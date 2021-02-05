package main

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/label"

	logrusAdapter "github.com/trevinteacutter/go-context-logger/pkg/adapters/logrus"
	"github.com/trevinteacutter/go-context-logger/pkg/logs"
)

type Awesome struct {
	logger logs.Logger
}

func NewAwesome(logger logs.Logger) *Awesome {
	return &Awesome{
		logger.WithBaseLabels(label.Any("struct", "Awesome")),
	}
}

func (a *Awesome) DoStuff(param string, labels ...label.KeyValue) {
	a.logger.LogWithLabels(param, label.Any("foo", "bar"), label.Any("baz", "baz"))
}

func main() {
	configuration := logrusAdapter.Configuration{
		Format:       "prettyjson",
		Verbose:      true,
		Prefix:       "program",
		MessageKey:   "message",
	}

	rootLabels := []label.KeyValue{label.Any("program", "meowmix")}
	baseLogger := logrusAdapter.CreateLogger(logrus.New()).WithConfiguration(configuration)
	thing := NewAwesome(baseLogger)

	thing.DoStuff("hahahahahaha", rootLabels...)
}
