package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// LogEvent - logger для событий в календаре
type LogEvent struct {
	StdOut *log.Entry
	File   *log.Entry
}

// NewLogEvent - создаем инстанцию
func NewLogEvent(logfile, level string) *LogEvent {
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
	switch level {
	case "info":
		loggerStdOut.SetLevel(log.InfoLevel)
	case "debug":
		loggerStdOut.SetLevel(log.DebugLevel)
	case "warn":
		loggerStdOut.SetLevel(log.WarnLevel)
	case "error":
		loggerStdOut.SetLevel(log.ErrorLevel)
	}

	// Создаем инстанция для logger'a в file
	loggerFile := log.New()

	// Настройки отображения - json
	loggerFile.Formatter = &log.JSONFormatter{}

	// Указываем уровень логирования - логируем все
	loggerFile.SetLevel(log.InfoLevel)

	// Создаем/открываем файл для логирвания
	logFile, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {

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
	log.StdOut.Infof(format, args...)
	log.File.Infof(format, args...)
}

// Infof - вывод информации
func (log *LogEvent) Infof(format string, args ...interface{}) {
	log.StdOut.Infof(format, args...)
	log.File.Infof(format, args...)
}
