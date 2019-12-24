package logger

import (
	"os"

	"github.com/mirecl/goalmanac/internal/adapters"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

// LogGRPC - logger для gRPC-сервера
type LogGRPC struct {
	StdOut *log.Entry
	File   *log.Entry
}

// NewLogGRPC - создаем инстанцию
func NewLogGRPC(cfg *adapters.Config) (*LogGRPC, error) {
	var StdOut *log.Entry
	var File *log.Entry

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
	logLevel, err := log.ParseLevel(cfg.LogGRPC.Level)
	if err != nil {
		return nil, err
	}
	loggerStdOut.SetLevel(logLevel)
	StdOut = loggerStdOut.WithFields(log.Fields{"type": "grpc"})

	// Создаем инстанция для logger'a в file
	loggerFile := log.New()

	// Настройки отображения - json
	loggerFile.Formatter = &log.JSONFormatter{}

	// Указываем уровень логирования - логируем все
	loggerFile.SetLevel(log.InfoLevel)

	// Создаем/открываем файл для логирвания
	logFile, err := os.OpenFile(cfg.LogGRPC.Path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}

	// Указываем вывод в file
	loggerFile.SetOutput(logFile)
	File = loggerFile.WithFields(log.Fields{"type": "grpc"})

	return &LogGRPC{StdOut: StdOut, File: File}, nil
}

// Errorf - вывод ошибок
func (l *LogGRPC) Errorf(status *codes.Code, path, format string, args ...interface{}) {
	if status == nil {
		l.StdOut.WithFields(log.Fields{"func": path}).Errorf(format, args...)
		l.File.WithFields(log.Fields{"func": path}).Errorf(format, args...)
		return
	}
	l.StdOut.WithFields(log.Fields{"func": path, "status": *status}).Errorf(format, args...)
	l.File.WithFields(log.Fields{"func": path, "status": *status}).Errorf(format, args...)
}

// Infof - вывод информации
func (l *LogGRPC) Infof(status *codes.Code, format string, args ...interface{}) {
	if status == nil {
		l.StdOut.Infof(format, args...)
		l.File.Infof(format, args...)
		return
	}
	l.StdOut.WithFields(log.Fields{"status": *status}).Infof(format, args...)
	l.File.WithFields(log.Fields{"status": *status}).Infof(format, args...)
}
