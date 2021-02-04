package basic

import (
	"github.com/trevinteacutter/go-context-logger/pkg/logs"
)

var _ logs.Field = (*Field)(nil)

type Field struct {
	key     string
	value   interface{}
	verbose bool
}

// CreateVerboseField is a helper to create a Field that is only available in verbose mode
func CreateVerboseField(key string, value interface{}) *Field {
	return createField(key, value, true)
}

// CreateField is a helper to create a Field that is always available
func CreateField(key string, value interface{}) *Field {
	return createField(key, value, false)
}

func createField(key string, value interface{}, verbose bool) *Field {
	return &Field{
		key:  key,
		value:  value,
		verbose: verbose,
	}
}

func (field *Field) Key() string {
	return field.key
}

func (field *Field) Value() interface{} {
	return field.value
}

func (field *Field) Verbose() bool {
	return field.verbose
}
