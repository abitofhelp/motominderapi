// Package interfaces contains contracts for entities and other objects.
package interfaces

// AuthService is a contract that provides authentication and authorization services.
type AuthService interface {
	IsAuthenticated() bool
	IsAuthorized(role string) bool
}
