package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Event - структура события в календаре
//NOTE: использую стороний пакет для uuid
type Event struct {
	ID        uuid.UUID  `json:"id"`
	User      string     `json:"user"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	StartTime *time.Time `json:"start"`
	EndTime   *time.Time `json:"end"`
}
