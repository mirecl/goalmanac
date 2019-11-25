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
	memdb, _ := NewMemEventStorage()
	for i := 0; i < 20; i++ {
		memdb.Save(context.Background(), event)
	}

	cnt, _ := memdb.GetCount(context.Background())
	require.Equal(t, cnt, 20)
}
