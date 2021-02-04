package main

import (
	"github.com/sirupsen/logrus"

	logrusAdapter "github.com/trevinteacutter/go-context-logger/pkg/adapters/logrus"
	"github.com/trevinteacutter/go-context-logger/pkg/logs"
	"github.com/trevinteacutter/go-context-logger/pkg/logs/basic"
)

type Awesome struct {
	logger logs.Logger
}

func NewAwesome(logger logs.Logger) *Awesome {
	ctx := basic.CreateContext("logger").WithFields(
		basic.CreateField("struct", "Awesome"),
	)

	return &Awesome{
		logger.WithBaseContext(ctx),
	}
}

func (a *Awesome) DoStuff(param string, ctx logs.Context) {
	newContext := ctx.WithChildren(
		basic.CreateContext("DoStuff").WithFields(
			basic.CreateField("foo", "bar"),
			basic.CreateVerboseField("baz", "baz"),
		),
	)

	a.logger.LogWithContext(param, newContext)
}

func main() {
	configuration := logs.Configuration{
		Format:       "prettyjson",
		Verbose:      true,
		OmitMetadata: true,
		Flatten:      true,
		Prefix:       "program",
		Separator:    ".",
		MessageKey:   "message",
	}

	root := basic.CreateRootContext(configuration)
	baseLogger := logrusAdapter.CreateLogger(logrus.New()).WithConfiguration(configuration)
	thing := NewAwesome(baseLogger)

	thing.DoStuff("hahahahahaha", root)
}
