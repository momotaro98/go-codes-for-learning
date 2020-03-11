package main

import (
	"context"

	logger "git.rarejob.com/shintaro.ikeda/platform_logging/new_logger"
)

type Service interface {
	SearchSample(ctx context.Context, sampleID string) (string, error)
}

type service struct{}

// NewService is ...
func NewService() Service {
	return newLoggingFilter(
		&service{},
		logger.DefaultLogger,
	)
}

func (s *service) SearchSample(ctx context.Context, sampleID string) (string, error) {
	return "SearchSample result", nil
}
