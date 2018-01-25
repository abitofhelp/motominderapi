// Package presenter performs the translation of a response message into a view model.
package presenter

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/viewmodel"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
	"github.com/go-ozzo/ozzo-validation"
)

// UpdateMotorcyclePresenter translates the response message from the UpdateMotorcycleInteractor to a view model.
type UpdateMotorcyclePresenter struct {
}

// NewUpdateMotorcyclePresenter creates a new instance of a UpdateMotorcyclePresenter.
// Returns (instance of UpdateMotorcyclePresenter, nil) on success, otherwise (nil, error).
func NewUpdateMotorcyclePresenter() (*UpdateMotorcyclePresenter, error) {

	presenter := &UpdateMotorcyclePresenter{}

	// All okay
	return presenter, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of UpdateMotorcycleViewModel, nil) on success, otherwise (nil, error)
func (presenter *UpdateMotorcyclePresenter) Handle(responseMessage *response.UpdateMotorcycleResponse) (*viewmodel.UpdateMotorcycleViewModel, error) {
	if responseMessage.Error != nil {
		return viewmodel.NewUpdateMotorcycleViewModel(responseMessage.ID, "Failed to update the motorcycle.", responseMessage.Error)
	}

	return viewmodel.NewUpdateMotorcycleViewModel(responseMessage.ID, "Successfully updated the motorcycle.", responseMessage.Error)
}

// Validate verifies that a UpdateMotorcyclePresenter's fields contain valid data.
// Returns (an instance of UpdateMotorcyclePresenter, nil) on success, otherwise (nil, error)
func (presenter UpdateMotorcyclePresenter) Validate() error {
	return validation.ValidateStruct(&presenter)
}
