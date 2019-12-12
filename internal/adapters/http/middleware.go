package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/mirecl/goalmanac/internal/adapters/http/validate"
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
			api.Logger.Infof(&o.status, "%s %s %s", r.RequestURI, r.Method, time.Since(start))
		} else {
			api.Logger.Errorf(&o.status, GetFunc(), "%s %s %s", r.RequestURI, r.Method, time.Since(start))
		}
	})
}

// validateHandler - handler для Middleware
func (api *APIServerHTTP) validateHandler(next http.HandlerFunc, a validate.Validater) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Чтение входных параметров
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			api.Error(w, err, http.StatusBadRequest, GetFunc())
			return
		}
		// Валидация данных
		result, err := a.Validate(body)
		if err != nil {
			api.Error(w, err, http.StatusBadRequest, GetFunc())
			return
		}
		// Формирование ответа
		if !result.Valid() {
			var errS error
			for _, errD := range result.Errors() {
				if errS == nil {
					errS = fmt.Errorf("%s: %s", errD.Field(), errD.Description())
					continue
				}
				errS = fmt.Errorf("%s: %s  %w", errD.Field(), errD.Description(), errS)
			}
			api.Error(w, errS, http.StatusBadRequest, GetFunc())
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		next(w, r)
	})
}
