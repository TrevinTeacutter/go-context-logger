package logs

import (
	"encoding/json"
	"fmt"
)

type Encoder interface {
	Encode(log Log) ([]byte, error)
}

var _ Encoder = (*JSONEncoder)(nil)

type JSONEncoder struct{}

func (encoder *JSONEncoder) Encode(log Log) ([]byte, error) {
	return json.Marshal(log)
}

type TextEncoder struct{}

func (encoder *TextEncoder) Encode(log Log) ([]byte, error) {
	return []byte(fmt.Sprintf("%s", log)), nil
}
