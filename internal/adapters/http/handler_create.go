package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
)

// ResCreateHTTPEventSuccess ...
type ResCreateHTTPEventSuccess struct {
	Result uuid.UUID `json:"result"`
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

	// Считываем входящие данные
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.Logger.Errorf("%s", err.Error())
		api.Helper.Error(w, err, http.StatusBadRequest)
		return
	}

	// Конвертируем время начала события
	startTime, err := api.Helper.time.Parse(req.StartTime)
	if err != nil {
		api.Logger.Errorf("%s", err.Error())
		api.Helper.Error(w, err, http.StatusBadRequest)
		return
	}

	// Определяем время окончания события
	timeEvent, err := time.ParseDuration(req.Duration)
	if err != nil {
		api.Logger.Errorf("%s", err.Error())
		api.Helper.Error(w, err, http.StatusBadRequest)
		return
	}
	endTime := startTime.Add(timeEvent)

	// Создаем событие
	new := &entities.Event{
		ID:        uuid.NewV4(),
		User:      req.User,
		Title:     req.Title,
		Body:      req.Body,
		StartTime: &startTime,
		EndTime:   &endTime}

	// Сохраняем события
	err = api.Event.Add(context.Background(), new)
	if err != nil {
		api.Logger.Errorf("%s", err.Error())
		api.Helper.Error(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResCreateHTTPEventSuccess{Result: new.ID})
}
