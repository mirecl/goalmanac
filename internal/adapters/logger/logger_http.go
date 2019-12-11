package logger

import (
	"os"

	"github.com/mirecl/goalmanac/internal/adapters"
	log "github.com/sirupsen/logrus"
)

// LogHTTP - logger для событий в http
type LogHTTP struct {
	StdOut *log.Entry
	File   *log.Entry
}

// NewLogHTTP - создаем инстанцию
func NewLogHTTP(cfg *adapters.Config) (*LogHTTP, error) {
	// Создаем инстанция для logger'a в Stdout
	loggerStdOut := log.New()
	// Настройки отображения
	loggerStdOut.Formatter = &log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		QuoteEmptyFields:       true,
		DisableLevelTruncation: true,
	}
	// Указываем вывод в Stdout
	loggerStdOut.SetOutput(os.Stdout)

	// Указываем уровень логирования для Stdout
	logLevel, err := log.ParseLevel(cfg.LogHTTP.Level)
	if err != nil {
		return nil, err
	}
	loggerStdOut.SetLevel(logLevel)

	// Создаем инстанция для logger'a в file
	loggerFile := log.New()

	// Настройки отображения - json
	loggerFile.Formatter = &log.JSONFormatter{}

	// Указываем уровень логирования - логируем все
	loggerFile.SetLevel(log.InfoLevel)

	// Создаем/открываем файл для логирвания
	logFile, err := os.OpenFile(cfg.LogEVENT.Path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	// Указываем вывод в file
	loggerFile.SetOutput(logFile)

	return &LogHTTP{
		StdOut: loggerStdOut.WithFields(log.Fields{
			"type": "http",
		}),
		File: loggerFile.WithFields(log.Fields{
			"type": "http",
		}),
	}, nil
}

// Errorf - вывод ошибок
func (l *LogHTTP) Errorf(format string, args ...interface{}) {
	l.StdOut.Errorf(format, args...)
	l.File.Errorf(format, args...)
}

// Infof - вывод информации
func (l *LogHTTP) Infof(format string, args ...interface{}) {
	l.StdOut.Infof(format, args...)
	l.File.Infof(format, args...)
}
