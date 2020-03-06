package logger

type Field struct {
	name  string
	trace bool
}

// CreateTraceField is a helper to create a Field that is only available in verbose mode
func CreateTraceField(name string) Field {
	return createField(name, true)
}

// CreateField is a helper to create a Field that is always available
func CreateField(name string) Field {
	return createField(name, false)
}

func createField(name string, trace bool) Field {
	return Field{
		name:  name,
		trace: trace,
	}
}

func (field *Field) shouldBeLogged(verbose bool) bool {
	return field.trace && !verbose
}
