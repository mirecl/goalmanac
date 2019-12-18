package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Event - структура события в календаре
//NOTE: использую стороний пакет для uuid
type Event struct {
	ID        uuid.UUID  `json:"id"    db:"id"`
	User      string     `json:"user"  db:"user"`
	Title     string     `json:"title" db:"title"`
	Body      string     `json:"body"  db:"body"`
	StartTime *time.Time `json:"start" db:"starttime"`
	EndTime   *time.Time `json:"end"   db:"endtime"`
}
