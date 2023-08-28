package constants

import (
	"errors"
	"net/http"

	"gogenggo/internals/types"
)

// Code for mapping
const (
	DefaultErrorCode         = "0"
	accessForbiddenErrorCode = "AF-1"
	accessDisabledErrorCode  = "AD-1"
	badRequestErrorCode      = "BR-1"
	dbMaintenanceErrorCode   = "DBM-1"
	tokenErrorCode           = "T-1"
	serverErrorCode          = "SE-1"
	notFoundDataErrorCode    = "NT-1"
	unauthorizedErrorCode    = "U-1"
)

// Error HTTP
var (
	ErrorAccessForbidden  = errors.New("ErrorAccessForbidden")
	ErrorAccessDisabled   = errors.New("ErrorAccessDisabled")
	ErrorBadRequest       = errors.New("ErrorBadRequest")
	ErrorDBMaintenance    = errors.New("ErrorDBMaintenance")
	ErrorEmptyToken       = errors.New("ErrorEmptyToken")
	ErrorInternalServer   = errors.New("ErrorInternalServer")
	ErrorInvalidFileType  = errors.New("ErrorInvalidFileType")
	ErrorInvalidToken     = errors.New("ErrorInvalidToken")
	ErrorNotFoundData     = errors.New("ErrorNotFoundData")
	ErrorUnauthorized     = errors.New("ErrorUnauthorized")
	ErrorWrongCredentials = errors.New("ErrorWrongCredentials")
)

var MapErrorCode = map[string]types.ErrorData{
	ErrorAccessForbidden.Error():  {HttpCode: http.StatusForbidden, ErrorCode: accessForbiddenErrorCode},
	ErrorAccessDisabled.Error():   {HttpCode: http.StatusForbidden, ErrorCode: accessForbiddenErrorCode},
	ErrorBadRequest.Error():       {HttpCode: http.StatusBadRequest, ErrorCode: badRequestErrorCode},
	ErrorDBMaintenance.Error():    {HttpCode: http.StatusInternalServerError, ErrorCode: dbMaintenanceErrorCode},
	ErrorEmptyToken.Error():       {HttpCode: http.StatusUnauthorized, ErrorCode: tokenErrorCode},
	ErrorInternalServer.Error():   {HttpCode: http.StatusInternalServerError, ErrorCode: serverErrorCode},
	ErrorInvalidFileType.Error():  {HttpCode: http.StatusBadRequest, ErrorCode: badRequestErrorCode},
	ErrorInvalidToken.Error():     {HttpCode: http.StatusUnauthorized, ErrorCode: tokenErrorCode},
	ErrorNotFoundData.Error():     {HttpCode: http.StatusNotFound, ErrorCode: notFoundDataErrorCode},
	ErrorUnauthorized.Error():     {HttpCode: http.StatusUnauthorized, ErrorCode: unauthorizedErrorCode},
	ErrorWrongCredentials.Error(): {HttpCode: http.StatusBadRequest, ErrorCode: badRequestErrorCode},
}

// Basic Error
var (
	ErrorNilValue = errors.New("ErrorNilValue")
)
