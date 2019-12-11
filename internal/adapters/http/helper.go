package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jinzhu/now"
)

// HelperHTTP ...
type HelperHTTP struct {
	time *now.Config
}

// CreateHelperHTTP ...
func CreateHelperHTTP(cfg *HelperHTTP) error {

	location, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}
	cfg.time = &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: location,
		TimeFormats:  []string{"02.01.2006, 15:04:05"},
	}

	return nil
}

type badHTTP struct {
	Error string `json:"error"`
}

// Error ...
func (h *HelperHTTP) Error(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(badHTTP{Error: err.Error()})
}
