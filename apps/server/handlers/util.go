package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gogenggo/apps/server/httputils"
	"gogenggo/internals/types"
	"gogenggo/internals/types/constants"
	"gogenggo/utils"

	"github.com/nicksnyder/go-i18n/i18n"
)

// Main used functions
func (h *Handler) writeSuccessResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	header := types.DefaultAPIResponseHeader{
		Status:     "Ok",
		StatusCode: http.StatusOK,
	}

	h.writeResponseAPI(w, r, header, data)
}

func (h *Handler) writeErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	T, _ := i18n.Tfunc(httputils.GetLang(r))
	message := T(err.Error())
	statusCode, errCode := h.getErrorStatusCode(err.Error())

	h.processErrorResponse(w, r, statusCode, errCode, message)
}

// End main used functions

func (h *Handler) writeErrorResponseWithMessage(w http.ResponseWriter, r *http.Request, err error, errMessage string) {
	statusCode, errCode := h.getErrorStatusCode(err.Error())

	h.processErrorResponse(w, r, statusCode, errCode, errMessage)
}

// Private func
func (h *Handler) writeResponseAPI(w http.ResponseWriter, r *http.Request, header types.DefaultAPIResponseHeader, data interface{}) {
	processTime := time.Since(utils.GetStartTimeProcessCtx(r.Context())).Seconds()
	header.ProcessTime = fmt.Sprint(processTime)

	resp := types.DefaultAPIResponse{
		Header: header,
		Data:   data,
	}

	encoded, err := json.Marshal(resp)
	if err != nil {
		log.Println("[Handler - writeResponse] Error marshalling json response, err: ", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}

func (h *Handler) processErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, errCode, message string) {
	header := types.DefaultAPIResponseHeader{
		Status:     http.StatusText(statusCode),
		StatusCode: statusCode,
		Message:    &message,
		ErrorCode:  &errCode,
	}

	data := types.DefaultAPISuccessResponse{
		IsSuccess: false,
	}

	h.writeResponseAPI(w, r, header, data)
}

func (h *Handler) getErrorStatusCode(err string) (int, string) {
	if errData, isExist := constants.MapErrorCode[err]; isExist {
		return errData.HttpCode, errData.ErrorCode
	}

	return http.StatusInternalServerError, constants.DefaultErrorCode
}
