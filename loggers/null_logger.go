package loggers

type NullLogger struct{}

func (nl *NullLogger) GetName() string {
	return "NullLogger"
}

func (nl *NullLogger) PlainInfo(message string) {
}
func (nl *NullLogger) Info(message string, platform string, boundary, name string, args ...interface{}) {
}
func (nl *NullLogger) Warning(message string, platform string, boundary, name string, args ...interface{}) {
}
func (nl *NullLogger) Error(message string, platform string, boundary, name string, args ...interface{}) {
}
func (nl *NullLogger) PlainError(message string) {
}
func (nl *NullLogger) Debug(message string, platform string, boundary, name string, args ...interface{}) {
}
func (nl *NullLogger) Close() error { return nil }
