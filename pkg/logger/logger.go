package logger

import (
	"context"
	"log"
)

type Logger struct{}

func InitLogger() Logger {
	return Logger{}
}

func (l Logger) PtintInfo(ctx context.Context, message string) {
	log.Printf("INFO %s\n", message)

}

func (l Logger) PrintError(ctx context.Context, message string, err error) {

	log.Printf("ERROR %v\n", err)
}
