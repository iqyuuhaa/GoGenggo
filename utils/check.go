package utils

import "strings"

func InArray(s string, str []string, isStrict bool) bool {
	for _, v := range str {
		if v == s {
			return true
		}

		if isStrict {
			if ret := strings.Compare(s, v); ret > -1 {
				return true
			}

			lowerValue := strings.ToLower(v)
			if ret := strings.Contains(s, lowerValue); ret {
				return true
			}
		}
	}

	return false
}
