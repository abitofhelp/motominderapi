// Package interfaces contains contracts for entities and other objects.
package interfaces

import "github.com/abitofhelp/motominderapi/clean/domain/enumerations"

// AuthService is a contract that provides authentication and authorization services.
type AuthService interface {
	IsAuthenticated() bool
	IsAuthorized(role enumerations.AuthorizationRole) bool
	Validate() error
}
