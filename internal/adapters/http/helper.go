package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"time"

	"github.com/jinzhu/now"
	"github.com/xeipuuv/gojsonschema"
)

// HelperHTTP ...
type HelperHTTP struct {
	time         *now.Config
	schemaCreate gojsonschema.JSONLoader
	schemaChange gojsonschema.JSONLoader
}

// CreateHelperHTTP ...
func CreateHelperHTTP(cfg *HelperHTTP) error {

	path, err := filepath.Abs("internal/adapters/http/validate/createEvent.json")
	if err != nil {
		return fmt.Errorf("Error in filepath.Abs %w", err)
	}
	s, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error in %s (%s) %w", GetFunc(), "ioutil.ReadFile", err)
	}
	cfg.schemaCreate = gojsonschema.NewBytesLoader(s)

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

func (h *HelperHTTP) validateCreate(body []byte) (*gojsonschema.Result, error) {
	loader := gojsonschema.NewBytesLoader(body)
	result, err := gojsonschema.Validate(h.schemaCreate, loader)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return result, nil
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

// GetFunc ...
func GetFunc() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
}
