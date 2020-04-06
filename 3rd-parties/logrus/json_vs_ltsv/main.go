package main

/*
目的
LTSV と JSON での出力の容量の差を比較する
*/

import (
	"os"

	"github.com/doloopwhile/logrusltsv"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const (
	N = 1000
)

func logging(l *logrus.Logger) {
	for i := 0; i < N; i++ {
		l.WithFields(log.Fields{
			"transaction": "Request-ID-1234567890",
		}).Info("This is INFO log level.")
	}
}

func openFile(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	return f
}

func ltsvOutput() {
	file := openFile(`ltsv.log`)
	defer file.Close()

	var l = logrus.New()
	{
		l.Formatter = &logrusltsv.Formatter{}
		l.Out = file
	}

	logging(l)
}

func jsonOutput() {
	file := openFile(`json.log`)
	defer file.Close()

	var l = logrus.New()
	{
		l.Formatter = &logrus.JSONFormatter{}
		l.Out = file
	}

	logging(l)
}

func main() {
	ltsvOutput()

	jsonOutput()
}
