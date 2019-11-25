package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/mirecl/goalmanac/internal/adapters/db"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	memdb, _ := db.NewMemStorage()
	uses := &EventUsecases{db: memdb}

	start := time.Now()
	end := time.Now()
	ctx := context.Background()

	for i := 0; i < 20; i++ {
		_ = uses.AddEvent(ctx, "Grazhdankov", "Golang", "Tutorial", &start, &end)
	}
	cnt, _ := uses.GetCountEvent(ctx)
	require.Equal(t, *cnt, 20)
}
