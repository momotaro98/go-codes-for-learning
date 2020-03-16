package main

import (
	"context"

	"github.com/momotaro98/go-codes-for-learning/microservice-logging/logger"
)

type logging struct {
	next       Service
	loggerFunc func() logger.Logger
}

// info is the application's common logging function
func (l *logging) info(ctx context.Context, msg string, fields ...logger.Field) {
	l.loggerFunc().Info(ctx.Value(transactionCtxKey), msg, fields...)
}

func newLoggingFilter(next Service, loggerFunc func() logger.Logger) Service {
	return &logging{
		next:       next,
		loggerFunc: loggerFunc,
	}
}

func (l *logging) SearchSample(ctx context.Context, sampleID string) (string, error) {
	out, err := l.next.SearchSample(ctx, sampleID)
	if err == nil {
		l.info(ctx, "search sample success",
			logger.F("sample-id", out),
		)
	}
	return out, err
}
