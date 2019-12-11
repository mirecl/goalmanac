package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mirecl/goalmanac/internal/domain/entities"
)

// ResAlldHTTPEventSuccess ...
type ResAlldHTTPEventSuccess struct {
	Result []*entities.Event `json:"result"`
}

// updateHandler - handler для пути /api/update_event
func (api *APIServerHTTP) allHandler(w http.ResponseWriter, r *http.Request) {
	data, err := api.Event.GetAll(context.Background())
	if err != nil {
		api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "api.Event.GetAll", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResAlldHTTPEventSuccess{Result: data})
}
