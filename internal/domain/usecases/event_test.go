package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/mirecl/goalmanac/internal/adapters/db"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	memdb, _ := db.NewMemStorage()
	uses := &EventUsecases{Storage: memdb}

	start := time.Now()
	end := time.Now()
	ctx := context.Background()

	new := &entities.Event{
		ID:        uuid.NewV4(),
		User:      "Andrei",
		Title:     "Golang",
		Body:      "Ttxt",
		StartTime: &start,
		EndTime:   &end,
	}

	for i := 0; i < 20; i++ {
		_ = uses.Add(ctx, new)
	}
	cnt, _ := uses.GetCount(ctx)
	require.Equal(t, *cnt, 20)
}
