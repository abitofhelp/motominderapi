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
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestUpdateMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil motorcycle repository fails properly.
func TestUpdateMotorcycleInteractor_MotorcycleRepositoryIsNil(t *testing.T) {

	// ARRANGE
	roles := make(map[authorizationrole.AuthorizationRole]bool)
	authService, _ := security.NewAuthService(true, roles)

	// ACT
	_, err := NewUpdateMotorcycleInteractor(nil, authService)

	// ASSERT
	assert.NotNil(t, err)
}

// TestUpdateMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil authorization service fails properly.
func TestUpdateMotorcycleInteractor_AuthServiceIsNil(t *testing.T) {

	// ARRANGE
	repo, _ := repository.NewMotorcycleRepository()

	// ACT
	_, err := NewUpdateMotorcycleInteractor(repo, nil)

	// ASSERT
	assert.NotNil(t, err)
}

// TestUpdateMotorcycleInteractor_NotAuthenticated verifies that a non-authenticated user fails properly.
func TestUpdateMotorcycleInteractor_NotAuthenticated(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(false, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycle, _ := entity.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")
	motorcycle.ID = 123
	motorcycleRequest, _ := request.NewUpdateMotorcycleRequest(123, motorcycle)
	interactor, _ := NewUpdateMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.NotNil(t, response.Error)
}

// TestUpdateMotorcycleInteractor_NotAuthorized verifies that an authenticated user lacking an authorization role cannot insert a motorcycle.
func TestUpdateMotorcycleInteractor_NotAuthorized(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: false,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycle, _ := entity.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")
	motorcycle.ID = 123
	motorcycleRequest, _ := request.NewUpdateMotorcycleRequest(123, motorcycle)
	interactor, _ := NewUpdateMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.NotNil(t, response.Error)
}

// TestUpdateMotorcycleInteractor_Update deletes a motorcycle in the repository.
func TestUpdateMotorcycleInteractor_Update(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()

	// Add a motorcycle so we can update it.
	insertInteractor, _ := NewInsertMotorcycleInteractor(repo, authService)
	insertRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	insertResponse, _ := insertInteractor.Handle(insertRequest)

	// Get the new motorcycle and change its vin.
	motorcycle, _, _ := repo.FindByID(insertResponse.ID)
	vin := "65432109876543210"
	motorcycle.Vin = vin
	updateRequest, _ := request.NewUpdateMotorcycleRequest(insertResponse.ID, motorcycle)
	updateInteractor, _ := NewUpdateMotorcycleInteractor(repo, authService)

	// ACT
	_, err := updateInteractor.Handle(updateRequest)
	motorcycle, _, _ = repo.FindByID(updateRequest.ID)

	// ASSERT
	assert.Nil(t, err)
	assert.True(t, motorcycle.Vin == vin)
}

// TestUpdateMotorcycleInteractor_NotExist attempts to delete a motorcycle in the repository that does not exist.
func TestUpdateMotorcycleInteractor_NotExist(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()

	// Add a motorcycle so we can seek it.
	insertInteractor, _ := NewInsertMotorcycleInteractor(repo, authService)
	insertRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	insertResponse, _ := insertInteractor.Handle(insertRequest)

	motorcycle, _, _ := repo.FindByID(insertResponse.ID)
	updateRequest, _ := request.NewUpdateMotorcycleRequest(123, motorcycle)
	updateInteractor, _ := NewUpdateMotorcycleInteractor(repo, authService)

	// ACT
	updateResponse, _ := updateInteractor.Handle(updateRequest)

	// ASSERT
	assert.NotNil(t, updateResponse.Error)
}
