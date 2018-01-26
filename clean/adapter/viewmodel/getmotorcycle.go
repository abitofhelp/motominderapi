// Package viewmodel translates a response message into a view model.
package viewmodel

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/dto"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// GetMotorcycleViewModel translates a GetMotorcycleResponse to a GetMotorcycleViewModel.
// by the Configuration ring.
type GetMotorcycleViewModel struct {
	Motorcycle *dto.MotorcycleDto `json:"motorcycle"`
	Message    string             `json:"message"`
	Error      error              `json:"error"`
}

// NewGetMotorcycleViewModel creates a new instance of a GetMotorcycleViewModel.
// Returns an (instance of GetMotorcycleViewModel, nil) on success, otherwise (nil, error)
func NewGetMotorcycleViewModel(motorcycle *dto.MotorcycleDto, message string, err error) (*GetMotorcycleViewModel, error) {

	viewModel := &GetMotorcycleViewModel{
		Motorcycle: motorcycle,
		Message:    message,
		Error:      err,
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

// Validate verifies that a GetMotorcycleViewModel's fields contain valid data.
// Returns (an instance of GetMotorcycleViewModel, nil) on success, otherwise (nil, error).
func (viewmodel GetMotorcycleViewModel) Validate() error {
	return validation.ValidateStruct(&viewmodel,
		// Motorcycle can be empty, but not nil
		validation.Field(&viewmodel.Motorcycle, validation.NotNil),

		// Message is required and it cannot be empty or nil.
		validation.Field(&viewmodel.Message, validation.NilOrNotEmpty),
	)
}
