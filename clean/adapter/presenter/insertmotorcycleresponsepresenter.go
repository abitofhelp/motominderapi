// Package presenter performs the translation of a response message into a view model.
package presenter

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/viewmodel"
	"github.com/abitofhelp/motominderapi/clean/usecase/responsemessage"
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
func (insertMotorcycleResponseMessagePresenter *InsertMotorcycleResponseMessagePresenter) Handle(responseMessage *responsemessage.InsertMotorcycleResponseMessage) (*viewmodel.InsertMotorcycleResponseMessageViewModel, error) {
	if responseMessage.Error != nil {
		return viewmodel.NewInsertMotorcycleResponseMessageViewModel(responseMessage.ID, "Failed to create the new motorcycle.", responseMessage.Error)
	}

	return viewmodel.NewInsertMotorcycleResponseMessageViewModel(responseMessage.ID, "Successfully created the new motorcycle.", responseMessage.Error)
}

// Validate verifies that a InsertMotorcycleResponseMessagePresenter's fields contain valid data.
// Returns (an instance of InsertMotorcycleResponseMessagePresenter, nil) on success, otherwise (nil, error)
func (insertMotorcycleResponseMessagePresenter InsertMotorcycleResponseMessagePresenter) Validate() error {
	return validation.ValidateStruct(&insertMotorcycleResponseMessagePresenter)
}
