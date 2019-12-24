package grpc

import (
	"context"

	"github.com/mirecl/goalmanac/internal/adapters/grpc/api"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create - создание события
func (g *APIServerGRPC) Create(ctx context.Context, r *api.EventCreate) (*api.ResponseOK, error) {
	// Определяем время старта и окончания события
	startTime, endTime, err := getTime(r.GetStarttime(), r.GetDuration())
	if err != nil {
		return nil, err
	}

	// Создаем событие
	new := &entities.Event{
		ID:        uuid.NewV4(),
		User:      r.GetUser(),
		Title:     r.GetTitle(),
		Body:      r.GetBody(),
		StartTime: startTime,
		EndTime:   endTime}

	// Сохраняем событие
	err = g.Event.Add(ctx, new)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &api.ResponseOK{Result: new.ID.String()}, nil
}
