// Package viewmodel translates a response message into a view model.
package viewmodel

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/dto"
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// ListMotorcyclesViewModel translates a ListMotorcyclesResponse to a ListMotorcyclesViewModel.
// by the Configuration ring.
type ListMotorcyclesViewModel struct {
	Motorcycles []dto.MotorcycleDto `json:"motorcycles"`
	Message     string              `json:"message"`
	Error       error               `json:"error"`
}

// NewListMotorcyclesViewModel creates a new instance of a ListMotorcyclesViewModel.
// Returns an (instance of ListMotorcyclesViewModel, nil) on success, otherwise (nil, error)
func NewListMotorcyclesViewModel(motorcycles []entity.Motorcycle, message string, err error) (*ListMotorcyclesViewModel, error) {
	// Ensure that we create an empty slice rather than the default for []entity.Motorcycle, which is a null pointer.
	motorcycleDtos := make([]dto.MotorcycleDto, 0)

	for i := 0; i < len(motorcycles); i++ {
		motorcycle := &dto.MotorcycleDto{
			ID:          motorcycles[i].ID,
			Make:        motorcycles[i].Make,
			Model:       motorcycles[i].Model,
			Year:        motorcycles[i].Year,
			Vin:         motorcycles[i].Vin,
			CreatedUtc:  motorcycles[i].CreatedUtc,
			ModifiedUtc: motorcycles[i].ModifiedUtc,
		}

		motorcycleDtos = append(motorcycleDtos, *motorcycle)

	}

	viewModel := &ListMotorcyclesViewModel{
		Motorcycles: motorcycleDtos,
		Message:     message,
		Error:       err,
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

// Validate verifies that a ListMotorcyclesViewModel's fields contain valid data.
// Returns (an instance of ListMotorcyclesViewModel, nil) on success, otherwise (nil, error).
func (viewmodel ListMotorcyclesViewModel) Validate() error {
	return validation.ValidateStruct(&viewmodel,
		// Motorcycles can be empty, but not nil
		validation.Field(&viewmodel.Motorcycles, validation.NotNil),

		// Message is required and it cannot be empty or nil.
		validation.Field(&viewmodel.Message, validation.NilOrNotEmpty),
	)
}
