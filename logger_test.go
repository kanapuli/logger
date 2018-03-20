package logger_test

import (
	"context"
	"os"
	"testing"

	"github.com/kanapuliAthavan/logger"
)

// test file creation
func Test_CreateFile(t *testing.T) {

	log := logger.Logger{
		Filename: "newlog.log",
		FileSize: 1,
		Filepath: "",
	}
	ctx := context.Background()
	log.Log(ctx, logger.INFO, "Write some initial log")
	_, err := os.Stat(log.Filepath + log.Filename)
	if os.IsNotExist(err) {
		t.Errorf("error creating log file: %v", err)
	}

}

//test file writing
func Test_WriteFile(t *testing.T) {

	log := logger.Logger{
		Filename: "newlog.log",
		FileSize: 1,
		Filepath: "",
	}
	ctx := context.Background()
	err := log.Log(ctx, logger.INFO, "Check if log is written")
	if err != nil {
		t.Errorf("error writing to file: %v", err)
	}
}
