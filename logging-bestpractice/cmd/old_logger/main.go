package main

import (
	"context"

	"git.rarejob.com/shintaro.ikeda/platform_logging/logger"
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

func (l *logging) info(ctx context.Context, msg string, appendFields ...logger.Field) {
	var fields = make(logger.Fields, len(appendFields)+5)

	fields = append(fields, appendFields...)

	fields = append(fields, normalLogFields(ctx)...)

	l.loggerFunc().Info(msg, fields...)
}

func normalLogFields(ctx context.Context) logger.Fields {
	return logger.Fields{
		logger.F("transaction", "Dummy-transaction"),
	}
}

func main() {
	s := NewService()
	ctx := context.Background()
	s.info(ctx, "AAA")
}
