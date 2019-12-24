package cmd

import (
	"github.com/mirecl/goalmanac/internal/adapters/db"
	grpc "github.com/mirecl/goalmanac/internal/adapters/grpc"
	mux "github.com/mirecl/goalmanac/internal/adapters/http"
	"github.com/mirecl/goalmanac/internal/adapters/logger"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Запуск gRPC-сервера",
	Long: `Запуск gRPC-сервера. За основу взят пакет google.golang.org/grpc.
	В конфигурациоонм файле должны быть указану host и port.`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          GRPCinit,
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}

// GRPCinit ...
func GRPCinit(cmd *cobra.Command, args []string) error {
	// Создаем logger для событий в Календаре
	loggerEvent, err := logger.NewLogEvent(&cfg)
	if err != nil {
		return err
	}

	// Создаем logger для событий в api grpc
	loggerGRPC, err := logger.NewLogGRPC(&cfg)
	if err != nil {
		return err
	}

	// Создаем инстанция БД - PostgreSQL
	db, err := db.NewSQLStorage(&cfg)
	if err != nil {
		return err
	}

	// Создаем интсанцию Бизнес-операцией с Календарем
	use := &usecases.EventUsecases{
		Storage: db,
		Logger:  loggerEvent,
	}

	// Создаем helper для HTTP
	var helper mux.HelperHTTP
	err = mux.CreateHelperHTTP(&helper)
	if err != nil {
		return err
	}

	// Создаем инстанцию GRPC API
	serverGRPC := &grpc.APIServerGRPC{
		Event:  use,
		Logger: loggerGRPC,
		Config: &cfg,
	}

	// Запускаем GRPC API
	return serverGRPC.Serve()
}
