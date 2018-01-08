package logger_test

import (
	"os"
	"testing"

	"github.com/kanapuliAthavan/logger"
)

func TestWriteNewLog(t *testing.T) {
	//remove file if exists
	os.Remove("Errorlog.log")
	l := logger.Logger{
		FileName:      "Errorlog.log", //file path is pwd. Filename without extension
		MaxSize:       200,
		DailyLog:      true,
		PushToS3:      false,
		ShouldArchive: false,
		TimeFormat:    "local",
	}
	n, err := l.Write([]byte("10 bytes\n"))
	if err != nil {
		t.Errorf("Error occured in writing log file %v", err)
	}
	if n != 9 {
		//n is the amount of bytes written
		t.Errorf("Expected 8 bytes written but got %d", n)
	}
}

func TestWriteOldLog(t *testing.T) {
	l := logger.Logger{
		FileName:      "Errorlog.log", //file path is pwd
		MaxSize:       200,
		DailyLog:      true,
		PushToS3:      false,
		ShouldArchive: false,
		TimeFormat:    "utc",
	}
	n, err := l.Write([]byte("10 bytes\n"))
	if err != nil {
		t.Errorf("Error occured in writing log file %v", err)
	}
	if n != 9 {
		//n is the amount of bytes written
		t.Errorf("Expected 8 bytes written but got %d", n)
	}
}
func TestWriteInvalidFile(t *testing.T) {
	l := new(logger.Logger)
	n, err := l.Write([]byte("10 bytes\n"))
	if err == nil {
		t.Errorf("Expected Error in writing file but got %v", err)
	}
	if n == 9 {
		//n is the amount of bytes written
		t.Errorf("Expected 0 bytes written but got %d", n)
	}
}

func TestFileClose(t *testing.T) {
	l := new(logger.Logger)
	err := l.Close()
	if err == nil {
		t.Errorf("Expected Error in CLosing File but got %v", err)
	}
}
