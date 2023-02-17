package harvest

type LogFields map[string]interface{}

type Logger interface {
	WithFields(fields LogFields) Logger
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}
