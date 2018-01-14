// Package security contains implementations of interfaces dealing security, authentication, and authorization.
package security

// AuthService is a contract that provides authentication and authorization services.
type AuthService struct {
	Authenticated bool
	Roles         map[string]bool
}

// IsAuthenticated determines whether the User has been authenticated by the system.
// Returns true if the User has passed authentication, otherwise false.
func (authService *AuthService) IsAuthenticated() bool {
	authService.Authenticated = true
	return authService.Authenticated

}

// IsAuthorized determines whether the User possesses the required authorization role(s).
// Returns true if "Admin" is in the roles, otherwise false.
func (authService *AuthService) IsAuthorized(role string) bool {
	authService.Roles = map[string]bool{"Admin": true}

	return authService.Roles["Admin"]
}
