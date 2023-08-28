package httputils

import (
	"encoding/json"
	"log"
	"net/http"

	"gogenggo/internals/types/constants"
)

func GetFromBodyJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return constants.ErrorBadRequest
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("[Utils - GetFromBodyJSON] Error Unmarshal request, err:", err)
		return constants.ErrorBadRequest
	}

	return nil
}

func GetFromFormValue(r *http.Request, index string, isRequired bool) (string, error) {
	v := r.FormValue(index)
	if v == "" && isRequired {
		return "", constants.ErrorBadRequest
	}

	return v, nil
}
