package utils

import (
	"fmt"
	"strings"
	"time"

	"gogenggo/internals/types/constants"
)

func ReplaceStringsFormat(s string, values map[string]string) (result string) {
	result = s
	for keyName, replaceStringKey := range constants.ReplacedStringKeys {
		result = strings.ReplaceAll(result, replaceStringKey, values[keyName])
	}

	return result
}

func NormalizeMsisdnFormat(msisdn string) (result string) {
	// if msisdn[0:1] != "0" || msisdn[0:3] != "+62" || msisdn[0:2] != "62" {
	// 	return ""
	// }

	result = msisdn
	if msisdn[0:1] == "0" {
		result = fmt.Sprintf("62%s", msisdn[1:])
	}

	if msisdn[0:1] == "+" {
		result = msisdn[1:]
	}

	return result
}

func FormatingDateTime(dateTime time.Time) (result string) {
	year, month, day := dateTime.Date()
	hour, minute, second := dateTime.Clock()

	return fmt.Sprintf("%d-%d-%d %d:%d:%d", year, int(month), day, hour, minute, second)
}
