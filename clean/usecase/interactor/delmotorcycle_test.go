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

// TestDeleteMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil motorcycle repository fails properly.
func TestDeleteMotorcycleInteractor_MotorcycleRepositoryIsNil(t *testing.T) {

	// ARRANGE
	roles := make(map[authorizationrole.AuthorizationRole]bool)
	authService, _ := security.NewAuthService(true, roles)

	// ACT
	_, err := NewDeleteMotorcycleInteractor(nil, authService)

	// ASSERT
	assert.NotNil(t, err)
}

// TestDeleteMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil authorization service fails properly.
func TestDeleteMotorcycleInteractor_AuthServiceIsNil(t *testing.T) {

	// ARRANGE
	repo, _ := repository.NewMotorcycleRepository()

	// ACT
	_, err := NewDeleteMotorcycleInteractor(repo, nil)

	// ASSERT
	assert.NotNil(t, err)
}

// TestDeleteMotorcycleInteractor_NotAuthenticated verifies that a non-authenticated user fails properly.
func TestDeleteMotorcycleInteractor_NotAuthenticated(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(false, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewDeleteMotorcycleRequest(123)
	interactor, _ := NewDeleteMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.NotNil(t, response.Error)
}

// TestDeleteMotorcycleInteractor_NotAuthorized verifies that an authenticated user lacking an authorization role cannot insert a motorcycle.
func TestDeleteMotorcycleInteractor_NotAuthorized(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: false,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewDeleteMotorcycleRequest(123)
	interactor, _ := NewDeleteMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.NotNil(t, response.Error)
}

// TestDeleteMotorcycleInteractor_Delete deletes a motorcycle from the repository.
func TestDeleteMotorcycleInteractor_Delete(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()

	// Add a motorcycle so we can delete it.
	insertInteractor, _ := NewInsertMotorcycleInteractor(repo, authService)
	insertRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	insertResponse, _ := insertInteractor.Handle(insertRequest)

	deleteRequest, _ := request.NewDeleteMotorcycleRequest(insertResponse.ID)
	deleteInteractor, _ := NewDeleteMotorcycleInteractor(repo, authService)

	// ACT
	deleteResponse, _ := deleteInteractor.Handle(deleteRequest)

	// ASSERT
	assert.True(t, deleteResponse.ID == insertResponse.ID)
	assert.Nil(t, deleteResponse.Error)
}

// TestDeleteMotorcycleInteractor_NotExist attempts to delete a motorcycle from the repository that does not exist.
func TestDeleteMotorcycleInteractor_NotExist(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()

	deleteRequest, _ := request.NewDeleteMotorcycleRequest(123)
	deleteInteractor, _ := NewDeleteMotorcycleInteractor(repo, authService)

	// ACT
	deleteResponse, _ := deleteInteractor.Handle(deleteRequest)

	// ASSERT
	assert.True(t, deleteResponse.ID == 123)
	assert.NotNil(t, deleteResponse.Error)
}
