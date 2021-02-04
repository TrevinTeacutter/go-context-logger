package logs

type Field interface {
	Key() string
	Value() interface{}
	Verbose() bool
}
