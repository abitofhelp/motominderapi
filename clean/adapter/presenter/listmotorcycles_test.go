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

// TestListMotorcyclesPresenter_Handle verifies that a response messages is translated into a proper view model.
func TestListMotorcyclesPresenter_Handle(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	insertRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	insertInteractor, _ := interactor.NewInsertMotorcycleInteractor(repo, authService)
	insertInteractor.Handle(insertRequest)

	listRequest, _ := request.NewListMotorcyclesRequest()
	listInteractor, _ := interactor.NewListMotorcyclesInteractor(repo, authService)
	listResponse, _ := listInteractor.Handle(listRequest)
	presenter, _ := NewListMotorcyclesPresenter()

	// ACT
	viewModel, _ := presenter.Handle(listResponse)

	// ASSERT
	assert.Nil(t, viewModel.Error)
}
