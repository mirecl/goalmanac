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
	db interfaces.EventStorage
}

//AddEvent - создание события в календаре
func (event *EventUsecases) AddEvent(ctx context.Context, user, title, body string, startTime *time.Time, endTime *time.Time) error {
	eventNew := &entities.Event{ID: uuid.NewV4(), User: user, Title: title, Body: body, StartTime: startTime, EndTime: endTime}
	if err := event.db.Save(ctx, eventNew); err != nil {
		return err
	}
	return nil
}

//GetCountEvent ...
func (event *EventUsecases) GetCountEvent(ctx context.Context) (*int, error) {
	cnt, err := event.db.GetCount(ctx)
	if err != nil {
		return nil, err
	}
	return &cnt, nil
}
