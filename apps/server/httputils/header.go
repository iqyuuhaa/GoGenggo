package httputils

import (
	"net/http"

	"gogenggo/internals/types/constants"
)

func GetLang(r *http.Request) string {
	lang := r.Header.Get("lang")
	if lang != "" {
		return lang
	}

	return constants.LangID
}
