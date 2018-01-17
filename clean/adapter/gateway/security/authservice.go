// Package security contains implementations of interfaces dealing security, authentication, and authorization.
package security

import (
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/go-ozzo/ozzo-validation"
)

// AuthService is a contract that provides authentication and authorization services.
type AuthService struct {
	Authenticated bool
	Roles         map[enumeration.AuthorizationRole]bool
}

// Validate verifies that an AuthService's fields contain valid data.
// Returns nil if the AuthService contains valid data, otherwise an error.
func (authService AuthService) Validate() error {
	return validation.ValidateStruct(&authService,
		// Authenticated defaults to false.

		// Roles cannot be nil.
		validation.Field(&authService.Roles, validation.NotNil),
	)
}

// NewAuthService creates a new instance of an AuthService.
// Returns (nil, error) when there is an error, otherwise (authService, nil).
func NewAuthService(authenticated bool, roles map[enumeration.AuthorizationRole]bool) (*AuthService, error) {

	authService := &AuthService{
		Authenticated: authenticated,
		Roles:         roles,
	}

	err := authService.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return authService, nil
}

// IsAuthenticated determines whether the User has been authenticated by the system.
// Returns true if the User has passed authentication, otherwise false.
func (authService *AuthService) IsAuthenticated() bool {
	return authService.Authenticated

}

// IsAuthorized determines whether the User possesses the required authorization role(s).
// Returns true if "Admin" is in the roles, otherwise false.
func (authService *AuthService) IsAuthorized(role enumeration.AuthorizationRole) bool {
	return authService.Roles[enumeration.AdminAuthorizationRole]
}
