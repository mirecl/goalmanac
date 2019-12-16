package usecases

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/adapters/db"
	"github.com/mirecl/goalmanac/internal/adapters/logger"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	"github.com/mirecl/goalmanac/internal/domain/errors"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

// TestAdd1 - добавление записи 20 раз
func TestAdd1(t *testing.T) {
	var cfg adapters.Config
	adapters.CreateConfig(".", &cfg)
	memdb, _ := db.NewMemStorage()

	uses := &EventUsecases{Storage: memdb, Logger: nil}

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

// TestAdd2 - добавление записи с некорректной датой
func TestAdd2(t *testing.T) {
	var cfg adapters.Config
	adapters.CreateConfig(".", &cfg)
	memdb, err := db.NewMemStorage()
	require.NoError(t, err)

	loggerEvent, err := logger.NewLogEvent(&cfg)
	require.NoError(t, err)
	os.Remove("event.log")

	uses := &EventUsecases{Storage: memdb, Logger: loggerEvent}

	start := time.Now()
	end := time.Now()
	ctx := context.Background()

	new := &entities.Event{
		ID:        uuid.NewV4(),
		User:      "Andrei",
		Title:     "Golang",
		Body:      "Txt",
		StartTime: &start,
		EndTime:   &end,
	}

	err = uses.Add(ctx, new)
	require.EqualError(t, err, errors.ErrAfterDay)
}

// TestDelete - удаление события в пустой таблицы
func TestDelete(t *testing.T) {
	var cfg adapters.Config
	adapters.CreateConfig(".", &cfg)
	memdb, err := db.NewMemStorage()
	require.NoError(t, err)

	loggerEvent, err := logger.NewLogEvent(&cfg)
	require.NoError(t, err)
	os.Remove("event.log")

	uses := &EventUsecases{Storage: memdb, Logger: loggerEvent}

	ctx := context.Background()

	err = uses.Delete(ctx, uuid.NewV4())
	require.EqualError(t, err, "No Data for Delete")
}

// TestUpdate - обновления события без ошибок
func TestUpdate(t *testing.T) {
	var cfg adapters.Config
	adapters.CreateConfig(".", &cfg)
	memdb, _ := db.NewMemStorage()

	uses := &EventUsecases{Storage: memdb, Logger: nil}

	start := time.Now().Local().Add(time.Duration(6) * time.Minute)
	end := time.Now().Local().Add(time.Duration(8) * time.Minute)
	ctx := context.Background()

	id := uuid.NewV4()
	new := &entities.Event{
		ID:        id,
		User:      "Andrei",
		Title:     "Golang",
		Body:      "Txt",
		StartTime: &start,
		EndTime:   &end,
	}

	err := uses.Add(ctx, new)
	require.NoError(t, err)

	new2 := &entities.Event{
		ID:        id,
		User:      "Ivan",
		Title:     "Golang",
		Body:      "Txt",
		StartTime: &start,
		EndTime:   &end,
	}
	err = uses.Update(ctx, new2)
	require.NoError(t, err)
}

// TestAll - получить все события
func TestAll(t *testing.T) {
	var err error
	var cfg adapters.Config
	adapters.CreateConfig(".", &cfg)
	memdb, _ := db.NewMemStorage()

	uses := &EventUsecases{Storage: memdb, Logger: nil}

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
		err = uses.Add(ctx, new)
		require.NoError(t, err)
	}
	data, err := uses.GetAll(ctx)
	require.NoError(t, err)
	require.Equal(t, len(data), 20)
}
