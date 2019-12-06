package adapters

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ConfigHTTP ...
type ConfigHTTP struct {
	Host         string        `mapstructure:"host"`
	Port         string        `mapstructure:"port"`
	Shutdown     time.Duration `mapstructure:"shutdown"`
	WriteTimeout time.Duration `mapstructure:"writetimeout"`
	ReadTimeout  time.Duration `mapstructure:"readtimeout"`
}

// ConfigHTTPLog ...
type ConfigHTTPLog struct {
	Level string `mapstructure:"level"`
	Path  string `mapstructure:"path"`
}

// ConfigEVENTLog ...
type ConfigEVENTLog struct {
	Level string `mapstructure:"level"`
	Path  string `mapstructure:"path"`
}

// Config ...
type Config struct {
	HTTP     *ConfigHTTP     `mapstructure:"http"`
	LogHTTP  *ConfigHTTPLog  `mapstructure:"log_http"`
	LogEVENT *ConfigEVENTLog `mapstructure:"log_event"`
}

// CreateConfig ...
func CreateConfig(file string, cfg *Config) error {
	// Проверка файла конфигурации
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.WithFields(log.Fields{"type": "cmd"}).Errorln(err.Error())
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)

	// Значения по умолчанию.
	viper.SetDefault("log_http", map[string]string{
		"level": "info",
		"path":  "http.log",
	})
	viper.SetDefault("log_event", map[string]string{
		"level": "info",
		"path":  "event.log",
	})
	viper.SetDefault("http", map[string]interface{}{
		"host":         "127.0.0.1",
		"port":         "8080",
		"shutdown":     5,
		"writetimeout": 15,
		"readtimeout":  15,
	})

	// Чтения настроек
	if err := viper.ReadInConfig(); err == nil {
		log.WithFields(log.Fields{"type": "cmd"}).Info("Using config file:", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.WithFields(log.Fields{"type": "cmd"}).Errorln(err.Error())
		return err
	}
	return nil
}
