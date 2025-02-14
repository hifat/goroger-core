package core

type LoggerField struct {
	Key       string
	Type      uint8
	Integer   int64
	String    string
	Interface any
}

type Logger interface {
	Skip(skip int) Logger
	Info(message string, fields ...LoggerField)
	Debug(message string, fields ...LoggerField)
	Warn(message string, fields ...LoggerField)
	Fatal(message string, fields ...LoggerField)
	Error(message any, fields ...LoggerField)
}
