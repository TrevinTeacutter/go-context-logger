package discard

import (
	"go.opentelemetry.io/otel/label"

	"github.com/trevinteacutter/go-context-logger/pkg/logs"
)

// Verify that it actually implements the interface
var _ logs.Logger = Logger{}

// Discard returns a valid Logger that discards all messages logged to it.
// It can be used whenever the caller is not interested in the logs.
func Discard() Logger {
	return Logger{}
}

// DiscardLogger is a Logger that discards all messages.
type Logger struct{}

func (l Logger) Error(_ error, _ string) {}

func (l Logger) Log(_ string) {}

func (l Logger) Verbose(_ string) {}

func (l Logger) ErrorWithLabels(_ error, _ string, _ ...label.KeyValue) {}

func (l Logger) LogWithLabels(_ string, _ ...label.KeyValue) {}

func (l Logger) VerboseWithLabels(_ string, _ ...label.KeyValue) {}

func (l Logger) WithBaseLabels(_ ...label.KeyValue) logs.Logger {
	return l
}

