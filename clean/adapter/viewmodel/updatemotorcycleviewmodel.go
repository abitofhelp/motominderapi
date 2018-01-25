// Package viewmodel translates a response message into a view model.
package viewmodel

import (
	"github.com/pkg/errors"

	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
	"github.com/go-ozzo/ozzo-validation"
)

// UpdateMotorcycleViewModel translates a UpdateMotorcycleResponse to a UpdateMotorcycleViewModel.
// by the Configuration ring.
type UpdateMotorcycleViewModel struct {
	ID      typedef.ID `json:"id"`
	Message string     `json:"message"`
	Error   error      `json:"error"`
}

// NewUpdateMotorcycleViewModel creates a new instance of a UpdateMotorcycleViewModel.
// Returns an (instance of UpdateMotorcycleViewModel, nil) on success, otherwise (nil, error)
func NewUpdateMotorcycleViewModel(id typedef.ID, message string, err error) (*UpdateMotorcycleViewModel, error) {

	viewModel := &UpdateMotorcycleViewModel{
		ID:      id,
		Message: message,
		Error:   err,
	}

	msgErr := viewModel.Validate()
	// If we have a response message with a failure and validation failed, we will wrap the original error with the validation error.
	if viewModel.Error != nil && msgErr != nil {
		return nil, errors.Wrap(viewModel.Error, msgErr.Error())
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

// Validate verifies that a UpdateMotorcycleViewModel's fields contain valid data.
// Returns (an instance of UpdateMotorcycleViewModel, nil) on success, otherwise (nil, error).
func (viewmodel UpdateMotorcycleViewModel) Validate() error {
	return validation.ValidateStruct(&viewmodel,
		// ID is required and it must be non-zero
		validation.Field(&viewmodel.ID, validation.Required, validation.Min(constant.MinEntityID)),
		// Message is required and it cannot be empty or nil.
		validation.Field(&viewmodel.Message, validation.Required, validation.NilOrNotEmpty),
	)
}
