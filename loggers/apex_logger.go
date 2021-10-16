package loggers

import "errors"

type ApexLogger struct {
}

func NewApexLogger() (*ApexLogger, error) {
	return &ApexLogger{}, nil
}

func (zl *ApexLogger) GetName() string {
	return "apex-logger"
}

func (zl *ApexLogger) PlainInfo(message string) {
}

func (zl *ApexLogger) Info(message string, platform string, boundary, name string, args ...interface{}) {
}

func (zl *ApexLogger) Warning(message string, platform string, boundary, name string, args ...interface{}) {
}

func (zl *ApexLogger) Error(message string, platform string, boundary, name string, args ...interface{}) {
}

func (zl *ApexLogger) PlainError(message string) {
}

func (zl *ApexLogger) Debug(message string, platform string, boundary, name string, args ...interface{}) {
}

func (zl *ApexLogger) Close() error {
	return errors.New("failed to close logger")
}
