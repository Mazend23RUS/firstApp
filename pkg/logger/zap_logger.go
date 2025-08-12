package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type ZapLogger struct {
	zaplog *zap.Logger
}

func NewZapLogger() (*ZapLogger, error) {

	loggerzap, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("ошибка создания zap логгера %v", err)
	}
	loggerzap.Info("создался логгер zap !!!")
	return &ZapLogger{
		zaplog: loggerzap,
	}, nil
}

func (z *ZapLogger) PrintInfo(ctx context.Context, msg string) {
	z.zaplog.Info("INFO работает логгер zap: " + msg + "\n")

}

func (z *ZapLogger) PrintError(ctx context.Context, msg string, err error) {
	z.zaplog.Error("ERROR работает логгер zap: " + msg + "\n")
}

func (z *ZapLogger) Sync() error {
	return z.Sync()
}
