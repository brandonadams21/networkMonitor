package logger

import (
    "log"
    "os"
)

func InitLogger(filepath string) *log.Logger {
    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %s", err)
    }

    return log.New(file, "", log.LstdFlags)
}
