package db

import (
	"context"
	"sync"

	"github.com/mirecl/goalmanac/internal/domain/entities"
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
