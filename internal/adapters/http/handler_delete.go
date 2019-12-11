package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// ResDelHTTPEventSuccess ...
type ResDelHTTPEventSuccess struct {
	Result string `json:"result"`
}

// ReqDelHTTPEvent ...
type ReqDelHTTPEvent struct {
	ID uuid.UUID `json:"id"`
}

// deleteHandler - handler для пути /api/delete_event
func (api *APIServerHTTP) deleteHandler(w http.ResponseWriter, r *http.Request) {
	var req ReqDelHTTPEvent
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "json.NewDecoder", err), http.StatusBadRequest)
		return
	}
	err = api.Event.Delete(context.Background(), req.ID)
	if err != nil {
		api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "api.Event.Delete", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResDelHTTPEventSuccess{Result: "OK"})
}
