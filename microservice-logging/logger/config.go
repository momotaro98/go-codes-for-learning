package logger

import (
	"io"
	"os"
)

// Config はLoggerのための設定の構造体です。
//
// formatter: Format type of logging, TEXT or JSON
// out:       io.Writer of the logger output
// minLevel:  Minimum level to out
type Config struct {
	formatter Formatter
	out       io.Writer
	minLevel  Level
}

// ConfigOption はFunctional Options Patternの
// Configのフィールドを設定するための関数の型です。
type ConfigOption func(*Config)

// NewConfig は *Config のコンストラクタです。
//
// マイクロサービスのルールであるデフォルトの設定
// Formatter: JSON
// Out:       STD OUT
// MinLevel:  Info
func NewConfig(options ...ConfigOption) *Config {
	config := &Config{
		formatter: Formatters.JSON,
		out:       os.Stdout,
		minLevel: Levels.Info,
	}
	for _, option := range options {
		option(config)
	}
	return config
}

func WithMinLevel(minLevel Level) ConfigOption {
	return func(c *Config) {
		c.minLevel = minLevel
	}
}

func WithFormatter(formatter Formatter) ConfigOption {
	return func(c *Config) {
		c.formatter = formatter
	}
}

func WithOut(out io.Writer) ConfigOption {
	return func(c *Config) {
		c.out = out
	}
}

