package logger

import (
	"os"

	"github.com/mirecl/goalmanac/internal/adapters"
	log "github.com/sirupsen/logrus"
)

// LogEvent - logger для событий в календаре
type LogEvent struct {
	StdOut *log.Entry
	File   *log.Entry
}

// NewLogEvent - создаем инстанцию
func NewLogEvent(cfg *adapters.Config) *LogEvent {
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
		log.WithFields(log.Fields{"type": "loggerEvent"}).Errorln(err.Error())
		os.Exit(0)
	}
	loggerStdOut.SetLevel(logLevel)

	// Создаем инстанция для logger'a в file
	loggerFile := log.New()

	// Настройки отображения - json
	loggerFile.Formatter = &log.JSONFormatter{}

	// Указываем уровень логирования - логируем все
	loggerFile.SetLevel(log.InfoLevel)

	// Создаем/открываем файл для логирвания
	logFile, err := os.OpenFile(cfg.LogHTTP.Path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.WithFields(log.Fields{"type": "loggerEvent"}).Errorln(err.Error())
		os.Exit(0)
	}

	// Указываем вывод в file
	loggerFile.SetOutput(logFile)

	return &LogEvent{
		StdOut: loggerStdOut.WithFields(log.Fields{
			"type": "event",
		}),
		File: loggerFile.WithFields(log.Fields{
			"type": "event",
		}),
	}
}

// Errorf - вывод ошибок
func (log *LogEvent) Errorf(format string, args ...interface{}) {
	log.StdOut.Errorf(format, args...)
	log.File.Errorf(format, args...)
}

// Infof - вывод информации
func (log *LogEvent) Infof(format string, args ...interface{}) {
	log.StdOut.Infof(format, args...)
	log.File.Infof(format, args...)
}
