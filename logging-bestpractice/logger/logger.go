package logger

import (
	"git.rarejob.com/shintaro.ikeda/platform_logging/logger/fio"
)

var (
	rootLogger = newLogger(NewConfig("default"))
)

// Logger is ...
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
}

// DefaultLogger ...
func DefaultLogger() Logger {
	return rootLogger
}

// Debug is ...
func Debug(msg string, fields ...Field) {
	rootLogger.Debug(msg, fields...)
}

// Info is ...
func Info(msg string, fields ...Field) {
	rootLogger.Info(msg, fields...)
}

// Warn is ...
func Warn(msg string, fields ...Field) {
	rootLogger.Warn(msg, fields...)
}

// Error is ...
func Error(msg string, fields ...Field) {
	rootLogger.Error(msg, fields...)
}

// Panic is ...
func Panic(msg string, fields ...Field) {
	rootLogger.Panic(msg, fields...)
}

// Close is ...
func Close() error {
	return fio.Close()
}

// SetupRootLogger is ...
func SetupRootLogger(configs ...*Config) {
	rootLogger = newLogger(configs...)
}

func newLogger(configs ...*Config) Logger {
	switch len(configs) {
	case 0:
		panic("configuration not found")
	case 1:
		return newStandardLogger(configs[0])
	default:
		return newMultiLogger(configs...)
	}
}
