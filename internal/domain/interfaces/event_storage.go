package interfaces

import (
	"context"

	"github.com/mirecl/goalmanac/internal/domain/entities"
)

//EventStorage - интерфейс для работы с usecases календаря
type EventStorage interface {
	SaveEvent(ctx context.Context, event *entities.Event) error
	GetCountEvent(ctx context.Context) (int, error)
}
