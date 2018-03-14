package logger

import (
	"context"
	"io"
	"time"
)

//Logger represents the active logging properties
type Logger struct {
	Filename  string    `json:"filename"`
	Timestamp time.Time `json:"timestamp"`
	FileSize  int       `json:"filesize"` //Filesize is in Megabytes
	Filepath  string    `json:"filepath"`
}

// LogLevel is the integer referring to INFO,DEBUG etc
type LogLevel int

const (
	//INFO -information level (0)
	INFO = iota
	//DEBUG -debug level (1)
	DEBUG
	//WARNING -warning level (2)
	WARNING
	//ERROR -error level (3)
	ERROR
	//FATAL -fatal error level (4)
	FATAL
)

//Logger implements io.WriterCloser
var _ io.WriteCloser = (*Logger)(nil)

//Close closes the logfile
func (logger *Logger) Close() error {
	return nil
}

//Write writes to the log file
func (logger *Logger) Write(data []byte) (int, error) {
	return 0, nil
}

//Log is the external interface to get log data
func Log(ctx context.Context, level LogLevel, data interface{}) error {
	return nil
}
