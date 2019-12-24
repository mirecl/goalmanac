package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (g *APIServerGRPC) unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	m, err := handler(ctx, req)
	// Считываем статус запроса
	errStatus, _ := status.FromError(err)
	status := errStatus.Code()
	// Формируем log
	if err != nil {
		g.Logger.Errorf(&status, F(), "%s %s %s", info.FullMethod, time.Since(start))
	} else {
		g.Logger.Infof(&status, "%s %s", info.FullMethod, time.Since(start))
	}
	return m, err
}
