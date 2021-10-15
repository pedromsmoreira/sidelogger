package loggers

type MockLogger struct {
	methodCalled bool
	methodName   string
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		methodCalled: false,
		methodName:   "",
	}
}

func (ml *MockLogger) GetName() string {
	return "MockLogger"
}

func (ml *MockLogger) PlainInfo(message string) {
	ml.methodCalled = true
	ml.methodName = "PlainInfo"
}

func (ml *MockLogger) Info(message string, platform string, boundary, name string, args ...interface{}) {
	ml.methodCalled = true
	ml.methodName = "Info"
}

func (ml *MockLogger) Warning(message string, platform string, boundary, name string, args ...interface{}) {
	ml.methodCalled = true
	ml.methodName = "Warning"
}

func (ml *MockLogger) Error(message string, platform string, boundary, name string, args ...interface{}) {
	ml.methodCalled = true
	ml.methodName = "Error"
}

func (ml *MockLogger) PlainError(message string) {
	ml.methodCalled = true
	ml.methodName = "PlainError"
}

func (ml *MockLogger) Debug(message string, platform string, boundary, name string, args ...interface{}) {
	ml.methodCalled = true
	ml.methodName = "Debug"
}

func (ml *MockLogger) Close() error {
	ml.methodCalled = true
	ml.methodName = "Close"
	return nil
}
