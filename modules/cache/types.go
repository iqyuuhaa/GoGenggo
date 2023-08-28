package cache

import (
	"sync"
	"time"

	"gogenggo/internals/platforms"
)

type CacheModules struct {
	Platform *platforms.PlatformModule
}

type CacheMap struct {
	sync.Mutex
	MapRequestsData map[int64]RequestTimeData
}

type RequestTimeData struct {
	SessionID string
	Counter   int
	Start     time.Time
	Finish    time.Time
}
