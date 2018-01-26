// Package interactor contains use cases, which contain the application specific business rules.
// Interactors encapsulate and implement all of the use cases of the system.  They orchestrate the
// flow of data to and from the entity, and can rely on their business rules to achieve the goals
// of the use case.  They do not have any dependencies, and are totally isolated from things like
// a database, UI or special frameworks, which exist in the outer rings.  They Will almost certainly
// require refactoring if details of the use case requirements change.
package interactor

import (
	"testing"

	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/stretchr/testify/assert"
)

// TestListMotorcyclesInteractor_MotorcycleRepositoryIsNil verifies that a nil motorcycle repository fails properly.
func TestListMotorcyclesInteractor_MotorcycleRepositoryIsNil(t *testing.T) {

	// ARRANGE
	roles := make(map[authorizationrole.AuthorizationRole]bool)
	authService, _ := security.NewAuthService(true, roles)

	// ACT
	_, err := NewListMotorcyclesInteractor(nil, authService)

	// ASSERT
	assert.NotNil(t, err)
}

// TestListMotorcyclesInteractor_MotorcycleRepositoryIsNil verifies that a nil authorization service fails properly.
func TestListMotorcyclesInteractor_AuthServiceIsNil(t *testing.T) {

	// ARRANGE
	repo, _ := repository.NewMotorcycleRepository()

	// ACT
	_, err := NewListMotorcyclesInteractor(repo, nil)

	// ASSERT
	assert.NotNil(t, err)
}

// TestListMotorcyclesInteractor_NotAuthenticated verifies that a non-authenticated user fails properly.
func TestListMotorcyclesInteractor_NotAuthenticated(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(false, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewListMotorcyclesRequest()
	interactor, _ := NewListMotorcyclesInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.Nil(t, response.Motorcycles)
	assert.NotNil(t, response.Error)
}

// TestListMotorcyclesInteractor_NotAuthorized verifies that an authenticated user lacking an authorization role cannot insert a motorcycle.
func TestListMotorcyclesInteractor_NotAuthorized(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: false,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewListMotorcyclesRequest()
	interactor, _ := NewListMotorcyclesInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.Nil(t, response.Motorcycles)
	assert.NotNil(t, response.Error)
}

// TestListMotorcyclesInteractor_EmptyList gets an empty list of motorcycles from the repository.
func TestListMotorcyclesInteractor_EmptyList(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewListMotorcyclesRequest()
	interactor, _ := NewListMotorcyclesInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(motorcycleRequest)

	// ASSERT
	assert.Len(t, response.Motorcycles, 0)
}

// TestListMotorcyclesInteractor gets a non-empty list of motorcycles from the repository.
func TestListMotorcyclesInteractor_NotEmptyList(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	listRequest, _ := request.NewListMotorcyclesRequest()
	listInteractor, _ := NewListMotorcyclesInteractor(repo, authService)

	insertInteractor, _ := NewInsertMotorcycleInteractor(repo, authService)
	insertRequest1, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	insertRequest2, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2009, "01234567999923456")
	insertInteractor.Handle(insertRequest1)
	insertInteractor.Handle(insertRequest2)

	// ACT
	response, _ := listInteractor.Handle(listRequest)

	// ASSERT
	assert.Len(t, response.Motorcycles, 2)
}
