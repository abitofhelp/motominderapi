// Package presenter performs the translation of a response message into a view model.
package presenter

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/viewmodel"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
	"github.com/go-ozzo/ozzo-validation"
)

// InsertMotorcyclePresenter translates the response message from the InsertMotorcycleInteractor to a view model.
type InsertMotorcyclePresenter struct {
}

// NewInsertMotorcyclePresenter creates a new instance of a InsertMotorcyclePresenter.
// Returns (instance of InsertMotorcyclePresenter, nil) on success, otherwise (nil, error).
func NewInsertMotorcyclePresenter() (*InsertMotorcyclePresenter, error) {

	presenter := &InsertMotorcyclePresenter{}

	// All okay
	return presenter, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of InsertMotorcycleViewModel, nil) on success, otherwise (nil, error)
func (presenter *InsertMotorcyclePresenter) Handle(responseMessage *response.InsertMotorcycleResponse) (*viewmodel.InsertMotorcycleViewModel, error) {
	if responseMessage.Error != nil {
		return viewmodel.NewInsertMotorcycleViewModel(responseMessage.ID, "Failed to create the new motorcycle.", responseMessage.Error)
	}

	return viewmodel.NewInsertMotorcycleViewModel(responseMessage.ID, "Successfully created the new motorcycle.", responseMessage.Error)
}

// Validate verifies that a InsertMotorcyclePresenter's fields contain valid data.
// Returns (an instance of InsertMotorcyclePresenter, nil) on success, otherwise (nil, error)
func (presenter InsertMotorcyclePresenter) Validate() error {
	return validation.ValidateStruct(&presenter)
}
