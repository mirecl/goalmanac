package logger

import (
	"os"

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

// Errorf ...
func (log *LogHTTP) Errorf(format string, args ...interface{}) {
	log.StdOut.Infof(format, args...)
	log.File.Infof(format, args...)
}

// Infof ...
func (log *LogHTTP) Infof(format string, args ...interface{}) {
	log.StdOut.Infof(format, args...)
	log.File.Infof(format, args...)
}
