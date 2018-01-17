// Package presenter implements unit tests for InsertMotorcycleResponseMessagePresentation.
package presenter

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/abitofhelp/motominderapi/clean/usecase/interactor"
	"github.com/abitofhelp/motominderapi/clean/usecase/requestmessage"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInsertMotorcycleResponseMessagePresentation_Handle verifies that a response messages is translated into a proper view model.
func TestInsertMotorcycleResponseMessagePresentation_Handle(t *testing.T) {

	// ARRANGE
	roles := map[enumeration.AuthorizationRole]bool{
		enumeration.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	request, _ := requestmessage.NewInsertMotorcycleRequestMessage("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := interactor.NewInsertMotorcycleInteractor(repo, authService)
	response, _ := interactor.Handle(request)
	presenter, _ := NewInsertMotorcycleResponseMessagePresenter()

	// ACT
	viewModel, _ := presenter.Handle(response)

	// ASSERT
	assert.Nil(t, viewModel.Error)
}
