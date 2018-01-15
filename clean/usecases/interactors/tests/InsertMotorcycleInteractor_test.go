// Package insertMotorcycleInteractorTests implements unit tests for InsertMotorcycleInteractorTests.
package insertMotorcycleInteractorTests

import (
	"github.com/abitofhelp/motominderapi/clean/adapters/gateways/repositories"
	"github.com/abitofhelp/motominderapi/clean/adapters/gateways/security"
	"github.com/abitofhelp/motominderapi/clean/domain/enumerations"
	"github.com/abitofhelp/motominderapi/clean/usecases/interactors"
	"github.com/abitofhelp/motominderapi/clean/usecases/requestmessages"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInsertMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil motorcycle repository fails properly.
func TestInsertMotorcycleInteractor_MotorcycleRepositoryIsNil(t *testing.T) {

	// ARRANGE
	roles := make(map[enumerations.AuthorizationRole]bool)
	authService, _ := security.NewAuthService(true, roles)

	// ACT
	_, err := interactors.NewInsertMotorcycleInteractor(nil, authService)

	// ASSERT
	assert.NotNil(t, err)
}

// TestInsertMotorcycleInteractor_MotorcycleRepositoryIsNil verifies that a nil authorization service fails properly.
func TestInsertMotorcycleInteractor_AuthServiceIsNil(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()

	// ACT
	_, err := interactors.NewInsertMotorcycleInteractor(repo, nil)

	// ASSERT
	assert.NotNil(t, err)
}

// TestInsertMotorcycleInteractor_NotAuthenticated verifies that a non-authenticated user fails properly.
func TestInsertMotorcycleInteractor_NotAuthenticated(t *testing.T) {

	// ARRANGE
	roles := make(map[enumerations.AuthorizationRole]bool)
	authService, _ := security.NewAuthService(false, roles)
	repo, _ := repositories.NewMotorcycleRepository()
	request, _ := requestmessages.NewInsertMotorcycleRequestMessage("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := interactors.NewInsertMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(request)

	// ASSERT
	assert.NotNil(t, response.Error)
}

// TestInsertMotorcycleInteractor_NotAuthorized verifies that an authenticated user lacking an authorization role cannot insert a motorcycle.
func TestInsertMotorcycleInteractor_NotAuthorized(t *testing.T) {

	// ARRANGE
	roles := map[enumerations.AuthorizationRole]bool{
		enumerations.AdminAuthorizationRole: false,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repositories.NewMotorcycleRepository()
	request, _ := requestmessages.NewInsertMotorcycleRequestMessage("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := interactors.NewInsertMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(request)

	// ASSERT
	assert.True(t, response.ID == -1)
	assert.NotNil(t, response.Error)
}

// TestInsertMotorcycleInteractor_Insert inserts a new motorcycle into the repository.
func TestInsertMotorcycleInteractor_Insert(t *testing.T) {

	// ARRANGE
	roles := map[enumerations.AuthorizationRole]bool{
		enumerations.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repositories.NewMotorcycleRepository()
	request, _ := requestmessages.NewInsertMotorcycleRequestMessage("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := interactors.NewInsertMotorcycleInteractor(repo, authService)

	// ACT
	response, _ := interactor.Handle(request)

	// ASSERT
	assert.True(t, response.ID > 0)
	assert.Nil(t, response.Error)
}

// TestInsertMotorcycleInteractor_Insert_VinAlreadyExists verifies that a duplicate motorcycle will not be created.
func TestInsertMotorcycleInteractor_Insert_VinAlreadyExists(t *testing.T) {

	// ARRANGE
	roles := map[enumerations.AuthorizationRole]bool{
		enumerations.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repositories.NewMotorcycleRepository()
	request, _ := requestmessages.NewInsertMotorcycleRequestMessage("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := interactors.NewInsertMotorcycleInteractor(repo, authService)
	interactor.Handle(request)

	// ACT
	response, _ := interactor.Handle(request)

	// ASSERT
	assert.NotNil(t, response.Error)
}
