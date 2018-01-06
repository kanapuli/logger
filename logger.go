package logger

import (
	"errors"
	"io"
	"os"
)

//Logger holds the configuration properties of the log
type Logger struct {
	//FileName is the name of the output log file.
	//FileName should be a complete path
	FileName string `json:"filename"`
	//TimeFormat specifies either 'utc' or 'local'
	TimeFormat string `json:"timeformat"`
	//Maxsize is the maximum size the log file can grow. Specified in kb
	MaxSize int `json:"maxsize"`
	//ShouldArchive = true archives the log file if it reaches the max size or in case if DailyLog is enabled, then the file is archived on every end of the day else the file remains unzipped
	ShouldArchive bool `json:"shouldarchive"`
	//DailyLog = true creates a new log file everyday or if false the same file is used until it reaches the max size
	DailyLog bool `json:"dailylog"`
	//PushToS3 = true pushes the archived log file to S3 everyday
	PushToS3 bool `json:"pushtos3"`
}

//Logger implements all the methods of io.WriterCloser
var _ io.WriteCloser = (*Logger)(nil)

//Write writes the log to l.FileName
func (l *Logger) Write([]byte) (int, error) {
	//Check if the logFile esists
	if _, err := os.Stat(l.FileName); os.IsNotExist(err) {
		//File does not exist. Create a new File
		_, err := os.Create(l.FileName)
		if err != nil {
			//-1 from the Write Method indicates error
			return -1, err
		}
	}

	return 0, errors.New("Partially Implemented")
}

//Close closes the log file which is opened to write
func (l *Logger) Close() error {
	return errors.New("Not Implemented")
}
