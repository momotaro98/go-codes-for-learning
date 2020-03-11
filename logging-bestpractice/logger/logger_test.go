package logger

import (
	"bytes"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetupRootLogger(t *testing.T) {
	tests := []struct {
		name     string
		logFunc  func()
		expected string
	}{
		{
			logFunc: func() {
				Debug("debug message",
					F("id", 123),
				)
			},
			expected: `level=debug msg="debug message" id=123`,
		},
		{
			logFunc: func() {
				Info("info message",
					F("id", 123),
				)
			},
			expected: `level=info msg="info message" id=123`,
		},
		{
			logFunc: func() {
				Warn("warning message",
					F("id", 123),
				)
			},
			expected: `level=warning msg="warning message" id=123`,
		},
		{
			logFunc: func() {
				Error("error message",
					E(errors.New("validation error")),
				)
			},
			expected: `level=error msg="error message" error="validation error"`,
		},
		{
			logFunc: func() {
				assert.Panics(t, func() {
					Panic("panic message",
						E(errors.New("unexpected error")),
					)
				})
			},
			expected: `level=panic msg="panic message" error="unexpected error"`,
		},
	}
	buff := new(bytes.Buffer)
	{
		SetupRootLogger(NewConfig("test",
			WithLevel(Levels.Debug),
			WithMaxLevel(Levels.Fatal),
			WithOut(buff),
		))
		if logger, ok := rootLogger.(*standardLogger); ok {
			logger.Formatter = &logrus.TextFormatter{
				DisableTimestamp: true,
			}
		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buff.Reset()
			test.logFunc()
			if len(test.expected) > 0 {
				assert.Equal(t, test.expected+"\n", buff.String())
			} else {
				assert.Empty(t, buff.String())
			}
		})
	}
}

func TestNewLogger(t *testing.T) {
	{
		configs := []*Config{
			NewConfig("test"),
		}
		actual := newLogger(configs...)
		expected := newStandardLogger(configs[0])
		assert.Equal(t, expected, actual)
	}
	{
		configs := []*Config{
			NewConfig("test1"),
			NewConfig("test2"),
		}
		actual := newLogger(configs...)
		expected := newMultiLogger(configs...)
		assert.Equal(t, expected, actual)
	}
	{
		configs := []*Config{}
		assert.PanicsWithValue(t, "configuration not found", func() {
			newLogger(configs...)
		})
	}
}
