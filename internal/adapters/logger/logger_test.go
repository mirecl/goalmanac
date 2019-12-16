package logger

import (
	"os"
	"testing"

	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	defer os.Remove("event.log")
	var cfg adapters.Config
	adapters.CreateConfig(".", &cfg)
	_, err := NewLogEvent(&cfg)
	require.NoError(t, err)
}

func TestCreateHTTP(t *testing.T) {
	defer os.Remove("http.log")
	var cfg adapters.Config
	adapters.CreateConfig(".", &cfg)
	_, err := NewLogHTTP(&cfg)
	require.NoError(t, err)
}
