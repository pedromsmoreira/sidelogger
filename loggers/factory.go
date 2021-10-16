package loggers

import (
	"fmt"
)

type LoggerFactory struct{}

func (lf *LoggerFactory) CreateLogger(name string) (Logger, error) {
	switch name {
	case "uberzap":
		return NewZapLogger()
	case "logrus":
		return NewLogrusLogger()
	case "apex":
		return NewApexLogger()
	default:
		return &NullLogger{}, fmt.Errorf("logger with name %v is not implemented", name)
	}
}
