package cache

import (
	"time"

	"gogenggo/config"
	"gogenggo/utils"
)

func GetUserLatestSession(userID int64) (result string) {
	if v, ok := cacheObject.MapRequestsData[userID]; ok && v.SessionID != "" {
		result = v.SessionID
	}

	return result
}

func SetUserRequest(userID int64) {
	counter := 1
	start := time.Now()
	sessionID := utils.GenerateUUID()

	if v, ok := cacheObject.MapRequestsData[userID]; ok {
		counter = v.Counter + 1
		if !v.Start.IsZero() {
			start = v.Start
		}

		if v.SessionID != "" && (!v.Finish.IsZero() && v.Finish.Sub(start).Minutes() < float64(config.Configs.Main.Cache.LifetimeSessionID)) {
			sessionID = v.SessionID
		}
	}

	if len(cacheObject.MapRequestsData) == 0 || cacheObject.MapRequestsData == nil {
		cacheObject.MapRequestsData = make(map[int64]RequestTimeData)
	}

	cacheObject.Lock()
	cacheObject.MapRequestsData[userID] = RequestTimeData{
		SessionID: sessionID,
		Counter:   counter,
		Start:     start,
		Finish:    time.Now(),
	}
	cacheObject.Unlock()
}

func DeleteUserRequest(userID int64) {
	cacheObject.MapRequestsData[userID] = RequestTimeData{}
}

func ClearAllRequestCache() {
	cacheObject.Lock()
	cacheObject.MapRequestsData = make(map[int64]RequestTimeData)
	cacheObject.Unlock()
}
