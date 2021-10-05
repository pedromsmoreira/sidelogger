package loggers

type Logger interface {
	PlainInfo(message string)
	Info(message string, platform string, boundary, name string, args ...interface{})
	Warning(message string, platform string, boundary, name string, args ...interface{})
	Error(message string, platform string, boundary, name string, args ...interface{})
	PlainError(message string)
	Debug(message string, platform string, boundary, name string, args ...interface{})
	Close() error
}
