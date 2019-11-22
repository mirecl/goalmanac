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
}

//NewMemEventStorage ...
func NewMemEventStorage() (*MemEventStorage, error) {
	return &MemEventStorage{db: make([]*entities.Event, 0, 100), mux: &sync.Mutex{}}, nil
}

//SaveEvent ...
func (mem *MemEventStorage) SaveEvent(ctx context.Context, event *entities.Event) error {
	mem.mux.Lock()
	mem.db = append(mem.db, event)
	mem.mux.Unlock()
	return nil
}

//GetCountEvent ...
func (mem *MemEventStorage) GetCountEvent(ctx context.Context) (int, error) {
	mem.mux.Lock()
	cnt := len(mem.db)
	mem.mux.Unlock()
	return cnt, nil
}
