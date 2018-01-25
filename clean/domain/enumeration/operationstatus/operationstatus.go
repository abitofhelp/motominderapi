// Package operationstatus defines operation status values for the application.
package operationstatus

// OperationStatus indicates the success or error of an operation.
type OperationStatus int

// The list of valid operation status values.
const (
	Undefined        = 0
	Ok               = 200
	Created          = 201
	NoContent        = 204
	Found            = 302
	BadRequest       = 400
	NotAuthenticated = 401
	NotAuthorized    = 403
	NotFound         = 404
	InternalError    = 500
)

// descriptions are the textual message for each operation status value.
var descriptions = [...]string{
	"Undefined",
	"Ok",
	"Created",
	"No Content",
	"Found",
	"Bad Request",
	"Not Authenticated",
	"Not Authorized",
	"Not Found",
	"Internal Error",
}

// ToString provides a description for the operation status value.
func (status OperationStatus) ToString() string {
	return descriptions[status-1]
}
