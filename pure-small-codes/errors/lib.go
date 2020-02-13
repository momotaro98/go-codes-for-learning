package main

import (
	"errors"
	"time"
)

var (
	ErrParse = errors.New("lib: parse error")
)

type LibTime time.Time

func ParseLibFunc(text string) (LibTime, error) {
	t, err := time.Parse(time.RFC3339, text)
	if err != nil {
		//return LibTime{}, &LibError{kind: "Parse", orgError: err}
		return LibTime{}, ErrParse
	}
	return LibTime(t), nil
}

type LibError struct {
	kind     string
	orgError error
}

func (l *LibError) Error() string {
	return "error occured in Lib"
}

func (l *LibError) Kind() string {
	return l.kind
}
