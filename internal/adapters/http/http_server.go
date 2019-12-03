package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mirecl/goalmanac/internal/domain/interfaces"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
	"github.com/spf13/viper"
)

//APIServerHTTP - структура для http-сервера
type APIServerHTTP struct {
	Event  *usecases.EventUsecases
	Logger interfaces.HTTPLogger
	Config *viper.Viper
}

//Serve ...
func (api *APIServerHTTP) Serve() error {
	// Загружаем конфигурацию для запуска http-сервера
	httpCfg := api.Config.GetStringMapString("http")
	host := httpCfg["host"]
	port := httpCfg["port"]
	wait, _ := strconv.Atoi(httpCfg["shutdown"])

	// Создаем Route
	r := mux.NewRouter()

	// Устанавливаем handler ддля /hello
	r.HandleFunc("/hello", api.helloHandler)

	// Устанавливаем Middleware для log
	r.Use(api.logHandler)

	api.Logger.Infof("Starting http server on %s:%s", host, port)

	// Создаем инстанцию http-сервера
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", host, port),
	}

	// Запускаем http-сервер
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			api.Logger.Errorf("%s", err)
		}
	}()

	// Перехватываем сигналы завершения
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Graceful Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(wait)*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	api.Logger.Infof("%s", "shutting down")
	return nil
}

// helloHandler - handler для пути /hello
func (api *APIServerHTTP) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

// logHandler - handler для Middleware
func (api *APIServerHTTP) logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		api.Logger.Infof("%s %s %s", r.RequestURI, r.Method, time.Since(start))
	})
}
