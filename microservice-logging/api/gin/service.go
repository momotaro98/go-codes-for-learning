package main

import (
	"context"

	"github.com/momotaro98/go-codes-for-learning/microservice-logging/logger"
)

type Service interface {
	SearchSample(ctx context.Context, sampleID string) (string, error)
}

type service struct{}

// NewService is ...
func NewService() Service {
	return newLoggingFilter(
		&service{},
		func() logger.Logger { return logger.Log },
	)
}

func (s *service) SearchSample(ctx context.Context, sampleID string) (string, error) {
	panic("SearchSample panic!")
	return "SearchSample result", nil
}
