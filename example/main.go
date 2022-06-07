package main

import (
	"errors"

	"github.com/trevinteacutter/go-context-logger/pkg/logs"
)

func main() {
	logger := logs.GetLogger()

	logger.Log("foo")

	logger.Error(errors.New("bar"))

	logger.Verbose("baz")
}
