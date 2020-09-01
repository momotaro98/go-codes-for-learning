package main

import (
	"encoding/json"
	"flag"
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
			// ログインエンドポイント
			Login,
			NewLoginTargeter("login_user.csv"),
		},
		{
			// 予定取得エンドポイント
			Schedule,
			NewScheduleTargeter(tokenChn),
		},
	}

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
				// 全アカウント分終了すると空の結果が返るのでそれはMetricsに含めないようにする
				if res == nil || res.Error == vegeta.ErrNoTargets.Error() {
					continue
				}
				switch t.ttype {
				case Login:
					// ログインから返ったJWTトークンを予定取得リクエストに引き渡す
					go PassToken(res.Body, tokenChn)
				}
				metrics.Lock()
				metrics.Add(res)
				metrics.Unlock()
			}
		}()
	}

	wg.Wait()
	metrics.Close()

	// 結果をJSONフォーマットにする
	b, err := json.Marshal(metrics)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(b)
}
