// Package security implements unit tests for the AuthService entity.
package security

import (
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/stretchr/testify/assert"
	"testing"
)

// AuthService_IsAuthenticated verifies an authenticated user has been detected.
func TestAuthService_IsAuthenticated(t *testing.T) {

	// ARRANGE
	roles := make(map[authorizationrole.AuthorizationRole]bool)

	// ACT
	authService, _ := NewAuthService(true, roles)

	// ASSERT
	assert.True(t, authService.Authenticated)
}

// TestAuthService_IsNotAuthenticated verifies a user who has not been authenticated is detected.
func TestAuthService_IsNotAuthenticated(t *testing.T) {

	// ARRANGE
	roles := make(map[authorizationrole.AuthorizationRole]bool)

	// ACT
	authService, _ := NewAuthService(false, roles)

	// ASSERT
	assert.False(t, authService.Authenticated)
}

// TestAuthService_HasARole verifies that a user is authorized for a specific role.
func TestAuthService_HasARole(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}

	// ACT
	authService, _ := NewAuthService(true, roles)

	// ASSERT
	assert.True(t, authService.IsAuthorized(authorizationrole.AdminAuthorizationRole))
}

// TestAuthService_DoesNotHaveRole verifies that a user is not authorized for a specific role.
func TestAuthService_DoesNotHaveRole(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: false,
	}

	// ACT
	authService, _ := NewAuthService(true, roles)

	// ASSERT
	assert.False(t, authService.IsAuthorized(authorizationrole.AdminAuthorizationRole))
}

// TestAuthService_DoesNotHaveRoleInMap verifies that a user does not have a specific authorization for a role, which should always return false.
func TestAuthService_DoesNotHaveRoleInMap(t *testing.T) {

	// ARRANGE
	roles := make(map[authorizationrole.AuthorizationRole]bool)

	// ACT
	authService, _ := NewAuthService(true, roles)

	// ASSERT
	assert.False(t, authService.IsAuthorized(authorizationrole.AdminAuthorizationRole))
}
