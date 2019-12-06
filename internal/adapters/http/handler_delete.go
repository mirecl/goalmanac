package http

import (
	"context"
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// ResDelHTTPEventBad ...
type ResDelHTTPEventBad struct {
	Error string `json:"error"`
}

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
		api.Logger.Errorf("%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResDelHTTPEventBad{Error: err.Error()})
		return
	}
	err = api.Event.Delete(context.Background(), req.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResDelHTTPEventBad{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResDelHTTPEventSuccess{Result: "OK"})
}
