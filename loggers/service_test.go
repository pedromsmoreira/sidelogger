package loggers

import (
	"fmt"
	"testing"
)

func TestLogWhenLogLevelIsInfoShouldCallInfo(t *testing.T) {
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
}

func TestLogWhenLogLevelIsWarningShouldCallWarning(t *testing.T) {
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
}

func TestLogWhenLogLevelIsErrorShouldCallError(t *testing.T) {
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
}

func TestLogWhenLogLevelIsDebugShouldCallDebug(t *testing.T) {
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
}

func TestLogWhenLogLevelDoesNotExistShouldReturnError(t *testing.T) {
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
}
