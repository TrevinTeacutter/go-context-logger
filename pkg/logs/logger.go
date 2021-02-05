package logs

import (
	"go.opentelemetry.io/otel/label"
)

type Logger interface {
	// Error is equivalent to logging with error level
	Error(err error, message string)
	// Log is equivalent to logging with info level
	Log(message string)
	// Log is equivalent to logging with debug/verbose level
	Verbose(message string)
	// ErrorWithLabels is equivalent to logging with error level, but includes the given context
	ErrorWithLabels(err error, message string, labels ...label.KeyValue)
	// LogWithLabels is equivalent to logging with info level, but includes the given context
	LogWithLabels(message string, labels ...label.KeyValue)
	// VerboseWithLabels is equivalent to logging with debug/verbose level, but includes the given context
	VerboseWithLabels(message string, labels ...label.KeyValue)
	// WithBaseLabels
	WithBaseLabels(labels ...label.KeyValue) Logger
}
