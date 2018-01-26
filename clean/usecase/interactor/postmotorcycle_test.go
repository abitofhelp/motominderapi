// Package interactor contains use cases, which contain the application specific business rules.
// Interactors encapsulate and implement all of the use cases of the system.  They orchestrate the
// flow of data to and from the entity, and can rely on their business rules to achieve the goals
// of the use case.  They do not have any dependencies, and are totally isolated from things like
// a database, UI or special frameworks, which exist in the outer rings.  They Will almost certainly
// require refactoring if details of the use case requirements change.
package interactor

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInsertMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil motorcycle repository fails properly.
func TestInsertMotorcycleInteractor_MotorcycleRepositoryIsNil(t *testing.T) {

	// ARRANGE
	roles := make(map[authorizationrole.AuthorizationRole]bool)
	authService, _ := security.NewAuthService(true, roles)

	// ACT
	_, err := NewInsertMotorcycleInteractor(nil, authService)

	// ASSERT
	assert.NotNil(t, err)
}

// TestInsertMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil authorization service fails properly.
func TestInsertMotorcycleInteractor_AuthServiceIsNil(t *testing.T) {

	// ARRANGE
	repo, _ := repository.NewMotorcycleRepository()

	// ACT
	_, err := NewInsertMotorcycleInteractor(repo, nil)

	// ASSERT
	assert.NotNil(t, err)
}

// TestInsertMotorcycleInteractor_NotAuthenticated verifies that a non-authenticated user fails properly.
func TestInsertMotorcycleInteractor_NotAuthenticated(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(false, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := NewInsertMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.True(t, response.ID == -1)
	assert.NotNil(t, response.Error)
}

// TestInsertMotorcycleInteractor_NotAuthorized verifies that an authenticated user lacking an authorization role cannot insert a motorcycle.
func TestInsertMotorcycleInteractor_NotAuthorized(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: false,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := NewInsertMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.True(t, response.ID == -1)
	assert.NotNil(t, response.Error)
}

// TestInsertMotorcycleInteractor_Insert inserts a new motorcycle into the repository.
func TestInsertMotorcycleInteractor_Insert(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := NewInsertMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.True(t, response.ID > 0)
	assert.Nil(t, response.Error)
}

// TestInsertMotorcycleInteractor_Insert_VinAlreadyExists verifies that a duplicate motorcycle will not be created.
func TestInsertMotorcycleInteractor_Insert_VinAlreadyExists(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := NewInsertMotorcycleInteractor(repo, authService)
	interactor.Handle(motorcycleRequest)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.NotNil(t, response.Error)
}
