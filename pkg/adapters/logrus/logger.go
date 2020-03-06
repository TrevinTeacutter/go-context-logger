package logrus

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/trevinteacutter/go-context-logger/pkg/logger"
)

var (
	// nolint: gochecknoglobals
	globalLogger = logrus.New()
)

// ConfigureAdapter is meant to setup the logrus logging adapter, however it only ever be called once for an application
func ConfigureAdapter(configuration logger.Configuration, baseLogger *logrus.Logger) {
	switch baseLogger {
	case nil:
		globalLogger = logrus.StandardLogger()
	default:
		globalLogger = baseLogger
	}

	globalLogger.SetOutput(os.Stdout)
	globalLogger.SetReportCaller(true)

	switch configuration.Verbose {
	case true:
		globalLogger.SetLevel(logrus.InfoLevel)
	default:
		globalLogger.SetLevel(logrus.TraceLevel)
	}

	switch configuration.Format {
	case "text":
		globalLogger.SetFormatter(&logrus.TextFormatter{
			DisableLevelTruncation: true,
			QuoteEmptyFields:       true,
		})
	case "prettyjson":
		globalLogger.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg:  configuration.MessageKey,
				logrus.FieldKeyFunc: "function",
			},
			PrettyPrint: true,
		})
	default:
		globalLogger.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg:  configuration.MessageKey,
				logrus.FieldKeyFunc: "function",
			},
		})
	}
}

type Logger struct {
	logger *logrus.Entry
}

func CreateLogger() *Logger {
	return &Logger{
		logger: logrus.NewEntry(globalLogger),
	}
}

func (logger *Logger) Error(message interface{}) {
	logger.logger.Error(message)
}

func (logger *Logger) Log(message interface{}) {
	logger.logger.Info(message)
}

func (logger *Logger) Trace(message interface{}) {
	logger.logger.Trace(message)
}

func (logger *Logger) ErrorWithContext(message interface{}, ctx *logger.Context) {
	fields := ctx.GetFields()
	logger.logger.WithFields(fields).Error(message)
}

func (logger *Logger) LogWithContext(message interface{}, ctx *logger.Context) {
	fields := ctx.GetFields()
	logger.logger.WithFields(fields).Info(message)
}

func (logger *Logger) TraceWithContext(message interface{}, ctx *logger.Context) {
	fields := ctx.GetFields()
	logger.logger.WithFields(fields).Trace(message)
}
