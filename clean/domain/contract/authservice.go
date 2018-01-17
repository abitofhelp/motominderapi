// Package contract contains contracts for entities and other objects.
package contract

import "github.com/abitofhelp/motominderapi/clean/domain/enumeration"

// AuthService is a contract that provides authentication and authorization services.
type AuthService interface {
	IsAuthenticated() bool
	IsAuthorized(role enumeration.AuthorizationRole) bool
	Validate() error
}
