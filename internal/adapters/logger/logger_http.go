package logger

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// LogHTTP ...
type LogHTTP struct {
	StdOut *log.Entry
	File   *log.Entry
}

// NewLogHTTP ...
func NewLogHTTP(logfile string) *LogHTTP {
	loggerStdOut := log.New()
	loggerStdOut.Formatter = &log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		QuoteEmptyFields:       true,
		DisableLevelTruncation: true,
	}
	loggerStdOut.SetOutput(os.Stdout)

	logFile, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {

	}
	loggerFile := log.New()
	loggerFile.Formatter = &log.JSONFormatter{}
	loggerFile.SetOutput(logFile)

	loggerStdOut.SetLevel(log.InfoLevel)
	loggerFile.SetLevel(log.InfoLevel)
	return &LogHTTP{
		StdOut: loggerStdOut.WithFields(log.Fields{
			"type": "http",
		}),
		File: loggerFile.WithFields(log.Fields{
			"type": "http",
		}),
	}
}

// Error ...
func (log *LogHTTP) Error(args ...string) {
	msg := strings.Join(args, " ")
	log.StdOut.Info(msg)
	log.File.Info(msg)
}

// Info ...
func (log *LogHTTP) Info(args ...string) {
	msg := strings.Join(args, " ")
	log.StdOut.Info(msg)
	log.File.Info(msg)
}

// Infof ...
func (log *LogHTTP) Infof(format string, args ...interface{}) {
	log.StdOut.Infof(format, args...)
	log.File.Infof(format, args...)
}
