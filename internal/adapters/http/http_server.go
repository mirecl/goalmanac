package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
)

//APIServerHTTP ...
type APIServerHTTP struct {
	Event *usecases.EventUsecases
}

//Serve ...
func (api *APIServerHTTP) Serve(host, port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", api.Hello)

	srv := &http.Server{
		Handler: r,
		Addr:    host + ":" + port,
	}
	return srv.ListenAndServe()
}

//Hello ...
func (api *APIServerHTTP) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
