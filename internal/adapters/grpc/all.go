package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mirecl/goalmanac/internal/adapters/grpc/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAll - все события
func (g *APIServerGRPC) GetAll(ctx context.Context, _ *empty.Empty) (*api.ResponseEvents, error) {
	events, err := g.Event.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	res := make([]*api.Event, 0, len(events))

	for _, e := range events {
		startTime, err := ptypes.TimestampProto(*e.StartTime)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		endTime, err := ptypes.TimestampProto(*e.EndTime)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		// Создаем событие
		new := &api.Event{
			Id:        e.ID.String(),
			User:      e.User,
			Title:     e.Title,
			Body:      e.Body,
			Starttime: startTime,
			Endtime:   endTime}
		res = append(res, new)
	}
	return &api.ResponseEvents{Result: res}, nil
}

// GetDayEvent - все события сегодня
func (g *APIServerGRPC) GetDayEvent(ctx context.Context, r *api.EventUser) (*api.ResponseEvents, error) {
	events, err := g.Event.GetForDay(ctx, r.GetUser())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	res := make([]*api.Event, 0, len(events))

	for _, e := range events {
		startTime, err := ptypes.TimestampProto(*e.StartTime)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		endTime, err := ptypes.TimestampProto(*e.EndTime)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		// Создаем событие
		new := &api.Event{
			Id:        e.ID.String(),
			User:      e.User,
			Title:     e.Title,
			Body:      e.Body,
			Starttime: startTime,
			Endtime:   endTime}
		res = append(res, new)
	}
	return &api.ResponseEvents{Result: res}, nil
}

// GetWeekEvent - все события на этой недели
func (g *APIServerGRPC) GetWeekEvent(ctx context.Context, r *api.EventUser) (*api.ResponseEvents, error) {
	events, err := g.Event.GetForWeek(ctx, r.GetUser())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	res := make([]*api.Event, 0, len(events))

	for _, e := range events {
		startTime, err := ptypes.TimestampProto(*e.StartTime)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		endTime, err := ptypes.TimestampProto(*e.EndTime)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		// Создаем событие
		new := &api.Event{
			Id:        e.ID.String(),
			User:      e.User,
			Title:     e.Title,
			Body:      e.Body,
			Starttime: startTime,
			Endtime:   endTime}
		res = append(res, new)
	}
	return &api.ResponseEvents{Result: res}, nil
}

// GetMonthEvent - все события в этом месяце
func (g *APIServerGRPC) GetMonthEvent(ctx context.Context, r *api.EventUser) (*api.ResponseEvents, error) {
	events, err := g.Event.GetForMonth(ctx, r.GetUser())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	res := make([]*api.Event, 0, len(events))

	for _, e := range events {
		startTime, err := ptypes.TimestampProto(*e.StartTime)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		endTime, err := ptypes.TimestampProto(*e.EndTime)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		// Создаем событие
		new := &api.Event{
			Id:        e.ID.String(),
			User:      e.User,
			Title:     e.Title,
			Body:      e.Body,
			Starttime: startTime,
			Endtime:   endTime}
		res = append(res, new)
	}
	return &api.ResponseEvents{Result: res}, nil
}
