package loggers

import (
	"errors"
)

type LogrusLogger struct {
}

func NewLogrusLogger() (*LogrusLogger, error) {

	return &LogrusLogger{}, nil
}

func (zl *LogrusLogger) GetName() string {
	return "logrus-logger"
}

func (zl *LogrusLogger) PlainInfo(message string) {
}

func (zl *LogrusLogger) Info(message string, platform string, boundary, name string, args ...interface{}) {
}

func (zl *LogrusLogger) Warning(message string, platform string, boundary, name string, args ...interface{}) {
}

func (zl *LogrusLogger) Error(message string, platform string, boundary, name string, args ...interface{}) {
}

func (zl *LogrusLogger) PlainError(message string) {
}

func (zl *LogrusLogger) Debug(message string, platform string, boundary, name string, args ...interface{}) {
}

func (zl *LogrusLogger) Close() error {
	return errors.New("failed to close logger")
}
