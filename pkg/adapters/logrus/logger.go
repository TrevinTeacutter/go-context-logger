package logrus

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/trevinteacutter/go-context-logger/pkg/logs"
)

var _ logs.Logger = (*Logger)(nil)

type Logger struct {
	logger  *logrus.Entry
	configuration logs.Configuration
	context logs.Context
}

func CreateLogger(baseLogger *logrus.Logger) *Logger {
	baseLogger.SetOutput(os.Stdout)
	baseLogger.SetReportCaller(true)

	return &Logger{
		logger: logrus.NewEntry(baseLogger),
	}

}

func (logger *Logger) Error(err error, message string) {
	logger.logger.WithFields(logger.loggerFields()).WithError(err).Error(message)
}

func (logger *Logger) Log(message string) {
	logger.logger.WithFields(logger.loggerFields()).Info(message)
}

func (logger *Logger) Verbose(message string) {
	logger.logger.WithFields(logger.loggerFields()).Trace(message)
}

func (logger *Logger) ErrorWithContext(err error, message string, ctx logs.Context) {
	logger.logger.WithFields(logger.loggerFields()).WithFields(logger.contextFields(ctx)).WithError(err).Error(message)
}

func (logger *Logger) LogWithContext(message string, ctx logs.Context) {
	logger.logger.WithFields(logger.loggerFields()).WithFields(logger.contextFields(ctx)).Info(message)
}

func (logger *Logger) VerboseWithContext(message string, ctx logs.Context) {
	logger.logger.WithFields(logger.loggerFields()).WithFields(logger.contextFields(ctx)).Trace(message)
}

func (logger *Logger) loggerFields() map[string]interface{} {
	return logger.context.Fields(logger.configuration.Verbose, logger.configuration.Flatten, logger.configuration.Separator)
}

func (logger *Logger) contextFields(ctx logs.Context) map[string]interface{} {
	return ctx.Fields(logger.configuration.Verbose, logger.configuration.Flatten, logger.configuration.Separator)
}

func (logger *Logger) WithConfiguration(configuration logs.Configuration) logs.Logger {
	newLogger := CreateLogger(logger.logger.Logger)

	newLogger.configuration = configuration

	switch configuration.Verbose {
	case true:
		newLogger.logger.Logger.SetLevel(logrus.InfoLevel)
	default:
		newLogger.logger.Logger.SetLevel(logrus.TraceLevel)
	}

	switch configuration.Format {
	case "text":
		newLogger.logger.Logger.SetFormatter(&logrus.TextFormatter{
			DisableLevelTruncation: true,
			QuoteEmptyFields:       true,
		})
	case "prettyjson":
		newLogger.logger.Logger.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg:  configuration.MessageKey,
				logrus.FieldKeyFunc: "function",
			},
			PrettyPrint: true,
		})
	default:
		newLogger.logger.Logger.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg:  configuration.MessageKey,
				logrus.FieldKeyFunc: "function",
			},
		})
	}

	return newLogger
}

func (logger *Logger) WithBaseContext(context logs.Context) logs.Logger {
	newLogger := CreateLogger(logger.logger.Logger)

	newLogger.context = context

	return newLogger
}
