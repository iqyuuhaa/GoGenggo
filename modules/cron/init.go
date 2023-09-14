package cron

import (
	"os"
	"time"

	"gogenggo/config"
	"gogenggo/internals/types/constants"
	"gogenggo/modules/cache"

	"github.com/go-co-op/gocron"
)

func InitCron() {
	timezone := os.Getenv(constants.Timezone)
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	loc, _ := time.LoadLocation(timezone)
	s := gocron.NewScheduler(loc)

	clearRequestCache(s)
	populateConfig(s)

	s.StartAsync()
}

func clearRequestCache(s *gocron.Scheduler) {
	s.Every(1).Day().At("00:00").Do(func() {
		cache.ClearAllRequestCache()
	})
}

func populateConfig(s *gocron.Scheduler) {
	s.Every(30).Seconds().Do(func() {
		config.LoadConfig()
	})
}
