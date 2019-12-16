package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mirecl/goalmanac/internal/domain/entities"
)

// ResAlldHTTPEventSuccess ...
type ResAlldHTTPEventSuccess struct {
	Result []*entities.Event `json:"result"`
}

// allHandler - handler для пути /api/all_event
func (api *APIServerHTTP) allHandler(w http.ResponseWriter, r *http.Request) {
	data, err := api.Event.GetAll(context.Background())
	if err != nil {
		api.Error(w, err, http.StatusBadRequest, F())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResAlldHTTPEventSuccess{Result: data})
}

// getDayHandler - handler для пути /api/events_for_day
func (api *APIServerHTTP) getDayHandler(w http.ResponseWriter, r *http.Request) {
	// Считываем аргументы
	args := r.URL.Query()
	user := args.Get("user")
	if user == "" {
		api.Error(w, errors.New("Не указан пользователь"), http.StatusBadRequest, F())
		return
	}
	data, err := api.Event.GetForDay(context.Background(), user)
	if err != nil {
		api.Error(w, err, http.StatusBadRequest, F())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResAlldHTTPEventSuccess{Result: data})
}

// getWeekHandler - handler для пути /api/events_for_week
func (api *APIServerHTTP) getWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Считываем аргументы
	args := r.URL.Query()
	user := args.Get("user")
	if user == "" {
		api.Error(w, errors.New("Не указан пользователь"), http.StatusBadRequest, F())
		return
	}
	data, err := api.Event.GetForWeek(context.Background(), user)
	if err != nil {
		api.Error(w, err, http.StatusBadRequest, F())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResAlldHTTPEventSuccess{Result: data})
}

// getMonthHandler - handler для пути /api/events_for_month
func (api *APIServerHTTP) getMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Считываем аргументы
	args := r.URL.Query()
	user := args.Get("user")
	if user == "" {
		api.Error(w, errors.New("Не указан пользователь"), http.StatusBadRequest, F())
		return
	}
	data, err := api.Event.GetForMonth(context.Background(), user)
	if err != nil {
		api.Error(w, err, http.StatusBadRequest, F())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResAlldHTTPEventSuccess{Result: data})
}
