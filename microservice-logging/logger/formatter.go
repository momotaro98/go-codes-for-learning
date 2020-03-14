package logger

import "github.com/sirupsen/logrus"

// Formatter is ...
type Formatter struct {
	logrus.Formatter
}

// Formatters is ...
var Formatters = struct {
	Text Formatter
	JSON Formatter
}{
	Text: Formatter{Formatter: &logrus.TextFormatter{}},
	JSON: Formatter{Formatter: &logrus.JSONFormatter{}},
}

