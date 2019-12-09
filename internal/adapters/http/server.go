package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/domain/interfaces"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
	"github.com/rs/cors"
)

//APIServerHTTP - структура для http-сервера
type APIServerHTTP struct {
	Event  *usecases.EventUsecases
	Logger interfaces.HTTPLogger
	Config *adapters.Config
	Helper *HelperHTTP
}

// Для SPA UI - Vue.js
type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

//Serve ...
func (api *APIServerHTTP) Serve() error {
	// Загружаем конфигурацию для запуска http-сервера
	host := api.Config.HTTP.Host
	port := api.Config.HTTP.Port
	wait := api.Config.HTTP.Shutdown
	writeTimeout := api.Config.HTTP.WriteTimeout
	readTimeout := api.Config.HTTP.ReadTimeout

	// Создаем Route
	r := mux.NewRouter()

	// Устанавливаем handler для /hello
	r.HandleFunc("/hello", api.helloHandler).Methods("GET")
	// Устанавливаем handler для /api/count_event
	r.HandleFunc("/api/count_event", api.cntHandler).Methods("GET")
	// Устанавливаем handler для /api/create_event
	r.HandleFunc("/api/create_event", api.createHandler).Methods("POST")
	// Устанавливаем handler для /api/create_event
	r.HandleFunc("/api/delete_event", api.deleteHandler).Methods("POST")
	// Устанавливаем handler для /api/update_event
	r.HandleFunc("/api/update_event", api.updateHandler).Methods("POST")
	// Устанавливаем handler для /api/all_event
	r.HandleFunc("/api/all_event", api.allHandler).Methods("GET")
	// Устанавливаем Middleware для log
	r.Use(api.logHandler)

	// Подключаем SPA
	spa := spaHandler{
		staticPath: "internal/ui",
		indexPath:  "index.html",
	}
	r.PathPrefix("/").Handler(spa)

	api.Logger.Infof("Starting http server on %s:%s", host, port)

	сrs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET"},
	})

	// Создаем инстанцию http-сервера
	srv := &http.Server{
		Handler:      сrs.Handler(r),
		Addr:         fmt.Sprintf("%s:%s", host, port),
		WriteTimeout: writeTimeout * time.Second,
		ReadTimeout:  readTimeout * time.Second,
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
	ctx, cancel := context.WithTimeout(context.Background(), wait*time.Second)
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
