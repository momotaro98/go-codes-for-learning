package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name     string
		config   *Config
		expected *Config
	}{
		{
			config: NewConfig("test"),
			expected: &Config{
				name:      "test",
				level:     Levels.Debug,
				maxLevel:  Levels.Fatal,
				formatter: Formatters.Text,
				out:       os.Stderr,
			},
		},
		{
			config: NewConfig("test",
				WithLevel(Levels.Info),
				WithMaxLevel(Levels.Warn),
				WithFormatter(Formatters.JSON),
				WithOut(os.Stdout),
			),
			expected: &Config{
				name:      "test",
				level:     Levels.Info,
				maxLevel:  Levels.Warn,
				formatter: Formatters.JSON,
				out:       os.Stdout,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config)
		})
	}
}
