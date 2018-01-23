// Package contract contains contracts for entities and other objects.
package contract

import (
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
)

// AuthService is a contract that provides authentication and authorization services.
type AuthService interface {
	IsAuthenticated() bool
	IsAuthorized(role authorizationrole.AuthorizationRole) bool
	Validate() error
}
