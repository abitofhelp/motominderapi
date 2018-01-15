// Package authServiceTests implements unit tests for the AuthService entity.
package authServiceTests

import (
	"github.com/abitofhelp/motominderapi/clean/adapters/gateways/security"
	"github.com/abitofhelp/motominderapi/clean/domain/enumerations"
	"github.com/stretchr/testify/assert"
	"testing"
)

// AuthService_IsAuthenticated verifies an authenticated user has been detected.
func TestAuthService_IsAuthenticated(t *testing.T) {

	// ARRANGE
	roles := make(map[enumerations.AuthorizationRole]bool)

	// ACT
	authService, _ := security.NewAuthService(true, roles)

	// ASSERT
	assert.True(t, authService.Authenticated)
}

// TestAuthService_IsNotAuthenticated verifies a user who has not been authenticated is detected.
func TestAuthService_IsNotAuthenticated(t *testing.T) {

	// ARRANGE
	roles := make(map[enumerations.AuthorizationRole]bool)

	// ACT
	authService, _ := security.NewAuthService(false, roles)

	// ASSERT
	assert.False(t, authService.Authenticated)
}

// TestAuthService_HasARole verifies that a user is authorized for a specific role.
func TestAuthService_HasARole(t *testing.T) {

	// ARRANGE
	roles := map[enumerations.AuthorizationRole]bool{
		enumerations.AdminAuthorizationRole: true,
	}

	// ACT
	authService, _ := security.NewAuthService(true, roles)

	// ASSERT
	assert.True(t, authService.IsAuthorized(enumerations.AdminAuthorizationRole))
}

// TestAuthService_DoesNotHaveRole verifies that a user is not authorized for a specific role.
func TestAuthService_DoesNotHaveRole(t *testing.T) {

	// ARRANGE
	roles := map[enumerations.AuthorizationRole]bool{
		enumerations.AdminAuthorizationRole: false,
	}

	// ACT
	authService, _ := security.NewAuthService(true, roles)

	// ASSERT
	assert.False(t, authService.IsAuthorized(enumerations.AdminAuthorizationRole))
}

// TestAuthService_DoesNotHaveRoleInMap verifies that a user does not have a specific authorization for a role, which should always return false.
func TestAuthService_DoesNotHaveRoleInMap(t *testing.T) {

	// ARRANGE
	roles := make(map[enumerations.AuthorizationRole]bool)

	// ACT
	authService, _ := security.NewAuthService(true, roles)

	// ASSERT
	assert.False(t, authService.IsAuthorized(enumerations.AdminAuthorizationRole))
}
