package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name     string
		actual   Level
		expected string
	}{
		{
			actual:   Levels.Debug,
			expected: "debug",
		},
		{
			actual:   Levels.Info,
			expected: "info",
		},
		{
			actual:   Levels.Warn,
			expected: "warn",
		},
		{
			actual:   Levels.Error,
			expected: "error",
		},
		{
			actual:   Levels.Panic,
			expected: "panic",
		},
		{
			actual:   Levels.Fatal,
			expected: "fatal",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.actual.String())
		})
	}
}

func TestLevels(t *testing.T) {
	assert.Equal(t, Levels, levels{
		Debug: 1,
		Info:  2,
		Warn:  3,
		Error: 4,
		Panic: 5,
		Fatal: 6,
	})
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected Level
	}{
		{
			actual:   "debug",
			expected: Levels.Debug,
		},
		{
			actual:   "info",
			expected: Levels.Info,
		},
		{
			actual:   "warn",
			expected: Levels.Warn,
		},
		{
			actual:   "error",
			expected: Levels.Error,
		},
		{
			actual:   "panic",
			expected: Levels.Panic,
		},
		{
			actual:   "fatal",
			expected: Levels.Fatal,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ParseLevel(test.actual))
		})
	}
}
