package logger

import (
	"context"
	"log"
)

type StdLogger struct{}

func InitLogger() *StdLogger {
	return &StdLogger{}
}

func (l *StdLogger) PrintInfo(ctx context.Context, message string) {
	log.Printf("INFO работает СТАНДАРТНЫЙ логгер %s\n", message)
}

func (l *StdLogger) PrintError(ctx context.Context, message string, err error) {
	log.Printf("ERROR работает СТАНДАРТНЫЙ логгер %v\n", err)
}

func (l *StdLogger) Sync() error {
	return nil
}
