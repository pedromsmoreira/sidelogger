package loggers

type LogRequest struct {
	LogLevel      string                 `json:"logLevel"`
	BuildId       string                 `json:"buildId"`
	Pipeline      string                 `json:"pipeline"`
	GitCommitHash string                 `json:"gitCommitHash"`
	Message       string                 `json:"message"`
	Platform      string                 `json:"platform"`
	Boundary      string                 `json:"boundary"`
	Name          string                 `json:"name"`
	Metadata      map[string]interface{} `json:"metadata"`
}
