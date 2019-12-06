package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
)

// ResUpdHTTPEventBad ...
type ResUpdHTTPEventBad struct {
	Error string `json:"error"`
}

// ResUpdHTTPEventSuccess ...
type ResUpdHTTPEventSuccess struct {
	Result string `json:"result"`
}

// ReqUpdHTTPEvent ...
type ReqUpdHTTPEvent struct {
	ID        uuid.UUID `json:"id"`
	User      string    `json:"user"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	StartTime string    `json:"start"`
	Duration  string    `json:"duration"`
}

// updateHandler - handler для пути /api/update_event
func (api *APIServerHTTP) updateHandler(w http.ResponseWriter, r *http.Request) {
	var req ReqUpdHTTPEvent
	err := json.NewDecoder(r.Body).Decode(&req)
	changeEvent := &entities.Event{
		ID:        req.ID,
		User:      req.User,
		Title:     req.Title,
		Body:      req.Body,
		StartTime: nil,
		EndTime:   nil}
	if err != nil {
		api.Logger.Errorf("%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResDelHTTPEventBad{Error: err.Error()})
		return
	}
	err = api.Event.Update(context.Background(), changeEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResUpdHTTPEventBad{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResUpdHTTPEventSuccess{Result: "OK"})
}
