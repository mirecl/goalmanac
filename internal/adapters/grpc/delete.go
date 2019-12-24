package grpc

import (
	"context"

	"github.com/mirecl/goalmanac/internal/adapters/grpc/api"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Delete - удалить событие по id
func (g *APIServerGRPC) Delete(ctx context.Context, r *api.EventDelete) (*api.ResponseOK, error) {
	id, err := uuid.FromString(r.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = g.Event.Delete(ctx, id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &api.ResponseOK{Result: "ok"}, nil
}
