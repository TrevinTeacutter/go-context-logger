package logrus

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/label"

	"github.com/trevinteacutter/go-context-logger/pkg/logs"
)

var _ logs.Logger = (*Logger)(nil)

type Logger struct {
	logger  *logrus.Entry
	configuration Configuration
	labels []label.KeyValue
}

func CreateLogger(baseLogger *logrus.Logger) *Logger {
	baseLogger.SetOutput(os.Stdout)
	baseLogger.SetReportCaller(true)

	return &Logger{
		logger: logrus.NewEntry(baseLogger),
	}

}

func (logger *Logger) Error(err error, message string) {
	labelsToFields(logger.logger, logger.labels...).WithError(err).Error(message)
}

func (logger *Logger) Log(message string) {
	labelsToFields(logger.logger, logger.labels...).Info(message)
}

func (logger *Logger) Verbose(message string) {
	labelsToFields(logger.logger, logger.labels...).Debug(message)
}

func (logger *Logger) ErrorWithLabels(err error, message string, labels ...label.KeyValue) {
	temp := labelsToFields(logger.logger, logger.labels...)

	temp = labelsToFields(temp, labels...)

	temp.WithError(err).Error(message)
}

func (logger *Logger) LogWithLabels(message string, labels ...label.KeyValue) {
	temp := labelsToFields(logger.logger, logger.labels...)

	temp = labelsToFields(temp, labels...)

	temp.Info(message)
}

func (logger *Logger) VerboseWithLabels(message string, labels ...label.KeyValue) {
	temp := labelsToFields(logger.logger, logger.labels...)

	temp = labelsToFields(temp, labels...)

	temp.Debug(message)
}

func labelsToFields(logger *logrus.Entry, labels ...label.KeyValue) *logrus.Entry {
	for _, field := range labels {
		logger = logger.WithField(string(field.Key), field.Value.Emit())
	}

	return logger
}

func (logger *Logger) WithConfiguration(configuration Configuration) logs.Logger {
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

func (logger *Logger) WithBaseLabels(labels ...label.KeyValue) logs.Logger {
	newLogger := CreateLogger(logger.logger.Logger)

	newLogger.labels = append(newLogger.labels, labels...)

	return newLogger
}
