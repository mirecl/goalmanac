package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// LogEvent ...
type LogEvent struct {
	StdOut *log.Entry
	File   *log.Entry
}

// NewLogEvent ...
func NewLogEvent(logfile, level string) *LogEvent {
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

	switch level {
	case "info":
		loggerStdOut.SetLevel(log.InfoLevel)
		loggerFile.SetLevel(log.InfoLevel)
	case "debug":
		loggerStdOut.SetLevel(log.DebugLevel)
		loggerFile.SetLevel(log.DebugLevel)
	case "warn":
		loggerStdOut.SetLevel(log.WarnLevel)
		loggerFile.SetLevel(log.WarnLevel)
	case "error":
		loggerStdOut.SetLevel(log.ErrorLevel)
		loggerFile.SetLevel(log.ErrorLevel)
	}
	return &LogEvent{
		StdOut: loggerStdOut.WithFields(log.Fields{
			"type": "event",
		}),
		File: loggerFile.WithFields(log.Fields{
			"type": "event",
		}),
	}
}

// Errorf ...
func (log *LogEvent) Errorf(format string, args ...interface{}) {
	log.StdOut.Infof(format, args...)
	log.File.Infof(format, args...)
}

// Infof ...
func (log *LogEvent) Infof(format string, args ...interface{}) {
	log.StdOut.Infof(format, args...)
	log.File.Infof(format, args...)
}
