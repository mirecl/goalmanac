package http

import (
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var re = regexp.MustCompile(`([2]0[0-9])`)

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
