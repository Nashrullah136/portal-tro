package logutils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"nashrul-be/crm/utils/localtime"
	"os"
)

var logger *log.Logger

func Init() error {
	logFile := fmt.Sprintf("logs-%s.logs", localtime.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	writer := io.MultiWriter(file)
	if os.Getenv("LOG") == "cli" {
		writer = io.MultiWriter(writer, os.Stdout)
	}
	logger = log.New(writer, "[INFO] ", log.LstdFlags|log.Llongfile)
	gin.DefaultWriter = writer
	return nil
}

func Get() *log.Logger {
	return logger
}
