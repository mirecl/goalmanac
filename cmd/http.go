package cmd

import (
	"github.com/mirecl/goalmanac/internal/adapters/db"
	mux "github.com/mirecl/goalmanac/internal/adapters/http"
	"github.com/mirecl/goalmanac/internal/adapters/logger"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	// Загружаем конфигурацию для запуска http-сервера
	httpCfg := viper.GetStringMapString("http")
	logCfg := viper.GetStringMapString("log")

	// считываем данные из config
	level := logCfg["level"]
	path := logCfg["path"]
	host := httpCfg["host"]
	port := httpCfg["port"]

	// Создаем logger для событий в Календаре
	loggerEvent := logger.NewLogEvent(path, level)

	// Создаем logger для событий в api http
	loggerHTTP := logger.NewLogHTTP(path)

	//Создаем инстанция БД в памяти
	memdb, _ := db.NewMemStorage()

	//Создаем интсанцию Бизнес-операцией с Календарем
	use := &usecases.EventUsecases{
		Storage: memdb,
		Logger:  loggerEvent,
	}

	//Создаем инстанцию  HTTP API
	server := &mux.APIServerHTTP{
		Event:  use,
		Logger: loggerHTTP,
	}
	//Запускаем http-сервер
	return server.Serve(host, port)
}
