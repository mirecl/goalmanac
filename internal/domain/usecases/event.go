package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/mirecl/goalmanac/internal/domain"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	"github.com/mirecl/goalmanac/internal/domain/errors"
	"github.com/mirecl/goalmanac/internal/domain/interfaces"
	uuid "github.com/satori/go.uuid"
)

// EventUsecases - структура для работы со всем внешними источниками данных
type EventUsecases struct {
	Storage interfaces.EventStorage
	Logger  interfaces.EventLogger
}

// Add - создание события в календаре
func (event *EventUsecases) Add(ctx context.Context, new *entities.Event) error {
	t := time.Now()
	if new.StartTime.Before(t) {
		event.Logger.Errorf(domain.GetFunc(), "%s %s", errors.ErrAfterDay, new.StartTime)
		return fmt.Errorf("%s", errors.ErrAfterDay)
	}
	if err := event.Storage.Save(ctx, new); err != nil {
		event.Logger.Errorf("%s %s", errors.ErrSaveEvent, err)
		return err
	}
	return nil
}

// GetCount - получить число общее число событий
func (event *EventUsecases) GetCount(ctx context.Context) (*int, error) {
	cnt, err := event.Storage.GetCount(ctx)
	if err != nil {
		event.Logger.Errorf(domain.GetFunc(), "%s: %s", errors.ErrGetCount, err)
		return nil, err
	}
	return &cnt, nil
}

// Delete - удалить событие по ID
func (event *EventUsecases) Delete(ctx context.Context, id uuid.UUID) error {
	err := event.Storage.Delete(ctx, id)
	if err != nil {
		event.Logger.Errorf(domain.GetFunc(), "%s", err)
		return err
	}
	return nil
}

// Update - изменения события по ID
func (event *EventUsecases) Update(ctx context.Context, e *entities.Event) error {
	err := event.Storage.Update(ctx, e)
	if err != nil {
		event.Logger.Errorf(domain.GetFunc(), "%s", err)
		return err
	}
	return nil
}

// GetAll - все события
func (event *EventUsecases) GetAll(ctx context.Context) ([]*entities.Event, error) {
	data, err := event.Storage.GetAll(ctx)
	if err != nil {
		event.Logger.Errorf(domain.GetFunc(), "%s", err)
		return nil, err
	}
	return data, nil
}
