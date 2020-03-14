package logger


import (
	"io"
	"os"
)

// Config is a struct for logger
//
// minLevel: Minimum level to log. When the minLevel is Warn,
// logger logs only "Warn", "Error", "Fatal"
//
// formatter: Format type of logging, TEXT or JSON
//
// out: io.Writer to log
type Config struct {
	minLevel  Level
	formatter Formatter
	out       io.Writer
}

// ConfigOption is ...
type ConfigOption func(*Config)

// WithMaxLevel is ...
func WithMinLevel(minLevel Level) ConfigOption {
	return func(c *Config) {
		c.minLevel = minLevel
	}
}

// WithFormatter is ...
func WithFormatter(formatter Formatter) ConfigOption {
	return func(c *Config) {
		c.formatter = formatter
	}
}

// WithOut is ...
func WithOut(out io.Writer) ConfigOption {
	return func(c *Config) {
		c.out = out
	}
}

// NewConfig makes Configuration struct of logging.
// This takes option args if needed as Functional options pattern.
// Defaults are MinLevel: Info, Formatter: JSON, Out: std out
func NewConfig(options ...ConfigOption) *Config {
	config := &Config{
		minLevel: Levels.Info,
		formatter: Formatters.JSON,
		out:       os.Stdout,
	}
	for _, option := range options {
		option(config)
	}
	return config
}
