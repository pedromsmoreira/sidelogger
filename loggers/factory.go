package loggers

import (
	"fmt"
)

func CreateLogger(name string) (Logger, error) {
	return &NullLogger{}, fmt.Errorf("logger with name %v is not implemented", name)
}
