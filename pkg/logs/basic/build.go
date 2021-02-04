package basic

import (
	"github.com/trevinteacutter/go-context-logger/pkg/logs"
)

const (
	buildPrefix = "build"
)

var (
	// Build represents the version of code in the binary
	Build = "snapshot" // nolint:gochecknoglobals
	// Commit SHA of the code used in the binary
	Commit = "unknown" // nolint:gochecknoglobals
	// Date the binary was built
	Date = "unknown" // nolint:gochecknoglobals
	// Version of go this binary was built from
	Version = "unknown" // nolint:gochecknoglobals
	// Program is the name of the binary/program
	Program = "go-context-logger" // nolint:gochecknoglobals
	// OS is the OS the binary was built for
	OS = "unknown" // nolint:gochecknoglobals
	// Architecture represents the processor architecture the binary was built for
	Architecture = "unknown" // nolint:gochecknoglobals
	// ARM is only used for the ARM architecture, delineates the version of ARM the binary was built for
	ARM = "" // nolint:gochecknoglobals
)

func CreateBuildContext() *Context {
	context := CreateContext(buildPrefix)

	context.fields = []logs.Field{
		CreateField("program", Program),
		CreateField("version", Build),
		CreateField("build", Commit),
		CreateField("os", OS),
		CreateField("architecture", Architecture),
		CreateField("date", Date),
		CreateField("goVersion", Version),
	}

	// Ignore warnings about this, we set this at build time so editors won't recognize this might change
	if ARM != "" {
		context.fields = append(context.fields, CreateField("armVersion", ARM))
	}

	return context
}
