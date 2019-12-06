package cmd

import (
	"github.com/mirecl/goalmanac/internal/adapters/db"
	mux "github.com/mirecl/goalmanac/internal/adapters/http"
	"github.com/mirecl/goalmanac/internal/adapters/logger"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Запуск http-сервера",
	Long: `Запуск http-сервера. За основу взят пакет https://github.com/gorilla/mux.
В конфигурациоонм файле должны быть указану ip и port.`,
	SilenceUsage: true,
	RunE:         HTTPinit,
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

// HTTPinit ...
func HTTPinit(cmd *cobra.Command, args []string) error {
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
	server := &mux.APIServerHTTP{
		Event:  use,
		Logger: loggerHTTP,
		Config: &cfg,
	}
	// Запускаем http-сервер
	return server.Serve()
}
