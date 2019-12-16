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
В конфигурациоонм файле должны быть указану host и port.`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          HTTPinit,
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

// HTTPinit ...
func HTTPinit(cmd *cobra.Command, args []string) error {
	// Создаем logger для событий в Календаре
	loggerEvent, err := logger.NewLogEvent(&cfg)
	if err != nil {
		return err
	}

	// Создаем logger для событий в api http
	loggerHTTP, err := logger.NewLogHTTP(&cfg)
	if err != nil {
		return err
	}

	// Создаем инстанция БД в памяти
	memdb, err := db.NewMemStorage()
	if err != nil {
		return err
	}

	// Создаем интсанцию Бизнес-операцией с Календарем
	use := &usecases.EventUsecases{
		Storage: memdb,
		Logger:  loggerEvent,
	}

	// Создаем helper для HTTP
	var helper mux.HelperHTTP
	err = mux.CreateHelperHTTP(&helper)
	if err != nil {
		return err
	}

	// Создаем инстанцию HTTP API
	server := &mux.APIServerHTTP{
		Event:  use,
		Logger: loggerHTTP,
		Config: &cfg,
		Helper: &helper,
	}
	// Запускаем http-сервер
	return server.Serve()
}
