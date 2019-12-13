package adapters

import (
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestCreateConfigNull(t *testing.T) {
	var cfg Config
	err := CreateConfig("", &cfg)
	require.NoError(t, err)
	require.Equal(t, cfg.HTTP.Host, "127.0.0.1")
	require.Equal(t, cfg.HTTP.Shutdown, time.Duration(5))
	require.Equal(t, cfg.HTTP.Port, "8080")
}

func TestCreateConfigFile(t *testing.T) {
	var cfg Config
	path, _ := filepath.Abs("../../config.yaml")
	err := CreateConfig(path, &cfg)
	require.NoError(t, err)
	require.Equal(t, cfg.HTTP.Host, "127.0.0.1")
	require.Equal(t, cfg.HTTP.Shutdown, time.Duration(5))
	require.Equal(t, cfg.HTTP.Port, "8800")
}
