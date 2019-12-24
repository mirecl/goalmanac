package grpc

import (
	"context"

	"github.com/mirecl/goalmanac/internal/adapters/grpc/api"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Update - обновление события
func (g *APIServerGRPC) Update(ctx context.Context, r *api.EventUpdate) (*api.ResponseOK, error) {
	// Определяем время старта и окончания события
	startTime, endTime, err := getTime(r.GetStarttime(), r.GetDuration())
	if err != nil {
		return nil, err
	}

	// Конвертируем uuid
	id, err := uuid.FromString(r.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Создаем событие
	new := &entities.Event{
		ID:        id,
		User:      r.GetUser(),
		Title:     r.GetTitle(),
		Body:      r.GetBody(),
		StartTime: startTime,
		EndTime:   endTime}

	// Сохраняем событие
	err = g.Event.Update(ctx, new)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &api.ResponseOK{Result: "ok"}, nil
}
