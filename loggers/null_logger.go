package loggers

type NullLogger struct{}

func (nl *NullLogger) info(message string, platform string, boundary, name string, args ...interface{}) {
}
func (nl *NullLogger) warning(message string, platform string, boundary, name string, args ...interface{}) {
}
func (nl *NullLogger) error(message string, platform string, boundary, name string, args ...interface{}) {
}
func (nl *NullLogger) debug(message string, platform string, boundary, name string, args ...interface{}) {
}
