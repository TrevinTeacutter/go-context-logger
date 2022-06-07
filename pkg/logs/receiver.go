package logs

import (
	"io"
)

type Receiver interface {
	Receive(log Log) error
}

var (
	_ Receiver = (*NoopReceiver)(nil)
	_ Receiver = (*StandardReceiver)(nil)
)

type NoopReceiver struct{}

func (r *NoopReceiver) Receive(_ Log) error {
	return nil
}

type StandardReceiver struct {
	Encoder   Encoder
	Writer    io.Writer
	Delimiter byte
}

func (r *StandardReceiver) Receive(log Log) error {
	rawLog, err := r.Encoder.Encode(log)
	if err != nil {
		return err
	}

	_, err = r.Writer.Write(append(rawLog, r.Delimiter))
	if err != nil {
		return err
	}

	return nil
}
