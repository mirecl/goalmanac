package db

import (
	"context"
	"sync"

	"github.com/mirecl/goalmanac/internal/domain/entities"
)

//MemEventStorage ...
type MemEventStorage struct {
	mux *sync.Mutex
	db  []*entities.Event
	cnt int
}

//NewMemEventStorage ...
func NewMemEventStorage() (*MemEventStorage, error) {
	return &MemEventStorage{db: make([]*entities.Event, 0, 100), mux: &sync.Mutex{}}, nil
}

//SaveEvent ...
func (m *MemEventStorage) SaveEvent(ctx context.Context, event *entities.Event) error {
	m.mux.Lock()
	m.db = append(m.db, event)
	m.cnt++
	m.mux.Unlock()
	return nil
}

//GetCountEvent ...
func (m *MemEventStorage) GetCountEvent(ctx context.Context) (int, error) {
	m.mux.Lock()
	cnt := m.cnt
	m.mux.Unlock()
	return cnt, nil
}
