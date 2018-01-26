// Package presenter implements unit tests for UpdateMotorcycleResponseMessagePresentation.
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

// TestUpdateMotorcyclePresenter_Handle verifies that a response messages is translated into a proper view model.
func TestUpdateMotorcyclePresenter_Handle(t *testing.T) {

	// ARRANGE
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()

	// Insert a motorcycle so we can update it.
	insertRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	insertInteractor, _ := interactor.NewInsertMotorcycleInteractor(repo, authService)
	insertResponse, _ := insertInteractor.Handle(insertRequest)

	motorcycle, _, _ := repo.FindByID(insertResponse.ID)
	motorcycle.Vin = "65432109876543210"

	updateRequest, _ := request.NewUpdateMotorcycleRequest(insertResponse.ID, motorcycle)
	updateInteractor, _ := interactor.NewUpdateMotorcycleInteractor(repo, authService)
	updateResponse, _ := updateInteractor.Handle(updateRequest)
	updatePresenter, _ := NewUpdateMotorcyclePresenter()

	// ACT
	viewModel, _ := updatePresenter.Handle(updateResponse)

	// ASSERT
	assert.Nil(t, viewModel.Error)
}
