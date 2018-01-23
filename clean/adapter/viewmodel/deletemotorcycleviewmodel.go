// Package viewmodel translates a response message into a view model.
package viewmodel

import (
	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/operationstatus"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
	"github.com/go-ozzo/ozzo-validation"
	errors "github.com/pjebs/jsonerror"
)

// DeleteMotorcycleViewModel translates a DeleteMotorcycleResponse to a DeleteMotorcycleViewModel.
// by the Configuration ring.
type DeleteMotorcycleViewModel struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

// NewDeleteMotorcycleViewModel creates a new instance of a DeleteMotorcycleViewModel.
// Returns an (instance of DeleteMotorcycleViewModel, nil) on success, otherwise (nil, error)
func NewDeleteMotorcycleViewModel(id int, message string, err error) (*DeleteMotorcycleViewModel, error) {

	viewModel := &DeleteMotorcycleViewModel{
		ID:      id,
		Message: message,
		Error:   err,
	}

	msgErr := viewModel.Validate()
	// If we have a response message with a failure and validation failed, we will wrap the original error with the validation error.
	if viewModel.Error != nil && msgErr != nil {
		ecol := errors.NewErrorCollection(errors.RejectDuplicates)
		ecol.AddErrors(viewModel.Error, msgErr)

		return nil, errors.New(viewModel.Error.(errors.JE).Code,
			operationstatus.StatusText(viewModel.Error.(errors.JE).Code), ecol.Error())
	}

	// If we have a response message that indicates success, but validation failed, we will return the validation error.
	if viewModel.Error == nil && msgErr != nil {
		return nil, msgErr
	}

	// If we have a response message that failed, but validation was successful, we will return response.
	if viewModel.Error != nil && msgErr == nil {
		return viewModel, nil
	}

	// Otherwise, all okay
	return viewModel, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of DeleteMotorcycleViewModel, nil) on success, otherwise (nil, error)
func (viewmodel *DeleteMotorcycleViewModel) Handle(responseMessage *response.DeleteMotorcycleResponse) (*DeleteMotorcycleViewModel, error) {
	if responseMessage.Error != nil {
		return NewDeleteMotorcycleViewModel(constant.InvalidEntityID, responseMessage.Error.Error(), responseMessage.Error)
	}

	return NewDeleteMotorcycleViewModel(responseMessage.ID, "Successfully deleted the motorcycle.", nil)
}

// Validate verifies that a DeleteMotorcycleViewModel's fields contain valid data.
// Returns (an instance of DeleteMotorcycleViewModel, nil) on success, otherwise (nil, error).
func (viewmodel DeleteMotorcycleViewModel) Validate() error {
	err := validation.ValidateStruct(&viewmodel,
		// ID is required and it must be non-zero
		validation.Field(&viewmodel.ID, validation.Required, validation.Min(constant.MinEntityID)),
		// Message is required and it cannot be empty or nil.
		validation.Field(&viewmodel.Message, validation.Required, validation.NilOrNotEmpty),
	)

	if err != nil {
		return errors.New(operationstatus.StatusInternalServerError, operationstatus.StatusText(operationstatus.StatusInternalServerError), err.Error())
	}

	return nil
}
