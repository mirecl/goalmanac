package interfaces

import (
	"context"

	"github.com/mirecl/goalmanac/internal/domain/entities"
)

// EventStorage - интерфейс для работы с событиями календаря
type EventStorage interface {
	Save(ctx context.Context, event *entities.Event) error
	GetCount(ctx context.Context) (int, error)
	//... далее продолжу, когда будет реальная БД
}

// EventLogger ...
type EventLogger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
}
