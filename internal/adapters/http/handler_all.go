package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mirecl/goalmanac/internal/domain/entities"
)

// ResAllHTTPEventBad ...
type ResAllHTTPEventBad struct {
	Error string `json:"error"`
}

// ResAlldHTTPEventSuccess ...
type ResAlldHTTPEventSuccess struct {
	Result []*entities.Event `json:"result"`
}

// updateHandler - handler для пути /api/update_event
func (api *APIServerHTTP) allHandler(w http.ResponseWriter, r *http.Request) {
	data, err := api.Event.GetAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ResUpdHTTPEventBad{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResAlldHTTPEventSuccess{Result: data})
}
