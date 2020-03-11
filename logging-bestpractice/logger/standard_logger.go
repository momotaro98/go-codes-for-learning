package logger

import (
	"github.com/sirupsen/logrus"
)

type standardLogger struct {
	*logrus.Logger
	config *Config
}

func (l *standardLogger) Debug(msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Debug) {
		l.withFields(fields).Debug(msg)
	}
}

func (l *standardLogger) Info(msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Info) {
		l.withFields(fields).Info(msg)
	}
}

func (l *standardLogger) Warn(msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Warn) {
		l.withFields(fields).Warn(msg)
	}
}

func (l *standardLogger) Error(msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Error) {
		l.withFields(fields).Error(msg)
	}
}

func (l *standardLogger) Panic(msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Panic) {
		l.withFields(fields).Panic(msg)
	}
}

func (l *standardLogger) withFields(fields []Field) *logrus.Entry {
	var cnvFields = make(logrus.Fields, len(fields))
	for _, f := range fields {
		if len(f.key) == 0 || f.value == nil {
			continue
		}
		cnvFields[f.key] = f.value
	}
	return l.Logger.WithFields(cnvFields)
}

func (l *standardLogger) enabledLogLevel(level Level) bool {
	return l.config.level <= level && level <= l.config.maxLevel
}

func newStandardLogger(config *Config) Logger {
	var l = logrus.New()
	{
		l.Level, _ = logrus.ParseLevel(config.level.String())
		l.Formatter = config.formatter
		l.Out = config.out
	}
	return &standardLogger{
		Logger: l,
		config: config,
	}
}
