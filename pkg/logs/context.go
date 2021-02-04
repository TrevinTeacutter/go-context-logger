package logs

type Context interface {
	Prefix() string
	Fields(verbose, flatten bool, separator string) map[string]interface{}
	Copy() Context
	WithFields(fields ...Field) Context
	WithChildren(children ...Context) Context
}
