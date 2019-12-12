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
func NewLogEvent(cfg *adapters.Config) (*LogEvent, error) {
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
	logLevel, err := log.ParseLevel(cfg.LogEVENT.Level)
	if err != nil {
		return nil, err
		// log.WithFields(log.Fields{"type": "loggerEvent"}).Errorln(err.Error())
		// os.Exit(0)
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
		return nil, err
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
	}, nil
}

// Errorf - вывод ошибок
func (l *LogEvent) Errorf(path, format string, args ...interface{}) {
	l.StdOut.WithFields(log.Fields{"func": path}).Errorf(format, args...)
	l.File.WithFields(log.Fields{"func": path}).Errorf(format, args...)
}

// Infof - вывод информации
func (l *LogEvent) Infof(format string, args ...interface{}) {
	l.StdOut.Infof(format, args...)
	l.File.Infof(format, args...)
}
