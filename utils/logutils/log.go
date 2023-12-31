package logutils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"nashrul-be/crm/utils/localtime"
	"os"
	"path/filepath"
)

var logger *log.Logger

func Init(logPath string) error {
	writer, err := GetWriter(logPath)
	if err != nil {
		return err
	}
	logger = log.New(writer, "[INFO] ", log.LstdFlags|log.Llongfile)
	gin.DefaultWriter = writer
	return nil
}

func GetWriter(logPath string) (io.Writer, error) {
	logFile := fmt.Sprintf("logs-%s.logs", localtime.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filepath.Join(logPath, logFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	writer := io.MultiWriter(file)
	if os.Getenv("LOG") == "cli" {
		writer = io.MultiWriter(writer, os.Stdout)
	}
	return writer, nil
}

func CliOnly() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

func Get() *log.Logger {
	return logger
}
