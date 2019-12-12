package interfaces

import (
	"context"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
)

// EventStorage - интерфейс для работы с событиями календаря
type EventStorage interface {
	Save(ctx context.Context, event *entities.Event) error
	GetCount(ctx context.Context) (int, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, event *entities.Event) error
	GetAll(ctx context.Context) ([]*entities.Event, error)
}

// EventLogger ...
type EventLogger interface {
	Errorf(path, format string, args ...interface{})
	Infof(format string, args ...interface{})
}
