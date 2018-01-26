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
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/operationstatus"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil motorcycle repository fails properly.
func TestGetMotorcycleInteractor_MotorcycleRepositoryIsNil(t *testing.T) {

	// ARRANGE
	roles := make(map[authorizationrole.AuthorizationRole]bool)
	authService, _ := security.NewAuthService(true, roles)

	// ACT
	_, err := NewGetMotorcycleInteractor(nil, authService)

	// ASSERT
	assert.NotNil(t, err)
}

// TestGetMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil authorization service fails properly.
func TestGetMotorcycleInteractor_AuthServiceIsNil(t *testing.T) {

	// ARRANGE
	repo, _ := repository.NewMotorcycleRepository()

	// ACT
	_, err := NewGetMotorcycleInteractor(repo, nil)

	// ASSERT
	assert.NotNil(t, err)
}

// TestGetMotorcycleInteractor_NotAuthenticated verifies that a non-authenticated user fails properly.
func TestGetMotorcycleInteractor_NotAuthenticated(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(false, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewGetMotorcycleRequest(123)
	interactor, _ := NewGetMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.NotNil(t, response.Error)
}

// TestGetMotorcycleInteractor_NotAuthorized verifies that an authenticated user lacking an authorization role cannot insert a motorcycle.
func TestGetMotorcycleInteractor_NotAuthorized(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: false,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewGetMotorcycleRequest(123)
	interactor, _ := NewGetMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.NotNil(t, response.Error)
}

// TestGetMotorcycleInteractor_Get gets a motorcycle from the repository.
func TestGetMotorcycleInteractor_Get(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()

	// Add a motorcycle so we can get it.
	insertInteractor, _ := NewInsertMotorcycleInteractor(repo, authService)
	insertRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	insertResponse, _ := insertInteractor.Handle(insertRequest)

	getRequest, _ := request.NewGetMotorcycleRequest(insertResponse.ID)
	getInteractor, _ := NewGetMotorcycleInteractor(repo, authService)

	// ACT
	getResponse, _ := getInteractor.Handle(getRequest)

	// ASSERT
	assert.True(t, getResponse.Motorcycle.ID == insertResponse.ID)
	assert.Nil(t, getResponse.Error)
}

// TestGetMotorcycleInteractor_NotExist attempts to get a motorcycle from the repository that does not exist.
func TestGetMotorcycleInteractor_NotExist(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()

	getRequest, _ := request.NewGetMotorcycleRequest(123)
	getInteractor, _ := NewGetMotorcycleInteractor(repo, authService)

	// ACT
	getResponse, _ := getInteractor.Handle(getRequest)

	// ASSERT
	assert.True(t, getResponse.Status == operationstatus.NotFound)
}
