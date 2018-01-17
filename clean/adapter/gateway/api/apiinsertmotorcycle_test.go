// Package api contains the restful web service.
package api

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/adapter/presenter"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/abitofhelp/motominderapi/clean/usecase/interactor"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestApi_InsertMotorcycle verifies that a new motorcycle can be successfully created.
func TestApi_InsertMotorcycle(t *testing.T) {

	// ARRANGE
	roles := map[enumeration.AuthorizationRole]bool{
		enumeration.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewInsertMotorcycleRequestMessage("Honda", "Shadow", 2006, "01234567890123456")
	motorcycleInteractor, _ := interactor.NewInsertMotorcycleInteractor(repo, authService)
	response, _ := motorcycleInteractor.Handle(motorcycleRequest)
	motorcyclePresenter, _ := presenter.NewInsertMotorcyclePresenter()

	// ACT
	viewModel, _ := motorcyclePresenter.Handle(response)

	// ASSERT
	assert.Nil(t, viewModel.Error)
}
