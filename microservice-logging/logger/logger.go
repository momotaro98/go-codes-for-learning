package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	XTransactionID = "X-Transaction-ID"

	logKeyOfXTxID       = "request-id"
	logKeyOfServiceName = "service-name"
)

var (
	// ServiceName は各アプリケーションのビルド時に ldflags を利用して埋め込まれます。
	// 下記がビルド時の例です。
	// go build -ldflags "-X path/microservice-logging/logger.ServiceName=MicroService01"
	ServiceName = "not-set"

	// Log は本パッケージのグローバルインスタンスです。
	Log = defaultLogger()
)

// Logger は本パッケージの公開インターフェースです。
type Logger interface {
	Debug(xTxID interface{}, msg string, fields ...Field)
	Info(xTxID interface{}, msg string, fields ...Field)
	Error(xTxID interface{}, msg string, fields ...Field)
}

func defaultLogger() Logger {
	return newLogger(NewConfig())
}

// NewLogger は Logger インターフェースのコンストラクタ です。
// 基本として、マイクロサービスのアプリケーションはこれを利用せずに Log インスタンスを利用することがルールです。
// 開発をする場合や調査をする場合などで Config を設定しこのメソッドを呼び出しましょう。
func NewLogger(conf *Config) Logger {
	return newLogger(conf)
}

type logger struct {
	*logrus.Logger
	config *Config
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

func (l *logger) Debug(xTxID interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Debug) {
		l.withFields(l.mergeToFields(xTxID, fields...)...).Debug(msg)
	}
}

func (l *logger) Info(xTxID interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Info) {
		l.withFields(l.mergeToFields(xTxID, fields...)...).Info(msg)
	}
}

func (l *logger) Error(xTxID interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Error) {
		l.withFields(l.mergeToFields(xTxID, fields...)...).Error(msg)
	}
}

func (l *logger) mergeToFields(xTxID interface{}, fields ...Field) []Field {
	// Required fields
	var requiredFields = []Field{
		{key: logKeyOfXTxID, value: xTxID},
		{key: logKeyOfServiceName, value: ServiceName},
	}

	mFields := make([]Field, 0, len(requiredFields)+len(fields))
	mFields = append(mFields, requiredFields...)
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
