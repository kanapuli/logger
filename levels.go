package logger

type Level int

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)
