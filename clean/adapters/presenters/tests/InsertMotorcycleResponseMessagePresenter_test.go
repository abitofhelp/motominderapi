// Package insertMotorcycleResponseMessagePresentationTests implements unit tests for InsertMotorcycleResponseMessagePresentation.
package insertMotorcycleResponseMessagePresentationTests

import (
	"github.com/abitofhelp/motominderapi/clean/adapters/gateways/repositories"
	"github.com/abitofhelp/motominderapi/clean/adapters/gateways/security"
	"github.com/abitofhelp/motominderapi/clean/adapters/presenters"
	"github.com/abitofhelp/motominderapi/clean/domain/enumerations"
	"github.com/abitofhelp/motominderapi/clean/usecases/interactors"
	"github.com/abitofhelp/motominderapi/clean/usecases/requestmessages"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInsertMotorcycleResponseMessagePresentation_Handle verifies that a response messages is translated into a proper view model.
func TestInsertMotorcycleResponseMessagePresentation_Handle(t *testing.T) {

	// ARRANGE
	roles := map[enumerations.AuthorizationRole]bool{
		enumerations.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repositories.NewMotorcycleRepository()
	request, _ := requestmessages.NewInsertMotorcycleRequestMessage("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := interactors.NewInsertMotorcycleInteractor(repo, authService)
	response, _ := interactor.Handle(request)
	presenter, _ := presenters.NewInsertMotorcycleResponseMessagePresenter()

	// ACT
	viewModel, _ := presenter.Handle(response)

	// ASSERT
	assert.Nil(t, viewModel.Error)
}
