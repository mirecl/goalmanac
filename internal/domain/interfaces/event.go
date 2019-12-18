package interfaces

import (
	"context"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
)

// EventStorage - интерфейс для работы с событиями календаря
type EventStorage interface {
	Save(ctx context.Context, event *entities.Event) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, event *entities.Event) error
	GetAll(ctx context.Context) ([]*entities.Event, error)
	GetForDay(ctx context.Context, user string) ([]*entities.Event, error)
	GetForWeek(ctx context.Context, user string) ([]*entities.Event, error)
	GetForMonth(ctx context.Context, user string) ([]*entities.Event, error)
}

// EventLogger ...
type EventLogger interface {
	Errorf(path, format string, args ...interface{})
	Infof(format string, args ...interface{})
}
