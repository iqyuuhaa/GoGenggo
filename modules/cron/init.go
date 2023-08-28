package cron

import (
	"os"
	"time"

	"gogenggo/modules/cache"

	"github.com/go-co-op/gocron"
)

func InitCron() {
	timezone := os.Getenv("TZ")
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	loc, _ := time.LoadLocation(timezone)
	s := gocron.NewScheduler(loc)
	s.Every(1).Day().At("00:00").Do(func() {
		cache.ClearAllRequestCache()
	})

	s.StartAsync()
}
