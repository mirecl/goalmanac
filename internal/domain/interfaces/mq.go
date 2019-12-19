package interfaces

import (
	"context"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
)

// MQStorage - интерфейс для работы с событиями календаря в MQ
type MQStorage interface {
	GetEventNotify(ctx context.Context, period string) ([]*entities.Event, error)
	ChangeStatusEventNotify(ctx context.Context, id uuid.UUID) error
}

// MQLogger ...
type MQLogger interface {
	Errorf(path, format string, args ...interface{})
	Infof(format string, args ...interface{})
}
