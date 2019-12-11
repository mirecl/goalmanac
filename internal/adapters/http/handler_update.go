package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
)

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

	// Считываем входящие данные
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "json.NewDecoder", err), http.StatusBadRequest)
		return
	}

	// Конвертируем время начала события
	startTime, err := api.Helper.time.Parse(req.StartTime)
	if err != nil {
		api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "api.Helper.time.Parse", err), http.StatusBadRequest)
		return
	}

	// Определяем время окончания события
	timeEvent, err := time.ParseDuration(req.Duration)
	if err != nil {
		api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "time.ParseDuration", err), http.StatusBadRequest)
		return
	}
	endTime := startTime.Add(timeEvent)

	changeEvent := &entities.Event{
		ID:        req.ID,
		User:      req.User,
		Title:     req.Title,
		Body:      req.Body,
		StartTime: &startTime,
		EndTime:   &endTime}

	err = api.Event.Update(context.Background(), changeEvent)
	if err != nil {
		api.Error(w, fmt.Errorf("Error in %s (%s) %w", GetFunc(), "api.Event.Update", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResUpdHTTPEventSuccess{Result: "OK"})
}
