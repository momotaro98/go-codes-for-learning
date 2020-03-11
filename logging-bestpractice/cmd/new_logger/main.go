package main

import (
	"context"

	logger "git.rarejob.com/shintaro.ikeda/platform_logging/new_logger"
)

type Service interface {
}

type service struct{}

type logging struct {
	next       Service
	loggerFunc func() logger.Logger
}

func newLoggingFilter(next Service, loggerFunc func() logger.Logger) *logging {
	return &logging{
		next:       next,
		loggerFunc: loggerFunc,
	}
}

// NewService is ...
func NewService() *logging {
	return newLoggingFilter(
		&service{},
		logger.DefaultLogger,
	)
}

// info is the application's common logging function
func (l *logging) info(ctx context.Context, msg string, fields ...logger.Field) {
	l.loggerFunc().Info(ctx.Value(XTxIDCtxKey), ServiceName, msg, fields...)
}

const (
	XTxIDCtxKey = "xKey"
	ServiceName = "event-management-2"
)

func main() {
	s := NewService()
	ctx := context.Background()
	ctx = context.WithValue(ctx, XTxIDCtxKey, "abcdefg")
	fields := []logger.Field{
		logger.F("app-specific-key", "App-Specific-Val"),
	}
	s.info(ctx, "Application message", fields...)
}
