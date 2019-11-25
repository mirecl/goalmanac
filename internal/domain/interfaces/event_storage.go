package interfaces

import (
	"context"

	"github.com/mirecl/goalmanac/internal/domain/entities"
)

//EventStorage - интерфейс для работы с usecases календаря
type EventStorage interface {
	Save(ctx context.Context, event *entities.Event) error
	GetCount(ctx context.Context) (int, error)
}
