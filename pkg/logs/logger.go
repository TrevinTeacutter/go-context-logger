package logs

type Logger interface {
	// Error is equivalent to logging with error level
	Error(err error, message string)
	// Log is equivalent to logging with info level
	Log(message string)
	// Log is equivalent to logging with debug/verbose level
	Verbose(message string)
	// ErrorWithContext is equivalent to logging with error level, but includes the given context
	ErrorWithContext(err error, message string, context Context)
	// LogWithContext is equivalent to logging with info level, but includes the given context
	LogWithContext(message string, context Context)
	// TraceWithContext is equivalent to logging with debug/verbose level, but includes the given context
	VerboseWithContext(message string, context Context)
	// WithConfiguration
	WithConfiguration(configuration Configuration) Logger
	// WithBaseContext
	WithBaseContext(context Context) Logger
}
