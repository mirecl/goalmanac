package usecases

import (
	"context"
	"time"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	"github.com/mirecl/goalmanac/internal/domain/interfaces"
	uuid "github.com/satori/go.uuid"
)

//EventUsecases - структура для работы со всем внешними источниками данных
type EventUsecases struct {
	EventStorage interfaces.EventStorage
}

//AddEvent - создание события в календаре
func (es *EventUsecases) AddEvent(ctx context.Context, user, title, body string, startTime *time.Time, endTime *time.Time) error {
	event := &entities.Event{ID: uuid.NewV4(), User: user, Title: title, Body: body, StartTime: startTime, EndTime: endTime}
	if err := es.EventStorage.SaveEvent(ctx, event); err != nil {
		return err
	}
	return nil
}

//GetCountEvent ...
func (es *EventUsecases) GetCountEvent(ctx context.Context) (*int, error) {
	cnt, err := es.EventStorage.GetCountEvent(ctx)
	if err != nil {
		return nil, err
	}
	return &cnt, nil
}
