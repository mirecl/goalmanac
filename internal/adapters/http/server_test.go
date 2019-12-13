package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/adapters/db"
	v "github.com/mirecl/goalmanac/internal/adapters/http/validate"
	"github.com/mirecl/goalmanac/internal/adapters/logger"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

// func TestRunServer(t *testing.T) {
// 	var buf bytes.Buffer
// 	log.SetOutput(&buf)
// 	log.SetOutput(os.Stdout)

// 	api, err := createAPI()
// 	require.NoError(t, err)
// 	go api.Serve()
// 	time.Sleep(20 * time.Second)

// 	t.Log(buf.String())
// 	require.Equal(t, "", "1")
// }
func TestHelloHandler(t *testing.T) {

	// Создаем инстанцию HTTP API
	api, err := createAPI()
	require.NoError(t, err)

	req, err := http.NewRequest("GET", "/hello", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.helloHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса запроса
	require.Equal(t, rr.Code, http.StatusOK)

	// Проверка ответа
	require.Equal(t, rr.Body.String(), "Hello")
}

func TestCountHandler(t *testing.T) {
	api, err := createAPI()
	require.NoError(t, err)

	req, err := http.NewRequest("GET", "/api/count_event", nil)
	require.NoError(t, err)

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

func TestCreateHandler(t *testing.T) {
	api, err := createAPI()
	require.NoError(t, err)

	requestBody, err := json.Marshal(map[string]string{
		"body":     "test",
		"duration": "10m",
		"start":    "2019-12-16T12:50:00.000Z",
		"title":    "Golang",
		"user":     "test",
	})

	req, err := http.NewRequest("POST", "/api/create_event", bytes.NewBuffer(requestBody))
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := api.logHandler(api.validateHandler(api.createHandler, v.Create))
	handler.ServeHTTP(rr, req)

	// Проверка статуса запроса
	require.Equal(t, rr.Code, http.StatusCreated)

	// Проверка ответа
	var resp ResCreateHTTPEventSuccess
	json.NewDecoder(rr.Body).Decode(&resp)
	require.NotEmpty(t, resp.Result)
}

func TestValidateHandler(t *testing.T) {
	api, err := createAPI()
	require.NoError(t, err)

	requestBody, err := json.Marshal(map[string]string{
		"body":     "test",
		"duration": "11m",
		"start":    "2019-12-16T12:50:00.000Z",
		"title":    "Golang",
		"user":     "test",
	})

	req, err := http.NewRequest("POST", "/api/create_event", bytes.NewBuffer(requestBody))
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := api.logHandler(api.validateHandler(api.createHandler, v.Create))
	handler.ServeHTTP(rr, req)

	// Проверка статуса запроса
	require.Equal(t, rr.Code, http.StatusBadRequest)

	// Проверка ответа
	var resp badHTTP
	json.NewDecoder(rr.Body).Decode(&resp)
	require.Equal(t, "duration: duration must be one of the following: \"10m\", \"20m\", \"30m\", \"40m\", \"50m\", \"60m\"", resp.Error)
}

func TestUpdateHandler(t *testing.T) {
	api, err := createAPI()
	require.NoError(t, err)

	requestBody, err := json.Marshal(map[string]string{
		"body":     "test",
		"duration": "10m",
		"start":    "2019-12-16T12:50:00.000Z",
		"title":    "Golang",
		"user":     "test",
	})

	req, err := http.NewRequest("POST", "/api/create_event", bytes.NewBuffer(requestBody))
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.createHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса запроса
	require.Equal(t, rr.Code, http.StatusCreated)

	// Проверка ответа
	var resp ResCreateHTTPEventSuccess
	json.NewDecoder(rr.Body).Decode(&resp)

	requestBodyUpd, err := json.Marshal(map[string]string{
		"id":       resp.Result.String(),
		"body":     "test2",
		"duration": "10m",
		"start":    "2019-12-16T12:50:00.000Z",
		"title":    "Golang",
		"user":     "test",
	})

	rr = httptest.NewRecorder()

	req, err = http.NewRequest("POST", "/api/update_event", bytes.NewBuffer(requestBodyUpd))
	require.NoError(t, err)

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(api.updateHandler)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var respUpd ResUpdHTTPEventSuccess
	json.NewDecoder(rr.Body).Decode(&respUpd)
	require.Equal(t, "OK", respUpd.Result)
}

func TestDeleteHandler(t *testing.T) {
	api, err := createAPI()
	require.NoError(t, err)

	id := uuid.NewV4()
	requestBody, err := json.Marshal(map[string]string{
		"id": id.String(),
	})

	req, err := http.NewRequest("POST", "/api/delete_event", bytes.NewBuffer(requestBody))
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.deleteHandler)
	handler.ServeHTTP(rr, req)

	// Проверка статуса запроса
	require.Equal(t, rr.Code, http.StatusBadRequest)

	// Проверка ответа
	var resp badHTTP
	json.NewDecoder(rr.Body).Decode(&resp)
	require.Equal(t, "No Data for Delete", resp.Error)
}

func createAPI() (*APIServerHTTP, error) {
	var cfg adapters.Config
	adapters.CreateConfig(".", &cfg)

	loggerEvent, err := logger.NewLogEvent(&cfg)
	if err != nil {
		return nil, err
	}

	// Создаем logger для событий в api http
	loggerHTTP, err := logger.NewLogHTTP(&cfg)
	if err != nil {
		return nil, err
	}
	// os.Remove("http.log")
	// os.Remove("event.log")

	// Создаем инстанция БД в памяти
	memdb, err := db.NewMemStorage()
	if err != nil {
		return nil, err
	}

	// Создаем интсанцию Бизнес-операцией с Календарем
	use := &usecases.EventUsecases{
		Storage: memdb,
		Logger:  loggerEvent,
	}

	var helper HelperHTTP
	err = CreateHelperHTTP(&helper)
	if err != nil {
		return nil, err
	}

	// Создаем инстанцию HTTP API
	return &APIServerHTTP{
		Event:  use,
		Logger: loggerHTTP,
		Config: &cfg,
		Helper: &helper,
	}, nil
}
