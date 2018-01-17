// Package api contains the restful web service.
package api

/*
import (
	"testing"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/usecase/requestmessage"
	"github.com/abitofhelp/motominderapi/clean/usecase/interactor"
	"github.com/abitofhelp/motominderapi/clean/adapter/presenter"
	"github.com/stretchr/testify/assert"
)

// TestApi_InsertMotorcycle verifies that a new motorcycle can be successfully created.
func TestApi_InsertMotorcycle(t *testing.T) {

	// ARRANGE
	roles := map[enumeration.AuthorizationRole]bool{
		enumeration.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	request, _ := requestmessage.NewInsertMotorcycleRequestMessage("Honda", "Shadow", 2006, "01234567890123456")
	interactor, _ := interactor.NewInsertMotorcycleInteractor(repo, authService)
	response, _ := interactor.Handle(request)
	presenter, _ := presenter.NewInsertMotorcycleResponseMessagePresenter()

	// ACT
	viewModel, _ := presenter.Handle(response)

	// ASSERT
	assert.Nil(t, viewModel.Error)
}
*/
