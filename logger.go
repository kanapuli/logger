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

	file *os.File
}

//Logger implements all the methods of io.WriterCloser
var _ io.WriteCloser = (*Logger)(nil)

//Write writes the log to l.FileName
func (l *Logger) Write(data []byte) (int, error) {
	//Check if the logFile exists
	if _, err := os.Stat(l.FileName); os.IsNotExist(err) {
		//File does not exist. Create a new File
		file, err := os.Create(l.FileName)
		if err != nil {
			//-1 from the Write Method indicates error
			return -1, err
		}
		//Assign the file object to the Logger File
		l.file = file

	} else {
		//file exists.
		//opent the file in a Writeonly or append mode and the permissions are set to write.
		l.file, err = os.OpenFile(l.FileName, os.O_WRONLY|os.O_APPEND, 444)
		if err != nil {
			return -1, err
		}
	}

	//Write data to the file
	n, err := l.file.Write(data)
	//close the file
	err = l.Close()
	if err != nil {
		return -1, err
	}

	return n, err
}

//Close closes the log file which is opened to write
func (l *Logger) Close() error {
	if l.file != nil {
		return close(l)
	}

	return errors.New("File is already nil")
}

//close calls the *file.Close() method.
func close(l *Logger) error {
	if l.file == nil {
		return errors.New("Invalid File")
	}
	var err error
	err = l.file.Close()
	l.file = nil
	return err
}
