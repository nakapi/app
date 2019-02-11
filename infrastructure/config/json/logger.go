package json

type LoggerConfig struct {
	Level         string `json:"level"`
	Encoding      string `json:"encoding"`
	EncoderConfig struct {
		MessageKey      string `json:"messageKey`
		LevelKey        string `json:"levelKey`
		TimeKey         string `json:"timeKey`
		NameKey         string `json:"nameKey`
		callerKey       string `json:"callerKey`
		StackTraceKey   string `json:"stacktraceKey`
		LevelEncoder    string `json:"levelEncoder"`
		TimeEncoder     string `json:"timeEncoder"`
		DurationEncoder string `json:"durationEncoder"`
		CallerEncoder   string `json:"callerEncoder"`
	}
	OutputPaths      []string `json:"outputPaths"`
	ErrorOutputPaths []string `json:"stderr`
}
