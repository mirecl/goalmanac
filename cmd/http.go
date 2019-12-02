package cmd

import (
	"errors"
	"fmt"

	"github.com/mirecl/goalmanac/internal/adapters/db"
	mux "github.com/mirecl/goalmanac/internal/adapters/http"
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
	// Загружаем конфигурацию с файла
	cfg := viper.GetStringMapString("http_listen")

	// Проверям host
	host, ok := cfg["ip"]
	if !ok {
		return errors.New("Укажите в config-файле переменную host")
	}

	// Проверям port
	port, ok := cfg["port"]
	if !ok {
		return errors.New("Укажите в config-файле переменную port")
	}

	fmt.Printf("Starting http server on host %s:%s\n", host, port)

	//Создаем инстанция БД в памяти
	memdb, _ := db.NewMemStorage()

	//Создаем интсанцию Бизнес-операцией с Календарем
	use := &usecases.EventUsecases{
		Storage: memdb,
	}
	//Создаемп инстанцию  HTTP API
	server := &mux.APIServerHTTP{
		Event: use,
	}
	//Запускаем http-сервер
	return server.Serve(host, port)
}
