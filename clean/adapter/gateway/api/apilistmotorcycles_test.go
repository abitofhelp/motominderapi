// Package api contains the restful web service.
package api

/*
// TestApi_ListMotorcycles verifies the retrieval of the list of motorcycles from the repository.
func TestApi_ListMotorcycles(t *testing.T) {

	// ARRANGE
	roles := map[enumeration.AuthorizationRole]bool{
		enumeration.AdminAuthorizationRole: true,
	}
	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repository.NewMotorcycleRepository()
	motorcycleRequest, _ := request.NewInsertMotorcycleRequest("Honda", "Shadow", 2006, "01234567890123456")
	motorcycleInteractor, _ := interactor.NewInsertMotorcycleInteractor(repo, authService)
	response, _ := motorcycleInteractor.Handle(motorcycleRequest)
	motorcyclePresenter, _ := presenter.NewInsertMotorcyclePresenter()

	// ACT
	viewModel, _ := motorcyclePresenter.Handle(response)

	// ASSERT
	assert.Nil(t, viewModel.Error)
}
*/
