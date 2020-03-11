package fio

import (
	"io"
	"path"

	"gopkg.in/natefinch/lumberjack.v2"
)

// FileConfig is ...
type FileConfig struct {
	logName      string
	logDirectory string
	maxSize      int  // megabytes
	maxBackups   int  // counts
	maxAge       int  // days
	compress     bool // disabled by default
	localTime    bool
}

// MaxSize is ...
func (c *FileConfig) MaxSize() int {
	return c.maxSize
}

// MaxBackups is ...
func (c *FileConfig) MaxBackups() int {
	return c.maxBackups
}

// MaxAge is ...
func (c *FileConfig) MaxAge() int {
	return c.maxAge
}

// IsCompress is ...
func (c *FileConfig) IsCompress() bool {
	return c.compress
}

// IsLocalTime is ...
func (c *FileConfig) IsLocalTime() bool {
	return c.localTime
}

// FilePath is ...
func (c *FileConfig) FilePath() string {
	if len(c.logDirectory) == 0 {
		return ""
	}
	return path.Join(
		c.logDirectory,
		c.logName,
	)
}

// FileConfigOption is ...
type FileConfigOption func(config *FileConfig)

// WithLogDirectory is ...
func WithLogDirectory(logDirectory string) FileConfigOption {
	return func(c *FileConfig) {
		c.logDirectory = logDirectory
	}
}

// WithMaxSize is ...
func WithMaxSize(maxSize int) FileConfigOption {
	return func(c *FileConfig) {
		c.maxSize = maxSize
	}
}

// WithMaxBackups is ...
func WithMaxBackups(maxBackups int) FileConfigOption {
	return func(c *FileConfig) {
		c.maxBackups = maxBackups
	}
}

// WithMaxAge is ...
func WithMaxAge(maxAge int) FileConfigOption {
	return func(c *FileConfig) {
		c.maxAge = maxAge
	}
}

// WithCompress is ...
func WithCompress(compress bool) FileConfigOption {
	return func(c *FileConfig) {
		c.compress = compress
	}
}

// WithLocalTime is ...
func WithLocalTime(localTime bool) FileConfigOption {
	return func(c *FileConfig) {
		c.localTime = localTime
	}
}

// NewFileConfig is ...
func NewFileConfig(logName string, options ...FileConfigOption) *FileConfig {
	config := &FileConfig{
		logName:      logName,
		logDirectory: "logs",
		maxSize:      10,
		maxBackups:   5,
		maxAge:       30,
		compress:     false,
		localTime:    true,
	}
	for _, option := range options {
		option(config)
	}
	return config
}

// NewFileWriter is ...
func NewFileWriter(config *FileConfig) io.Writer {
	out := &lumberjack.Logger{
		Filename:   config.FilePath(),   // output file path
		MaxSize:    config.MaxSize(),    // megabytes
		MaxBackups: config.MaxBackups(), // counts
		MaxAge:     config.MaxAge(),     // days
		Compress:   config.IsCompress(), // disabled by default
		LocalTime:  config.IsLocalTime(),
	}
	// append rotate caller function
	rotateFunctionCaller.append(func() error {
		return out.Rotate()
	})
	// append close caller function
	closeFunctionCaller.append(func() error {
		return out.Close()
	})
	return out
}
