package logger

import (
	"sync"

	"git.rarejob.com/shintaro.ikeda/platform_logging/logger/fio"
	"github.com/robfig/cron"
)

var (
	timer *cron.Cron
	once  sync.Once
)

// StartCronTimer is ...
func StartCronTimer(expression string) {
	once.Do(func() {
		timer = cron.New()
		timer.AddFunc(expression, func() {
			fio.Rotate()
		})
		timer.Start()
	})
}

// StopCronTimer is ...
func StopCronTimer() error {
	if timer != nil {
		timer.Stop()
	}
	return Close()
}
