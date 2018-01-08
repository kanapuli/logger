package logger

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//Logger holds the configuration properties of the log
type Logger struct {
	//FileName is the name of the output log file.
	//FileName should be a complete path without extension
	FileName string `json:"filename"`
	//TimeFormat specifies either 'utc' or 'local'
	//Timeformat to be appended in filename is RFC3399
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
	//filename is the filename appended with the time string
	filename string
}

//Logger implements all the methods of io.WriterCloser
var _ io.WriteCloser = (*Logger)(nil)

//Write writes the log to l.FileName
func (l *Logger) Write(data []byte) (int, error) {
	//Check if the logFile exists
	//Create the filename format
	_ = l.prepareFileName()
	if _, err := os.Stat(l.filename); os.IsNotExist(err) {
		//File does not exist. Create a new File
		file, err := os.Create(l.filename)
		if err != nil {
			//-1 from the Write Method indicates error
			return -1, err
		}
		//Assign the file object to the Logger File
		l.file = file

	} else {
		//file exists.
		//opent the file in a Writeonly or append mode and the permissions are set to write.
		l.file, err = os.OpenFile(l.filename, os.O_WRONLY|os.O_APPEND, 0644)
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

//prepareFileName prepares the name of the log file
//The Datetime has to be appended to the log file and the extension will be appended
func (l *Logger) prepareFileName() error {
	if strings.TrimSpace(l.TimeFormat) == "" {
		return errors.New("Time Format not set. FileName will be set as mentioned without log time information")
	}
	if l.TimeFormat == "local" {
		//Append local time
		l.filename = fmt.Sprintf("%s-%v.%s", l.FileName, time.Now().Format(time.RFC3339), "log")
		return nil
	}
	//Append utc time
	l.filename = fmt.Sprintf("%s-%v.%s", l.FileName, time.Now().UTC().Format(time.RFC3339), "log")
	return nil
}
