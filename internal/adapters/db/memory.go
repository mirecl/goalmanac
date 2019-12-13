package db

import (
	"context"
	"sync"

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
