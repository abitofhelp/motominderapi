// Package presenter implements unit tests for InsertMotorcycleResponseMessagePresentation.
package presenter

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/usecase/interactor"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInsertMotorcyclePresenter_Handle verifies that a response messages is translated into a proper view model.
func TestInsertMotorcyclePresenter_Handle(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	motorcycleInteractor, _ := interactor.NewInsertMotorcycleInteractor(repo, authService)
	response, _ := motorcycleInteractor.Handle(motorcycleRequest)
	presenter, _ := NewInsertMotorcyclePresenter()

	// ACT
	viewModel, _ := presenter.Handle(response)

	// ASSERT
	assert.Nil(t, viewModel.Error)
}
