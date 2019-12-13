package db

import (
	"context"
	"testing"
	"time"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestUpdate1(t *testing.T) {
	memdb, err := NewMemStorage()
	require.NoError(t, err)

	start := time.Now()
	end := time.Now()
	event := &entities.Event{ID: uuid.NewV4(), User: "Grazhdankov", Title: "Golang", Body: "Tutorial and  big test", StartTime: &start, EndTime: &end}

	err = memdb.Update(context.Background(), event)
	require.EqualError(t, err, ErrNotFoundUpdate.Error())
}

func TestUpdate2(t *testing.T) {
	memdb, err := NewMemStorage()
	require.NoError(t, err)

	start := time.Now()
	end := time.Now()
	event := &entities.Event{ID: uuid.NewV4(), User: "Grazhdankov", Title: "Golang", Body: "Tutorial and  big test", StartTime: &start, EndTime: &end}

	err = memdb.Save(context.Background(), event)
	require.NoError(t, err)

	err = memdb.Update(context.Background(), event)
	require.NoError(t, err)
}

func TestUpdate3(t *testing.T) {
	memdb, err := NewMemStorage()
	require.NoError(t, err)

	start := time.Now()
	end := time.Now()
	event := &entities.Event{ID: uuid.NewV4(), User: "Grazhdankov", Title: "Golang", Body: "Tutorial and  big test", StartTime: &start, EndTime: &end}

	err = memdb.Save(context.Background(), event)
	require.NoError(t, err)

	event1 := &entities.Event{ID: uuid.NewV4(), User: "Grazhdankov", Title: "Golang", Body: "Tutorial and  big test", StartTime: &start, EndTime: &end}

	err = memdb.Update(context.Background(), event1)
	require.EqualError(t, err, ErrNotFoundUpdate.Error())
}
