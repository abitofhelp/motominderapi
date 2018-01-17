// Package viewModels translates a response message into a view model.
package viewmodel

import (
	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/abitofhelp/motominderapi/clean/usecase/responsemessage"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// InsertMotorcycleResponseMessageViewModel translates a InsertMotorcycleResponseMessage to a InsertMotorcycleResponseMessageViewModel.
// by the Configuration ring.
type InsertMotorcycleResponseMessageViewModel struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

// NewInsertMotorcycleResponseMessageViewModel creates a new instance of a InsertMotorcycleResponseMessageViewModel.
// Returns an (instance of InsertMotorcycleResponseMessageViewModel, nil) on success, otherwise (nil, error)
func NewInsertMotorcycleResponseMessageViewModel(id int, message string, err error) (*InsertMotorcycleResponseMessageViewModel, error) {

	viewModel := &InsertMotorcycleResponseMessageViewModel{
		ID:      id,
		Message: message,
		Error:   err,
	}

	msgErr := viewModel.Validate()
	if msgErr != nil {
		// We had an error validating the response message,
		// so we will wrap the original error with the validation error.
		return nil, errors.Wrap(msgErr, viewModel.Error.Error())
	}

	// All okay
	return viewModel, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of InsertMotorcycleResponseMessageViewModel, nil) on success, otherwise (nil, error)
func (insertMotorcycleResponseMessageViewModel *InsertMotorcycleResponseMessageViewModel) Handle(responseMessage *responsemessage.InsertMotorcycleResponseMessage) (*InsertMotorcycleResponseMessageViewModel, error) {
	if responseMessage.Error != nil {
		return NewInsertMotorcycleResponseMessageViewModel(constant.InvalidEntityID, responseMessage.Error.Error(), responseMessage.Error)
	}

	return NewInsertMotorcycleResponseMessageViewModel(responseMessage.ID, "Successfully inserted a new motorcycle.", nil)
}

// Validate verifies that a InsertMotorcycleResponseMessageViewModel's fields contain valid data.
// Returns (an instance of InsertMotorcycleResponseMessageViewModel, nil) on success, otherwise (nil, error).
func (insertMotorcycleResponseMessageViewModel InsertMotorcycleResponseMessageViewModel) Validate() error {
	return validation.ValidateStruct(&insertMotorcycleResponseMessageViewModel,
		// ID is required and it must be non-zero
		validation.Field(&insertMotorcycleResponseMessageViewModel.ID, validation.Required, validation.Min(constant.MinEntityID)),
		// Message is required and it cannot be empty or nil.
		validation.Field(&insertMotorcycleResponseMessageViewModel.Message, validation.Required, validation.NilOrNotEmpty),
	)
}
