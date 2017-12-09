package logger

import (
	"errors"
	"io"
	"sync"
	"time"
)

//Make sure Logger implements io.WriteCLoser
var _ io.WriteCloser = (*Logger)(nil)

//Logger is an io.WriteCloser that writes to a file
type Logger struct {
	//FileName is the file to write logs.Files will be written in the same directory
	FileName string `json:"filename"`
	//LocalTime value decides if the file to be written using system's local time
	LocalTime bool `json:"localtime"`
	mu        sync.Mutex
}

var (
	currentTime = time.Now()
)

//Write implements the io.Writer
func (l *Logger) Write(data []byte) (int, error) {
	return 1, errors.New("Not Implemented")
}

//Close implements Closer
func (l *Logger) Close() error {
	return errors.New("Not Implemented")
}
