package logger

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestStandardLogger(t *testing.T) {
	tests := []struct {
		name     string
		config   *Config
		logFunc  func(logger Logger)
		expected string
	}{
		// write log test patterns
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Debug("debug message",
					F("id", 123),
				)
			},
			expected: `level=debug msg="debug message" id=123`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Info("info message",
					F("id", 123),
				)
			},
			expected: `level=info msg="info message" id=123`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Warn("warning message",
					F("id", 123),
				)
			},
			expected: `level=warning msg="warning message" id=123`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				logger.Error("error message",
					E(errors.New("validation error")),
				)
			},
			expected: `level=error msg="error message" error="validation error"`,
		},
		{
			config: NewConfig("test"),
			logFunc: func(logger Logger) {
				assert.Panics(t, func() {
					logger.Panic("panic message",
						E(errors.New("unexpected error")),
					)
				})
			},
			expected: `level=panic msg="panic message" error="unexpected error"`,
		},
		// not write log test patterns
		{
			config: NewConfig("test",
				WithLevel(Levels.Info),
			),
			logFunc: func(logger Logger) {
				logger.Debug("debug message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Warn),
			),
			logFunc: func(logger Logger) {
				logger.Info("info message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Error),
			),
			logFunc: func(logger Logger) {
				logger.Warn("warning message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Panic),
			),
			logFunc: func(logger Logger) {
				logger.Error("error message")
			},
			expected: ``,
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Fatal),
			),
			logFunc: func(logger Logger) {
				logger.Panic("panic message")
			},
			expected: ``,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger, buff := mockStandardLogger(test.config)
			test.logFunc(logger)
			if len(test.expected) > 0 {
				assert.Equal(t, test.expected+"\n", buff.String())
			} else {
				assert.Empty(t, buff.String())
			}
		})
	}
}

func TestStandardLogger_Close(t *testing.T) {

}

func TestNewStandardLogger(t *testing.T) {
	config := NewConfig("test",
		WithLevel(Levels.Debug),
		WithFormatter(Formatters.JSON),
	)
	actual := newStandardLogger(config)
	expected := &standardLogger{
		Logger: logrus.New(),
		config: config,
	}
	{
		expected.Level = logrus.DebugLevel
		expected.Formatter = Formatters.JSON
		expected.Out = os.Stderr
	}
	assert.Equal(t, expected, actual)
}

func mockStandardLogger(config *Config) (Logger, *bytes.Buffer) {
	buffer := new(bytes.Buffer)
	{
		config.out = buffer
	}
	logger := newStandardLogger(config).(*standardLogger)
	{
		logger.Formatter = &logrus.TextFormatter{
			DisableTimestamp: true,
		}
	}
	return logger, buffer
}
