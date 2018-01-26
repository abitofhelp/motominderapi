// Package presenter implements unit tests for GetMotorcycleResponseMessagePresentation.
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

// TestGetMotorcyclePresenter_Handle verifies that a response messages is translated into a proper view model.
func TestGetMotorcyclePresenter_Handle(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()

	// Insert a motorcycle so we can get it.
	insertRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	insertInteractor, _ := interactor.NewInsertMotorcycleInteractor(repo, authService)
	insertResponse, _ := insertInteractor.Handle(insertRequest)

	getRequest, _ := request.NewGetMotorcycleRequest(insertResponse.ID)
	getInteractor, _ := interactor.NewGetMotorcycleInteractor(repo, authService)
	getResponse, _ := getInteractor.Handle(getRequest)
	getPresenter, _ := NewGetMotorcyclePresenter()

	// ACT
	viewModel, _ := getPresenter.Handle(getResponse)

	// ASSERT
	assert.Nil(t, viewModel.Error)
	assert.True(t, insertResponse.ID == viewModel.Motorcycle.ID)
}
