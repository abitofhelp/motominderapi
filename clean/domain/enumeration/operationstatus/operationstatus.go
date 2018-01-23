// Package enumeration defines enumerations for the application.
package operationstatus

import "github.com/golang/go/src/pkg/net/http"

// OperationStatus is the status of a repository operation.
type OperationStatus int

// Various operation status values.
const (
	StatusOK        = 200
	StatusCreated   = 201
	StatusNoContent = 204

	StatusFound       = 302
	StatusNotModified = 304

	StatusBadRequest         = 400
	StatusUnauthorized       = 401 // Not authenticated.
	StatusForbidden          = 403 // Not authorized.
	StatusNotFound           = 404
	StatusConflict           = 409
	StatusPreconditionFailed = 412

	StatusInternalServerError = 500
	StatusNotImplemented      = 501
)

// statusText converts an OperationStatus to a string.
var statusText = map[int]string{
	StatusOK:        "OK",
	StatusCreated:   "Created",
	StatusNoContent: "No Content",

	StatusFound:       "Found",
	StatusNotModified: "Not Modified",

	StatusBadRequest:         "Bad Request",
	StatusUnauthorized:       "Not Authenticated",
	StatusForbidden:          "Not authorized",
	StatusNotFound:           "Not Found",
	StatusConflict:           "Conflict",
	StatusPreconditionFailed: "Precondition Failed",

	StatusInternalServerError: "Internal Server Error",
	StatusNotImplemented:      "Not Implemented",
}

// StatusText returns a text for the repository status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}

// MapOperationStatusToHttpStatus maps operation status values to HTTP status codes.  There will not always be a 1:1 correspondence.
func ToHttpStatus(status int) int {
	switch status {
	case StatusOK:
		return http.StatusOK
	case StatusCreated:
		return http.StatusCreated
	case StatusNoContent:
		return http.StatusNoContent
	case StatusFound:
		return http.StatusFound
	case StatusNotModified:
		return http.StatusNotModified
	case StatusBadRequest:
		return http.StatusBadRequest
	case StatusUnauthorized:
		return http.StatusUnauthorized
	case StatusForbidden:
		return http.StatusForbidden
	case StatusNotFound:
		return http.StatusNotFound
	case StatusConflict:
		return http.StatusConflict
	case StatusPreconditionFailed:
		return http.StatusPreconditionFailed
	case StatusInternalServerError:
		return http.StatusInternalServerError
	case StatusNotImplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusNotImplemented
	}
}
