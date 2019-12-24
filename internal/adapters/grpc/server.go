package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/adapters/grpc/api"
	"github.com/mirecl/goalmanac/internal/domain/interfaces"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
	"google.golang.org/grpc"
)

//APIServerGRPC - структура для gRPC-сервера
type APIServerGRPC struct {
	Event  *usecases.EventUsecases
	Logger interfaces.GRPCLogger
	Config *adapters.Config
}

// Serve - запуск gRPC-сервера
func (g *APIServerGRPC) Serve() error {
	host := g.Config.GRPC.Host
	port := g.Config.GRPC.Port
	serv, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return err
	}
	// Создаем цепочку middleware
	middleware := grpc_middleware.ChainUnaryServer(g.unaryInterceptor)
	g.Logger.Infof(nil, "Starting gRPC server on %s:%s", host, port)

	// Создаем gRPC-сервер
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware))
	api.RegisterAlmanacServer(grpcServer, g)

	// Запускаем gRPC-сервер
	go func() {
		if err := grpcServer.Serve(serv); err != nil {
			g.Logger.Errorf(nil, F(), "%s", err)
		}
	}()

	// Перехватываем сигналы завершения
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Graceful Shutdown
	stopped := make(chan struct{}, 1)
	go func() {
		grpcServer.GracefulStop()
		g.Logger.Infof(nil, "%s", "gRPC-server - Graceful Shutdown")
		close(stopped)
	}()

	// Ожидаем завершение сервиса
	t := time.NewTimer(10 * time.Second)
	select {
	case <-t.C:
		grpcServer.Stop()
		g.Logger.Infof(nil, "%s", "gRPC-server - Shutdown")
	case <-stopped:
		t.Stop()
	}
	return nil
}
