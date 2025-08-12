package loggerinterface

import "context"

type Logger interface {
	PrintInfo(ctx context.Context, msg string)
	PrintError(ctx context.Context, msg string, err error)
	Sync() error
}
