// Package enumerations defines enumerations for the application.
package enumeration

// AuthorizationRole is an authorization given to an authenticated user to access a resource.
type AuthorizationRole int

const (
	// NoAuthorizationRole is when the user does not require any authorization roles to access resources.
	NoAuthorizationRole = iota
	// AdminAuthorizationRole is a user with authorization to access administrative resources.
	AdminAuthorizationRole
	// AccountingAuthorizationRole is a user with authorization to access accounting resources.
	AccountingAuthorizationRole
	// GeneralAuthorizationRole is a user with administrative authorization to access general resources.
	GeneralAuthorizationRole
)

// Admin provides a useful string for the value.
// Returns a string for the Admin constant.
func (AuthorizationRole) Admin() string {
	return "Admin"
}

// Accounting provides a useful string for the value.
// Returns a string for the Accounting constant.
func (AuthorizationRole) Accounting() string {
	return "Accounting"
}

// General provides a useful string for the value.
// Returns a string for the General constant.
func (AuthorizationRole) General() string {
	return "General"
}
