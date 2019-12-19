package db

import (
	"context"
	"testing"
	"time"

	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	start := time.Now()
	end := time.Now()
	event := &entities.Event{ID: uuid.NewV4(), User: "Grazhdankov", Title: "Golang", Body: "Tutorial and  big test", StartTime: &start, EndTime: &end}

	memdb, err := NewMemStorage()
	require.NoError(t, err)

	for i := 0; i < 20; i++ {
		err = memdb.Save(context.Background(), event)
		require.NoError(t, err)
	}
}

func TestAll(t *testing.T) {
	memdb, err := NewMemStorage()
	require.NoError(t, err)

	events, err := memdb.GetAll(context.Background())
	require.NoError(t, err)

	require.Equal(t, len(events), 0)
}
