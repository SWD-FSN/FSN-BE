package util

import (
	"log"
	"os"
)

func GetLogConfig() *log.Logger {
	return log.New(os.Stdout, "[ERROR] ", log.LstdFlags)
}
