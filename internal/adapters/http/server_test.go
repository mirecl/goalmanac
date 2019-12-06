package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/adapters/db"
	"github.com/mirecl/goalmanac/internal/adapters/logger"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
	"github.com/stretchr/testify/require"
)

func TestHelloHandler(t *testing.T) {
	// Создаем инстанцию HTTP API
	api := createAPI()

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.helloHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса запроса
	require.Equal(t, rr.Code, http.StatusOK)

	// Проверка ответа
	require.Equal(t, rr.Body.String(), "Hello")
}

func TestCountHandler(t *testing.T) {
	api := createAPI()

	req, err := http.NewRequest("GET", "/api/count_event", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.cntHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса запроса
	require.Equal(t, rr.Code, http.StatusOK)

	// Проверка ответа
	var resp cntHTTPEvent
	json.NewDecoder(rr.Body).Decode(&resp)
	require.Equal(t, resp.Result, 0)
}

func createAPI() APIServerHTTP {
	path, _ := filepath.Abs("../../../config.yaml")
	var cfg adapters.Config
	adapters.CreateConfig(path, &cfg)

	// Создаем logger для событий в Календаре
	loggerEvent := logger.NewLogEvent(&cfg)

	// Создаем logger для событий в api http
	loggerHTTP := logger.NewLogHTTP(&cfg)

	// Создаем инстанция БД в памяти
	memdb, _ := db.NewMemStorage()

	// Создаем интсанцию Бизнес-операцией с Календарем
	use := &usecases.EventUsecases{
		Storage: memdb,
		Logger:  loggerEvent,
	}

	// Создаем инстанцию HTTP API
	return APIServerHTTP{
		Event:  use,
		Logger: loggerHTTP,
		Config: &cfg,
	}
}
