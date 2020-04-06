package main

import (
	"io/ioutil"
	"testing"

	"github.com/doloopwhile/logrusltsv"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func loggigForBenech(l *logrus.Logger) {
	l.WithFields(log.Fields{
		"transaction": "Request-ID-1234567890",
	}).Info("This is INFO log level.")
}

func BenchmarkLTSV(b *testing.B) {
	var l = logrus.New()
	{
		l.Formatter = &logrusltsv.Formatter{}
		l.Out = ioutil.Discard
	}

	for i := 0; i < b.N; i++ {
		loggigForBenech(l)
	}
}

func BenchmarkJSON(b *testing.B) {
	var l = logrus.New()
	{
		l.Formatter = &logrus.JSONFormatter{}
		l.Out = ioutil.Discard
	}

	for i := 0; i < b.N; i++ {
		loggigForBenech(l)
	}
}
