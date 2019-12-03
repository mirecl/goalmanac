package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mirecl/goalmanac/internal/domain/interfaces"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
)

//APIServerHTTP ...
type APIServerHTTP struct {
	Event  *usecases.EventUsecases
	Logger interfaces.HTTPLogger
}

//Serve ...
func (api *APIServerHTTP) Serve(host, port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/hello", api.Hello)
	r.Use(api.logHandler)

	api.Logger.Infof("Starting http server on %s:%s", host, port)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", host, port),
	}

	return srv.ListenAndServe()
}

//Hello ...
func (api *APIServerHTTP) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// logHandler ...
func (api *APIServerHTTP) logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		api.Logger.Infof("%s %s %s", r.RequestURI, r.Method, time.Since(start))
	})
}
