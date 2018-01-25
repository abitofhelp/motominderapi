// Package authorizationrole defines authorization roles for the application.
package authorizationrole

// AuthorizationRole is an authorization given to an authenticated user to access a resource.
type AuthorizationRole int

// The list of valid authorization role values.
const (
	// UndefinedAuthorizationRole is when an authorization role has not been assigned.
	UndefinedAuthorizationRole = 0
	// NoAuthorizationRole is when the user does not require any authorization roles to access resources.
	NoAuthorizationRole = iota
	// AdminAuthorizationRole is a user with authorization to access administrative resources.
	AdminAuthorizationRole
	// AccountingAuthorizationRole is a user with authorization to access accounting resources.
	AccountingAuthorizationRole
	// GeneralAuthorizationRole is a user with administrative authorization to access general resources.
	GeneralAuthorizationRole
)

// descriptions are the textual message for each authorization role value.
var descriptions = [...]string{
	"Undefined",
	"None",
	"Admin",
	"Accounting",
	"General",
}

// ToString provides a description for the authorization role value.
func (role AuthorizationRole) ToString() string {
	return descriptions[role-1]
}
