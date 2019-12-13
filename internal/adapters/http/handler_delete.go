package http

import (
	"context"
	"encoding/json"
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
		api.Error(w, err, http.StatusBadRequest, F())
		return
	}
	err = api.Event.Delete(context.Background(), req.ID)
	if err != nil {
		api.Error(w, err, http.StatusBadRequest, F())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResDelHTTPEventSuccess{Result: "OK"})
}
