package usecases

import (
	"context"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	"github.com/mirecl/goalmanac/internal/domain/interfaces"
)

//EventUsecases - структура для работы со всем внешними источниками данных
type EventUsecases struct {
	Storage interfaces.EventStorage
	Logger  interfaces.EventLogger
}

//Add - создание события в календаре
func (event *EventUsecases) Add(ctx context.Context, new *entities.Event) error {
	if err := event.Storage.Save(ctx, new); err != nil {
		return err
	}
	return nil
}

//GetCount - получить число общее число событий
func (event *EventUsecases) GetCount(ctx context.Context) (*int, error) {
	cnt, err := event.Storage.GetCount(ctx)
	if err != nil {
		return nil, err
	}
	return &cnt, nil
}
