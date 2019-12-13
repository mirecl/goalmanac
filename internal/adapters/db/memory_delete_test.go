package db

import (
	"context"
	"testing"
	"time"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestDelete1(t *testing.T) {
	memdb, err := NewMemStorage()
	require.NoError(t, err)

	err = memdb.Delete(context.Background(), uuid.NewV4())
	require.EqualError(t, err, ErrNotFoundDelete.Error())
}

func TestDelete2(t *testing.T) {
	start := time.Now()
	end := time.Now()
	id := uuid.NewV4()
	event := &entities.Event{ID: id, User: "Grazhdankov", Title: "Golang", Body: "Tutorial and  big test", StartTime: &start, EndTime: &end}

	memdb, err := NewMemStorage()
	require.NoError(t, err)

	err = memdb.Save(context.Background(), event)
	require.NoError(t, err)

	err = memdb.Delete(context.Background(), id)
	require.NoError(t, err)
}

func TestDelete3(t *testing.T) {
	start := time.Now()
	end := time.Now()
	id := uuid.NewV4()
	event := &entities.Event{ID: id, User: "Grazhdankov", Title: "Golang", Body: "Tutorial and  big test", StartTime: &start, EndTime: &end}

	memdb, err := NewMemStorage()
	require.NoError(t, err)

	err = memdb.Save(context.Background(), event)
	require.NoError(t, err)

	err = memdb.Delete(context.Background(), uuid.NewV4())
	require.EqualError(t, err, ErrNotFoundDelete.Error())
}
