package new_logger

import (
	"github.com/sirupsen/logrus"
)

const (
	XTransactionID = "X-Transaction-ID"

	logKeyOfXTxID       = "transaction"
	logKeyOfServiceName = "service-name"
)

// DefaultLogger is a constructor of Logger interface with default config settings.
func DefaultLogger() Logger {
	return newLogger(NewConfig())
}

// NewLogger is a constructor of Logger interface which takes configuration.
func NewLogger(conf *Config) Logger {
	return NewLogger(conf)
}

func newLogger(config *Config) Logger {
	var l = logrus.New()
	{
		l.Level, _ = logrus.ParseLevel(config.minLevel.String())
		l.Formatter = config.formatter
		l.Out = config.out
	}
	return &logger{
		Logger: l,
		config: config,
	}
}

type LoggerContext struct {
	XTxID       string
	ServiceName string
	Msg         string
}

// Logger is an interface of Logging
type Logger interface {

	// Existing:  // Info(msg string, fields ...Field)
	// Solution1: // Info(xTxID, serviceName, msg string, fields ...Field)
	// Solution2: // Info(loggerCtx LoggerContext, fields ...Field)
	//
	// Solution1 は共通性が強い反面、変更に弱い
	// Solution2 は変更に強い反面、共通性に弱い

	Debug(xTxID, serviceName interface{}, msg string, fields ...Field)
	Info(xTxID, serviceName interface{}, msg string, fields ...Field)
	Warn(xTxID, serviceName interface{}, msg string, fields ...Field)
	Error(xTxID, serviceName interface{}, msg string, fields ...Field)
	Panic(xTxID, serviceName interface{}, msg string, fields ...Field)
}

type logger struct {
	*logrus.Logger
	config *Config
}

func (l *logger) Debug(xTxID, serviceName interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Debug) {
		l.withFields(l.mergeToFields(xTxID, serviceName, fields...)...).Debug(msg)
	}
}

func (l *logger) Info(xTxID, serviceName interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Info) {
		l.withFields(l.mergeToFields(xTxID, serviceName, fields...)...).Info(msg)
	}
}

func (l *logger) Warn(xTxID, serviceName interface{}, msg string, fields ...Field) {
	//if l.enabledLogLevel(Levels.Warn) {
	//	l.withFields(fields...).Warn(msg)
	//}

	if l.enabledLogLevel(Levels.Warn) {
		l.withFields(l.mergeToFields(xTxID, serviceName, fields...)...).Warn(msg)
	}
}

func (l *logger) Error(xTxID, serviceName interface{}, msg string, fields ...Field) {
	//func (l *logger) Error(msg string, fields ...Field) {
	//if l.enabledLogLevel(Levels.Error) {
	//	l.withFields(fields...).Error(msg)
	//}

	if l.enabledLogLevel(Levels.Error) {
		l.withFields(l.mergeToFields(xTxID, serviceName, fields...)...).Error(msg)
	}
}

func (l *logger) Panic(xTxID, serviceName interface{}, msg string, fields ...Field) {
	//if l.enabledLogLevel(Levels.Panic) {
	//	l.withFields(fields...).Panic(msg)
	//}

	if l.enabledLogLevel(Levels.Panic) {
		l.withFields(l.mergeToFields(xTxID, serviceName, fields...)...).Panic(msg)
	}
}

func (l *logger) mergeToFields(xTxID, serviceName interface{}, fields ...Field) []Field {
	mFields := make([]Field, 0, 2+len(fields)) // 2 means xTxID and serviceName
	xTxIDField := Field{key: logKeyOfXTxID, value: xTxID}
	serviceNameField := Field{key: logKeyOfServiceName, value: serviceName}
	mFields = append(mFields, xTxIDField, serviceNameField)
	return append(mFields, fields...)
}

func (l *logger) withFields(fields ...Field) *logrus.Entry {
	var cnvFields = make(logrus.Fields, len(fields))
	for _, f := range fields {
		if len(f.key) == 0 || f.value == nil {
			continue
		}
		cnvFields[f.key] = f.value
	}
	return l.Logger.WithFields(cnvFields)
}

func (l *logger) enabledLogLevel(level Level) bool {
	return l.config.minLevel <= level
}
