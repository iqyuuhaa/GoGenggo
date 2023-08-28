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
		}
	}

	return false
}
