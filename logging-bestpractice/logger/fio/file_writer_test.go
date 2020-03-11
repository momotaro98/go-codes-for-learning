package fio

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestFileConfig(t *testing.T) {
	tests := []struct {
		name     string
		config   *FileConfig
		expected *FileConfig
	}{
		{
			config: NewFileConfig("app.log"),
			expected: &FileConfig{
				logName:      "app.log",
				logDirectory: "logs",
				maxSize:      10,
				maxBackups:   5,
				maxAge:       30,
				compress:     false,
				localTime:    true,
			},
		},
		{
			config: NewFileConfig("app.log",
				WithLogDirectory("/var/log/event"),
				WithMaxSize(100),
				WithMaxBackups(10),
				WithMaxAge(60),
				WithCompress(true),
				WithLocalTime(false),
			),
			expected: &FileConfig{
				logName:      "app.log",
				logDirectory: "/var/log/event",
				maxSize:      100,
				maxBackups:   10,
				maxAge:       60,
				compress:     true,
				localTime:    false,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config)
		})
	}
}

func TestFileConfig_FilePath(t *testing.T) {
	tests := []struct {
		name     string
		config   *FileConfig
		expected string
	}{
		{
			config: NewFileConfig("app.log",
				WithLogDirectory(""),
			),
			expected: "",
		},
		{
			config:   NewFileConfig("app.log"),
			expected: "logs/app.log",
		},
		{
			config: NewFileConfig("app.log",
				WithLogDirectory("/var/log/event"),
			),
			expected: "/var/log/event/app.log",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config.FilePath())
		})
	}
}

func TestNewFileWriter(t *testing.T) {
	config := &FileConfig{
		logName:      "app.log",
		logDirectory: "",
		maxSize:      100,
		maxBackups:   10,
		maxAge:       60,
		compress:     true,
		localTime:    false,
	}
	actual := NewFileWriter(config)
	expected := &lumberjack.Logger{
		Filename:   config.FilePath(),   // output file path
		MaxSize:    config.MaxSize(),    // megabytes
		MaxBackups: config.MaxBackups(), // counts
		MaxAge:     config.MaxAge(),     // days
		Compress:   config.IsCompress(), // disabled by default
		LocalTime:  config.IsLocalTime(),
	}
	assert.Equal(t, expected, actual)
}
