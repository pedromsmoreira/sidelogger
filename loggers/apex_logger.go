package loggers

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
)

type ApexLogger struct {
	log *log.Logger
}

func NewApexLogger() (*ApexLogger, error) {
	return &ApexLogger{
		log: &log.Logger{
			Handler: json.New(os.Stdout),
			Level:   log.InfoLevel,
		},
	}, nil
}

func (al *ApexLogger) GetName() string {
	return "apex-logger"
}

func (al *ApexLogger) PlainInfo(message string) {
	al.log.Info(message)
}

func (al *ApexLogger) Info(message string, args ...interface{}) {
	al.log.WithFields(
		log.Fields{
			"eventSeverity": "Info",
			"metadata":      args,
		},
	).Info(message)
}

func (al *ApexLogger) Warning(message string, args ...interface{}) {
	al.log.WithFields(
		log.Fields{
			"eventSeverity": "Warning",
			"metadata":      args,
		},
	).Warn(message)
}

func (al *ApexLogger) Error(message string, args ...interface{}) {
	al.log.WithFields(
		log.Fields{
			"eventSeverity": "Error",
			"metadata":      args,
		},
	).Error(message)
}

func (al *ApexLogger) PlainError(message string) {
	al.log.Error(message)
}

func (al *ApexLogger) Debug(message string, args ...interface{}) {
	al.log.WithFields(
		log.Fields{
			"eventSeverity": "Debug",
			"metadata":      args,
		},
	).Debug(message)
}

func (al *ApexLogger) Close() error {
	return nil
}
