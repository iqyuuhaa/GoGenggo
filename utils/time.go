package utils

import "time"

func GetCurrentDayTime() (result string) {
	hour := time.Now().Hour()

	if hour >= 6 && hour <= 12 {
		result = "pagi"
	} else if hour >= 12 && hour <= 15 {
		result = "siang"
	} else if hour >= 15 && hour <= 18 {
		result = "sore"
	} else {
		result = "malam"
	}

	return result
}
