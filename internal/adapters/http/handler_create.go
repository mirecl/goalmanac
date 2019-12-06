package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
)

// ResCreateHTTPEventSuccess ...
type ResCreateHTTPEventSuccess struct {
	Result uuid.UUID `json:"result"`
}

// ResCreateHTTPEventBad ...
type ResCreateHTTPEventBad struct {
	Error string `json:"error"`
}

// ReqCreateHTTPEvent ...
type ReqCreateHTTPEvent struct {
	User      string `json:"user"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	StartTime string `json:"start"`
	Duration  string `json:"duration"`
}

// addHandler - handler для пути /api/create_event
func (api *APIServerHTTP) createHandler(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateHTTPEvent
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.Logger.Errorf("%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResCreateHTTPEventBad{Error: err.Error()})
		return
	}
	new := &entities.Event{
		ID:        uuid.NewV4(),
		User:      req.User,
		Title:     req.Title,
		Body:      req.Body,
		StartTime: nil,
		EndTime:   nil}
	err = api.Event.Add(context.Background(), new)
	if err != nil {
		api.Logger.Errorf("%s", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResCreateHTTPEventSuccess{Result: new.ID})
}
