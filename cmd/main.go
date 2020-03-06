package main

import (
	"github.com/sirupsen/logrus"

	logrusAdapter "github.com/trevinteacutter/go-context-logger/pkg/adapters/logrus"
	logger2 "github.com/trevinteacutter/go-context-logger/pkg/logger"
)

func main() {
	loggerConfig := logger2.Configuration{
		Format:       "prettyjson",
		Verbose:      true,
		OmitMetadata: true,
		Flatten:      true,
		Prefix:       "",
		Separator:    ".",
		MessageKey:   "message",
	}

	logrusAdapter.ConfigureAdapter(loggerConfig, logrus.StandardLogger())

	logger := logrusAdapter.CreateLogger()
	loggerContext := logger2.CreateRootContext(loggerConfig)

	specialContext := loggerContext.
		WithFields(map[logger2.Field]interface{}{
			logger2.CreateField("foo"): "bar",
		}).
		WithChildren(
			logger2.CreateContext("baz").WithFields(map[logger2.Field]interface{}{
				logger2.CreateField("haha"): "no",
			}))

	logger.Log("Hello World!")
	logger.LogWithContext("Hello Empty Context!", loggerContext)
	logger.LogWithContext("Hello Context!", specialContext)
}
