package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Event - структура события в календаре
//NOTE: использую стороний пакет для uuid
type Event struct {
	ID        uuid.UUID
	User      string
	Title     string
	Body      string
	StartTime *time.Time
	EndTime   *time.Time
}
