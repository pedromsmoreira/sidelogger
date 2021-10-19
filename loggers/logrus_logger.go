package loggers

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger() (*LogrusLogger, error) {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	}
	logger.Out = os.Stdout
	logger.Level = logrus.InfoLevel

	return &LogrusLogger{
		log: logger,
	}, nil
}

func (ll *LogrusLogger) GetName() string {
	return "logrus-logger"
}

func (ll *LogrusLogger) PlainInfo(message string) {
	ll.log.Info(message)
}

func (ll *LogrusLogger) Info(message string, platform string, boundary, name string, args ...interface{}) {
	ll.log.WithFields(
		logrus.Fields{
			"eventSeverity": "Info",
			"platform":      platform,
			"boundary":      boundary,
			"name":          name,
			"metadata":      args,
		},
	).Info(message)
}

func (ll *LogrusLogger) Warning(message string, platform string, boundary, name string, args ...interface{}) {
	ll.log.WithFields(
		logrus.Fields{
			"eventSeverity": "Warning",
			"platform":      platform,
			"boundary":      boundary,
			"name":          name,
			"metadata":      args,
		},
	).Warn(message)
}

func (ll *LogrusLogger) Error(message string, platform string, boundary, name string, args ...interface{}) {
	ll.log.WithFields(
		logrus.Fields{
			"eventSeverity": "Error",
			"platform":      platform,
			"boundary":      boundary,
			"name":          name,
			"metadata":      args,
		},
	).Error(message)
}

func (ll *LogrusLogger) PlainError(message string) {
	ll.log.Error(message)
}

func (ll *LogrusLogger) Debug(message string, platform string, boundary, name string, args ...interface{}) {
	ll.log.WithFields(
		logrus.Fields{
			"eventSeverity": "Debug",
			"platform":      platform,
			"boundary":      boundary,
			"name":          name,
			"metadata":      args,
		},
	).Debug(message)
}

func (ll *LogrusLogger) Close() error {
	return ll.log.Writer().Close()
}
