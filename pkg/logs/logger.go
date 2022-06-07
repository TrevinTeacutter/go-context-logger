package logs

import (
	"os"
	"sync"

	"go.opentelemetry.io/otel/attribute"
)

const (
	MetaFieldKeyPrefix = "meta"
	NameFieldKey       = MetaFieldKeyPrefix + ".name"
	PackageFieldKey    = MetaFieldKeyPrefix + ".package"
	TypeFieldKey       = MetaFieldKeyPrefix + ".type"
)

var (
	initializer sync.Once

	singleton *Logger
)

func GetLogger() *Logger {
	initializer.Do(func() {
		singleton = &Logger{
			root: Route{
				Children: []Route{
					{
						Matchers: []Matcher{
							&AttributeMatcher{
								Key:      "level",
								Operator: "NOT_EQUALS",
								Value:    attribute.StringValue("debug"),
							},
						},
						Receiver: &StandardReceiver{
							Encoder:   &TextEncoder{},
							Writer:    os.Stdout,
							Delimiter: '\n',
						},
					},
				},
				Receiver: &NoopReceiver{},
			},
		}
	})

	return singleton
}

type Logger struct {
	root Route
}

func (l *Logger) Error(err error, attributes ...attribute.KeyValue) {
	log := Log{
		Message:    err.Error(),
		Attributes: append(attributes, attribute.String("level", "error")),
	}
	receivers := l.root.Route(log)

	for _, receiver := range receivers {
		_ = receiver.Receive(log)
	}
}

func (l *Logger) Log(message string, attributes ...attribute.KeyValue) {
	log := Log{
		Message:    message,
		Attributes: append(attributes, attribute.String("level", "info")),
	}
	receivers := l.root.Route(log)

	for _, receiver := range receivers {
		_ = receiver.Receive(log)
	}
}

func (l *Logger) Verbose(message string, attributes ...attribute.KeyValue) {
	log := Log{
		Message:    message,
		Attributes: append(attributes, attribute.String("level", "debug")),
	}
	receivers := l.root.Route(log)

	for _, receiver := range receivers {
		_ = receiver.Receive(log)
	}
}
