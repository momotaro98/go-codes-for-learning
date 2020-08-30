package main

import (
	"fmt"
	"net/http"

	vegeta "github.com/tsenart/vegeta/lib"
)

type ScheduleTargeter struct {
	Method string
	URL    string
	Header http.Header
}

func NewScheduleTargeter(tokenChn <-chan string) vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		target := ScheduleTargeter{
			Method: http.MethodGet,
			URL:    "http://localhost:8090/schedule",
			Header: map[string][]string{
				"Content-Type": {"application/json"},
			},
		}

		token, ok := <-tokenChn
		if !ok {
			return vegeta.ErrNoTargets
		}

		tgt.Method = target.Method
		tgt.URL = target.URL
		tgt.Header = target.Header
		if val := token; ok {
			tgt.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", val)}
		}

		return nil
	}
}
