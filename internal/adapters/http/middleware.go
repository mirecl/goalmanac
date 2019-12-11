package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var re = regexp.MustCompile(`([2-3]0[0-9])`)

type responseObserver struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *responseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}

// logHandler - handler для Middleware
func (api *APIServerHTTP) logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		o := &responseObserver{ResponseWriter: w}
		next.ServeHTTP(o, r)
		if re.MatchString(strconv.Itoa(o.status)) {
			api.Logger.Infof("%s %s %s %d", r.RequestURI, r.Method, time.Since(start), o.status)
		} else {
			api.Logger.Errorf("%s %s %s %d", r.RequestURI, r.Method, time.Since(start), o.status)
		}
	})
}

// validateHandlerCreate - handler для Middleware
func (api *APIServerHTTP) validateHandlerCreate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "ioutil.ReadAll", err), http.StatusBadRequest)
			return
		}
		result, err := api.Helper.validateCreate(body)
		if err != nil {
			api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "api.Helper.validateCreate", err), http.StatusBadRequest)
			return
		}
		if !result.Valid() {
			err = fmt.Errorf("Ошибка валидации данных в коде: %s)", GetFunc())
			for _, errD := range result.Errors() {
				err = fmt.Errorf("%s: %s %s %w", errD.Field(), errD.Description(), GetFunc(), err)
			}
			api.Error(w, err, http.StatusBadRequest)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		next(w, r)
	})
}
