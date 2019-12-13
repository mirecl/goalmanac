package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
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
		TimeFormats:  []string{"2006-01-02T15:04:00.000Z"},
	}

	return nil
}

type badHTTP struct {
	Error string `json:"error"`
}

// Error ...
func (api *APIServerHTTP) Error(w http.ResponseWriter, err error, code int, path string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(badHTTP{Error: err.Error()})
	api.Logger.Errorf(&code, path, "%s", err.Error())
}

// F ...
func F() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
}
