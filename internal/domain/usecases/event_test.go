package usecases

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/adapters/db"
	"github.com/mirecl/goalmanac/internal/adapters/logger"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	path, _ := filepath.Abs("../../../config.yaml")
	var cfg adapters.Config
	adapters.CreateConfig(path, &cfg)
	memdb, _ := db.NewMemStorage()
	loggerEvent, _ := logger.NewLogEvent(&cfg)

	uses := &EventUsecases{Storage: memdb, Logger: loggerEvent}

	start := time.Now().Local().Add(time.Duration(6) * time.Minute)
	end := time.Now().Local().Add(time.Duration(8) * time.Minute)
	ctx := context.Background()

	new := &entities.Event{
		ID:        uuid.NewV4(),
		User:      "Andrei",
		Title:     "Golang",
		Body:      "Txt",
		StartTime: &start,
		EndTime:   &end,
	}

	for i := 0; i < 20; i++ {
		_ = uses.Add(ctx, new)
	}
	cnt, _ := uses.GetCount(ctx)
	require.Equal(t, *cnt, 20)
}
