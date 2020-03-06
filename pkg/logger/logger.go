package logger

type Logger interface {
	// Error is equivalent to logging with error level
	Error(message interface{})
	// Log is equivalent to logging with info level
	Log(message interface{})
	// Log is equivalent to logging with debug/trace level
	Trace(message interface{})
	// ErrorWithContext is equivalent to logging with error level, but includes the given context
	ErrorWithContext(message interface{}, context *Context)
	// LogWithContext is equivalent to logging with info level, but includes the given context
	LogWithContext(message interface{}, context *Context)
	// TraceWithContext is equivalent to logging with debug/trace level, but includes the given context
	TraceWithContext(message interface{}, context *Context)
}
