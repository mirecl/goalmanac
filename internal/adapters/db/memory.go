package db

import (
	"context"
	"sync"
	"time"

	"github.com/jinzhu/now"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
)

//MemEventStorage - структура БД в памяти
type MemEventStorage struct {
	mux *sync.Mutex
	db  []*entities.Event
	cnt int
}

//NewMemStorage - создаем инстанцию БД в памяти
func NewMemStorage() (*MemEventStorage, error) {
	return &MemEventStorage{db: make([]*entities.Event, 0, 100), mux: &sync.Mutex{}}, nil
}

//Save - сохраняем событие
func (m *MemEventStorage) Save(ctx context.Context, event *entities.Event) error {
	m.mux.Lock()
	m.db = append(m.db, event)
	m.cnt++
	m.mux.Unlock()
	return nil
}

//GetCount - получить общее количество событий
func (m *MemEventStorage) GetCount(ctx context.Context) (int, error) {
	m.mux.Lock()
	cnt := m.cnt
	m.mux.Unlock()
	return cnt, nil
}

// Delete ...
func (m *MemEventStorage) Delete(ctx context.Context, id uuid.UUID) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	if m.cnt == 0 {
		return ErrNotFoundDelete
	}
	for i, event := range m.db {
		if event.ID == id {
			m.db[i] = m.db[m.cnt-1]
			m.db = m.db[:m.cnt-1]
			m.cnt--
			return nil
		}
	}
	return ErrNotFoundDelete
}

// Update ...
func (m *MemEventStorage) Update(ctx context.Context, e *entities.Event) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	if m.cnt == 0 {
		return ErrNotFoundUpdate
	}
	for i, event := range m.db {
		if event.ID == e.ID {
			m.db[i] = e
			return nil
		}
	}
	return ErrNotFoundUpdate
}

// GetAll ...
func (m *MemEventStorage) GetAll(ctx context.Context) ([]*entities.Event, error) {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.db, nil
}

// GetForDay ...
func (m *MemEventStorage) GetForDay(ctx context.Context, user string) ([]*entities.Event, error) {
	event := make([]*entities.Event, 0, len(m.db))
	m.mux.Lock()
	s := now.BeginningOfDay()
	end := now.EndOfDay()
	defer m.mux.Unlock()
	for _, e := range m.db {
		if s.Before(*e.StartTime) && end.After(*e.StartTime) && e.User == user {
			event = append(event, e)
		}
	}
	return event, nil
}

// GetForWeek ...
func (m *MemEventStorage) GetForWeek(ctx context.Context, user string) ([]*entities.Event, error) {
	event := make([]*entities.Event, 0, len(m.db))
	location, err := time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}
	cfgTime := &now.Config{WeekStartDay: time.Monday, TimeLocation: location}
	s := cfgTime.With(time.Now()).BeginningOfWeek()
	end := cfgTime.With(time.Now()).EndOfWeek()
	m.mux.Lock()
	defer m.mux.Unlock()
	for _, e := range m.db {
		if s.Before(*e.StartTime) && end.After(*e.StartTime) && e.User == user {
			event = append(event, e)
		}
	}
	return event, nil
}

// GetForMonth ...
func (m *MemEventStorage) GetForMonth(ctx context.Context, user string) ([]*entities.Event, error) {
	event := make([]*entities.Event, 0, len(m.db))
	location, err := time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}
	cfgTime := &now.Config{WeekStartDay: time.Monday, TimeLocation: location}
	s := cfgTime.With(time.Now()).BeginningOfMonth()
	end := cfgTime.With(time.Now()).EndOfMonth()
	m.mux.Lock()
	defer m.mux.Unlock()
	for _, e := range m.db {
		if s.Before(*e.StartTime) && end.After(*e.StartTime) && e.User == user {
			event = append(event, e)
		}
	}
	return event, nil
}
