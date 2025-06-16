package loggerinterface

import "context"

type Logger interface {
	Info(ctx context.Context, msg string)
	Error(ctx context.Context, msg string, err error)
}
