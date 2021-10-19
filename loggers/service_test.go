package loggers

import (
	"fmt"
	"testing"
)

func TestLoggerService(t *testing.T) {
	t.Run("log when log level is Info should call info", func(t *testing.T) {
		ml := NewMockLogger()
		lsvc := &LoggerService{
			logger: ml,
		}

		err := lsvc.Log(&LogRequest{
			LogLevel: "Info",
		})

		if err != nil && ml.methodCalled && ml.methodName == "Info" {
			t.Errorf(fmt.Sprintf("Error when logging %v", err.Error()))
		}
	})

	t.Run("log when log level is Warning should call Warning", func(t *testing.T) {
		ml := NewMockLogger()
		lsvc := &LoggerService{
			logger: ml,
		}

		err := lsvc.Log(&LogRequest{
			LogLevel: "Warning",
		})

		if err != nil && ml.methodCalled && ml.methodName == "Warning" {
			t.Errorf(fmt.Sprintf("Error when logging %v", err.Error()))
		}
	})

	t.Run("log when log level is Error should call Error", func(t *testing.T) {
		ml := NewMockLogger()
		lsvc := &LoggerService{
			logger: ml,
		}

		err := lsvc.Log(&LogRequest{
			LogLevel: "Error",
		})

		if err != nil && ml.methodCalled && ml.methodName == "Error" {
			t.Errorf(fmt.Sprintf("Error when logging %v", err.Error()))
		}
	})

	t.Run("log when log level is Debug should call Debug", func(t *testing.T) {
		ml := NewMockLogger()
		lsvc := &LoggerService{
			logger: ml,
		}

		err := lsvc.Log(&LogRequest{
			LogLevel: "Debug",
		})

		if err != nil && ml.methodCalled && ml.methodName == "Debug" {
			t.Errorf(fmt.Sprintf("Error when logging %v", err.Error()))
		}
	})

	t.Run("log when log level does not exist should return error", func(t *testing.T) {
		ml := NewMockLogger()
		lsvc := &LoggerService{
			logger: ml,
		}

		err := lsvc.Log(&LogRequest{
			LogLevel: "Stuff",
		})

		if err == nil {
			t.Errorf("Should return error")
		}
	})
}
