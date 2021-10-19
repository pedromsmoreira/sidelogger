package loggers

import (
	"fmt"
)

type LoggerService struct {
	logger Logger
}

func NewLoggerService(logger Logger) *LoggerService {
	return &LoggerService{
		logger: logger,
	}
}

func (ls *LoggerService) Log(request *LogRequest) error {
	switch fmt.Sprintf("%v", request.LogLevel) {
	case "Info":
		ls.logger.Info(request.Message, request.Metadata)
		return nil
	case "Warning":
		ls.logger.Warning(request.Message, request.Metadata)
		return nil
	case "Error":
		ls.logger.Error(request.Message, request.Metadata)
		return nil
	case "Debug":
		ls.logger.Debug(request.Message, request.Metadata)
		return nil
	default:
		return fmt.Errorf("LogLevel %v is not available", request.LogLevel)
	}
}

func (ls *LoggerService) SimpleLog(message string) error {
	ls.logger.PlainInfo("Log Sucess")
	return nil
}
