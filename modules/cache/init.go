package cache

import (
	"gogenggo/internals/platforms"
)

var cacheModule *CacheModules
var cacheObject CacheMap

func InitServerCache(platform *platforms.PlatformModule) error {
	if cacheModule == nil {
		cacheModule = &CacheModules{
			Platform: platform,
		}
	}

	if err := cacheModule.populateCache(); err != nil {
		return err
	}

	return nil
}

func (c *CacheModules) populateCache() error {
	c.populateInitialData()

	return nil
}

func (c *CacheModules) populateInitialData() {
	cacheObject.MapRequestsData = make(map[int64]RequestTimeData)
}
