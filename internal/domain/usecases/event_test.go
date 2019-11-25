package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/mirecl/goalmanac/internal/adapters/db"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	memdb, _ := db.NewMemEventStorage()
	uses := &EventUsecases{db: memdb}
	start := time.Now()
	end := time.Now()
	for i := 0; i < 20; i++ {
		_ = uses.AddEvent(context.Background(), "Grazhdankov", "Golang", "Tutorial and  big test", &start, &end)
	}
	cnt, _ := uses.GetCountEvent(context.Background())
	require.Equal(t, *cnt, 20)
}
