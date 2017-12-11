package logger

import (
	"errors"
	"fmt"
	"io"
	"os"
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
	file      *os.File
	size      int64
}

var (
	currentTime = time.Now()
)

//Write implements the io.Writer. Write checks if the dataToWrite is greater than the max filesize of logger.
func (l *Logger) Write(data []byte) (int, error) {
	lengthToWrite := int64(len(data))
	if lengthToWrite > l.size {
		return -1, fmt.Errorf("The file size %d is greater than the max logger size %d", lengthToWrite, l.size)
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	bytesWritten, err := l.file.Write(data)
	l.size += int64(bytesWritten)
	return bytesWritten, err
}

//Close implements Closer
func (l *Logger) Close() error {
	return errors.New("Not Implemented")
}
