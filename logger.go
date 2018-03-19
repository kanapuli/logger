package logger

import (
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

//Logger represents the active logging properties
type Logger struct {
	Filename  string    `json:"filename"`  //FileName refers to Log File name
	Timestamp time.Time `json:"timestamp"` //Timestamp is the epoch timestamp when the log is received
	FileSize  int       `json:"filesize"`  //FileSize is the size of the file in Megabytes
	Filepath  string    `json:"filepath"`  //FilePath is the path of the log file
	file      *os.File
	mu        sync.Mutex
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
	err := logger.close()
	if err != nil {
		log.Printf("Error closing log file: %v", err)

	}
}

func (logger *Logger) close() error {
	if logger.file != nil {
		logger.mu.Lock()
		logger.file.Close()
		logger.mu.Unlock()

	}
}

//Write writes to the log file
func (logger *Logger) Write(data []byte) (int, error) {
	if logger.file == nil {
		logger.createLogFile()
	}
	return 0, nil
}

//Log is the external interface to get log data
func Log(ctx context.Context, level LogLevel, data interface{}) error {
	return nil
}

func (logger *Logger) createLogFile() {
	path := logger.Filepath + logger.Filename
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Fatalf("Error in creating log file: %v", err)
		}
		logger.file = file
	}

}
