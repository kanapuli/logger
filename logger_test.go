package logger_test

import (
	"context"
	"os"
	"testing"

	"github.com/kanapuliAthavan/logger"
)

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
		t.Errorf("Error creating log file: %v", err)
	}

}
