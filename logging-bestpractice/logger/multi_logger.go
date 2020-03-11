package logger

type multiLogger struct {
	loggers []Logger
}

func (l *multiLogger) Debug(msg string, fields ...Field) {
	for _, logger := range l.loggers {
		logger.Debug(msg, fields...)
	}
}

func (l *multiLogger) Info(msg string, fields ...Field) {
	for _, logger := range l.loggers {
		logger.Info(msg, fields...)
	}
}

func (l *multiLogger) Warn(msg string, fields ...Field) {
	for _, logger := range l.loggers {
		logger.Warn(msg, fields...)
	}
}

func (l *multiLogger) Error(msg string, fields ...Field) {
	for _, logger := range l.loggers {
		logger.Error(msg, fields...)
	}
}

func (l *multiLogger) Panic(msg string, fields ...Field) {
	for _, logger := range l.loggers {
		logger.Panic(msg, fields...)
	}
}

func newMultiLogger(configs ...*Config) Logger {
	var loggers = make([]Logger, len(configs))
	for i, config := range configs {
		loggers[i] = newStandardLogger(config)
	}
	return &multiLogger{
		loggers: loggers,
	}
}
