package logs

import (
	"go.opentelemetry.io/otel/attribute"
)

type Log struct {
	Message    string
	Attributes []attribute.KeyValue
}
