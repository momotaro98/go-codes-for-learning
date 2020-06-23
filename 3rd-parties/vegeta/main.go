package main

import (
	"encoding/json"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	/*
		f, err := os.Open(`target.txt`)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		targeter := vegeta.NewHTTPTargeter(f, nil, nil)
	*/

	targeter := NewScenarioTargeter()

	rate := vegeta.Rate{Freq: 5, Per: time.Second}
	dur := time.Duration(5) * time.Second

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, dur, "sample-attacker") {
		metrics.Add(res)
	}
	metrics.Close()

	b, err := json.Marshal(metrics)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(b)
}
