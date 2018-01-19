// Package response contains the response messages for the use cases.
package response

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// ListMotorcyclesResponse is a simple dto containing the response data from the ListMotorcyclesInteractor.
type ListMotorcyclesResponse struct {
	Motorcycles []entity.Motorcycle
	Error       error `json:"error"`
}

// NewListMotorcyclesResponse creates a new instance of a ListMotorcyclesResponse.
// Returns (nil, error) when there is an error, otherwise (ListMotorcyclesResponse, nil).
func NewListMotorcyclesResponse(motorcycles []entity.Motorcycle, err error) (*ListMotorcyclesResponse, error) {

	// We return a (nil, error) only when validation of the response message fails, not for whether the
	// response message indicates failure.

	listResponse := &ListMotorcyclesResponse{
		Motorcycles: motorcycles,
		Error:       err,
	}

	msgErr := listResponse.Validate()

	// If we have a response message with a failure and validation failed, we will wrap the original error with the validation error.
	if listResponse.Error != nil && msgErr != nil {
		return nil, errors.Wrap(msgErr, listResponse.Error.Error())
	}

	// If we have a response message that indicates success, but validation failed, we will return the validation error.
	if listResponse.Error == nil && msgErr != nil {
		return nil, msgErr
	}

	// If we have a response message that failed, but validation was successful, we will return response.
	if listResponse.Error != nil && msgErr == nil {
		return listResponse, nil
	}

	// Otherwise, all okay
	return listResponse, nil
}

// Validate verifies that a ListMotorcyclesResponse's fields contain valid data.
// Returns nil if the ListMotorcyclesResponse contains valid data, otherwise an error.
func (response ListMotorcyclesResponse) Validate() error {
	return validation.ValidateStruct(&response,
		// Motorcycles can be empty, but not nil
		validation.Field(&response.Motorcycles, validation.NotNil))
}
