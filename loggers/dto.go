package loggers

type LogRequest struct {
	LogLevel string                 `json:"logLevel"`
	Message  string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata"`
}
