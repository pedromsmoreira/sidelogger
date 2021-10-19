package loggers

type Logger interface {
	GetName() string
	PlainInfo(message string)
	Info(message string, args ...interface{})
	Warning(message string, args ...interface{})
	Error(message string, args ...interface{})
	PlainError(message string)
	Debug(message string, args ...interface{})
	Close() error
}
