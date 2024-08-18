package errors

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 
 */

import (
	"net/http"
)

var httpCode = map[ErrorCode]int{
	UNKNOWN_ERROR:         http.StatusInternalServerError,
	DATA_INVALID:          http.StatusBadRequest,
	STATUS_PAGE_NOT_FOUND: http.StatusNotFound,
	UNAUTHORIZED:          http.StatusUnauthorized,
	FAILED_RETRIEVE_DATA:  http.StatusInternalServerError,
	FAILED_CREATE_DATA:    http.StatusInternalServerError,
}
