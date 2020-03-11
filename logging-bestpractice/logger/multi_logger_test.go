package logger

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMultiLogger(t *testing.T) {
	configs := []*Config{
		NewConfig("app",
			WithLevel(Levels.Debug),
			WithMaxLevel(Levels.Warn),
		),
		NewConfig("error",
			WithLevel(Levels.Error),
			WithMaxLevel(Levels.Fatal),
		),
	}

	tests := []struct {
		name     string
		configs  []*Config
		logFunc  func(logger Logger)
		expected []string
	}{
		{
			configs: configs,
			logFunc: func(logger Logger) {
				logger.Debug("debug message")
			},
			expected: []string{
				`level=debug msg="debug message"`,
				``,
			},
		},
		{
			configs: configs,
			logFunc: func(logger Logger) {
				logger.Info("info message")
			},
			expected: []string{
				`level=info msg="info message"`,
				``,
			},
		},
		{
			configs: configs,
			logFunc: func(logger Logger) {
				logger.Warn("warning message")
			},
			expected: []string{
				`level=warning msg="warning message"`,
				``,
			},
		},
		{
			configs: configs,
			logFunc: func(logger Logger) {
				logger.Error("error message")
			},
			expected: []string{
				``,
				`level=error msg="error message"`,
			},
		},
		{
			configs: configs,
			logFunc: func(logger Logger) {
				assert.Panics(t, func() {
					logger.Panic("panic message")
				})
			},
			expected: []string{
				``,
				`level=panic msg="panic message"`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger, buffs := mockMultiLogger(test.configs)
			test.logFunc(logger)
			if assert.Equal(t, len(test.expected), len(buffs)) {
				for i, expected := range test.expected {
					if len(expected) > 0 {
						assert.Equal(t, expected+"\n", buffs[i].String())
					} else {
						assert.Empty(t, buffs[i].String())
					}
				}
			}
		})
	}
}

func TestNewMultiLogger(t *testing.T) {
	configs := []*Config{
		NewConfig("app",
			WithLevel(Levels.Debug),
			WithMaxLevel(Levels.Warn),
		),
		NewConfig("error",
			WithLevel(Levels.Error),
			WithMaxLevel(Levels.Fatal),
		),
	}
	actual := newMultiLogger(configs...)
	expected := &multiLogger{
		loggers: []Logger{
			newStandardLogger(configs[0]),
			newStandardLogger(configs[1]),
		},
	}
	assert.Equal(t, expected, actual)
}

func mockMultiLogger(configs []*Config) (Logger, []*bytes.Buffer) {
	buffers := make([]*bytes.Buffer, len(configs))
	{
		for i, config := range configs {
			buffers[i] = new(bytes.Buffer)
			config.out = buffers[i]
		}
	}
	logger := newMultiLogger(configs...).(*multiLogger)
	{
		for _, l := range logger.loggers {
			if stdLogger, ok := l.(*standardLogger); ok {
				stdLogger.Formatter = &logrus.TextFormatter{
					DisableTimestamp: true,
				}
			}
		}
	}
	return logger, buffers
}
