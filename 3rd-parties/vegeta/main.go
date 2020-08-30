package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

type TargeterType int

const (
	Login TargeterType = iota
	Schedule
)

type MetricsMutex struct {
	sync.Mutex
	vegeta.Metrics
}

type ProductTargeter struct {
	ttype    TargeterType
	targeter vegeta.Targeter
}

func main() {
	var (
		rate     = flag.Int("rate", 10, "Number of requests per time unit [0 = infinity] (default 10/1s)")
		duration = flag.Int("duration", 10, "Duration of the test [0 = forever]")
	)
	flag.Parse()

	tokenChn := make(chan string)

	targeters := []*ProductTargeter{
		{
			Login,
			NewLoginTargeter("login_user.csv"),
		},
		{
			Schedule,
			NewScheduleTargeter(tokenChn),
		},
	}

	fmt.Println(*rate)
	rt := vegeta.Rate{Freq: *rate / len(targeters), Per: time.Second}
	dur := time.Duration(*duration) * time.Second

	var metrics MetricsMutex

	var wg sync.WaitGroup
	for _, t := range targeters {
		wg.Add(1)
		t := t
		go func() {
			defer wg.Done()
			for res := range vegeta.NewAttacker().Attack(t.targeter, rt, dur, string(t.ttype)) {
				if res == nil || res.Error == vegeta.ErrNoTargets.Error() {
					continue
				}
				switch t.ttype {
				case Login:
					go PassToken(res.Body, tokenChn)
				}
				metrics.Lock()
				metrics.Add(res)
				metrics.Unlock()
			}
		}()
	}

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for res := range vegeta.NewAttacker().Attack(loginTargeter, rate, dur, "login-attacker") {
	//		go PassToken(res.Body, tokenChn)
	//		metrics.Lock()
	//		metrics.Add(res)
	//		metrics.Unlock()
	//	}
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for res := range vegeta.NewAttacker().Attack(scheduleTargeter, rate, dur, "schedule-attacker") {
	//		metrics.Lock()
	//		metrics.Add(res)
	//		metrics.Unlock()
	//	}
	//}()

	//loop:
	//	for {
	//		select {
	//		case res := <-vegeta.NewAttacker().Attack(loginTargeter, rate, dur, "login-attacker"):
	//			fmt.Println("login dadada")
	//			if res == nil || res.Error == vegeta.ErrNoTargets.Error() {
	//				continue
	//			}
	//			go PassToken(res.Body, tokenChn)
	//			metrics.Lock()
	//			metrics.Add(res)
	//			metrics.Unlock()
	//		case res := <-vegeta.NewAttacker().Attack(scheduleTargeter, rate, dur, "schedule-attacker"):
	//			if res == nil {
	//				continue
	//			}
	//			if res.Error == vegeta.ErrNoTargets.Error() {
	//				fmt.Println("break?")
	//				fmt.Println(res)
	//				break loop
	//			}
	//			metrics.Lock()
	//			metrics.Add(res)
	//			metrics.Unlock()
	//		}
	//	}

	wg.Wait()
	metrics.Close()

	b, err := json.Marshal(metrics)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(b)
}
