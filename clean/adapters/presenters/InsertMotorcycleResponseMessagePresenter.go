// Package presenters performs the translation of a response message into a view model.
package presenters

import (
	"github.com/abitofhelp/motominderapi/clean/adapters/viewModels"
	"github.com/abitofhelp/motominderapi/clean/usecases/responsemessages"
	"github.com/go-ozzo/ozzo-validation"
)

// InsertMotorcycleResponseMessagePresenter translates the response message from the InsertMotorcycleInteractor to a view model.
type InsertMotorcycleResponseMessagePresenter struct {
}

// NewInsertMotorcycleResponseMessagePresenter creates a new instance of a InsertMotorcycleResponseMessagePresenter.
// Returns (instance of InsertMotorcycleResponseMessagePresenter, nil) on success, otherwise (nil, error).
func NewInsertMotorcycleResponseMessagePresenter() (*InsertMotorcycleResponseMessagePresenter, error) {

	presenter := &InsertMotorcycleResponseMessagePresenter{}

	// All okay
	return presenter, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of InsertMotorcycleResponseMessageViewModel, nil) on success, otherwise (nil, error)
func (insertMotorcycleResponseMessagePresenter *InsertMotorcycleResponseMessagePresenter) Handle(responseMessage *responsemessages.InsertMotorcycleResponseMessage) (*viewModels.InsertMotorcycleResponseMessageViewModel, error) {
	if responseMessage.Error != nil {
		return viewModels.NewInsertMotorcycleResponseMessageViewModel(responseMessage.ID, "Failed to create the new motorcycle.", responseMessage.Error)
	}

	return viewModels.NewInsertMotorcycleResponseMessageViewModel(responseMessage.ID, "Successfully created the new motorcycle.", responseMessage.Error)
}

// Validate verifies that a InsertMotorcycleResponseMessagePresenter's fields contain valid data.
// Returns (an instance of InsertMotorcycleResponseMessagePresenter, nil) on success, otherwise (nil, error)
func (insertMotorcycleResponseMessagePresenter InsertMotorcycleResponseMessagePresenter) Validate() error {
	return validation.ValidateStruct(&insertMotorcycleResponseMessagePresenter)
}
