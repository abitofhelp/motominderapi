// Package response contains the response messages for the use cases.
package response

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// InsertMotorcycleResponse is a simple dto containing the response data from the InsertMotorcycleInteractor.
type InsertMotorcycleResponse struct {
	ID    int   `json:"id"`
	Error error `json:"error"`
}

// NewInsertMotorcycleResponse creates a new instance of a MotorcycleResponseMessage.
// Returns (nil, error) when there is an error, otherwise (InsertMotorcycleResponse, nil).
func NewInsertMotorcycleResponse(id int, err error) (*InsertMotorcycleResponse, error) {

	// We return a (nil, error) only when validation of the response message fails, not for whether the
	// response message indicates failure.

	motorcycleResponse := &InsertMotorcycleResponse{
		ID:    id,
		Error: err,
	}

	msgErr := motorcycleResponse.Validate()

	// If were have a response message with a failure and validation failed, we will wrap the original error with the validation error.
	if motorcycleResponse.Error != nil && msgErr != nil {
		return nil, errors.Wrap(msgErr, motorcycleResponse.Error.Error())
	}

	// If were have a response message that indicates success, but validation failed, we will return the validation error.
	if motorcycleResponse.Error == nil && msgErr != nil {
		return nil, msgErr
	}

	// If were have a response message that failed, but validation was successful, we will return response.
	if motorcycleResponse.Error != nil && msgErr == nil {
		return motorcycleResponse, nil
	}

	// Otherwise, all okay
	return motorcycleResponse, nil
}

// Validate verifies that a InsertMotorcycleResponse's fields contain valid data.
// Returns nil if the InsertMotorcycleResponse contains valid data, otherwise an error.
func (response InsertMotorcycleResponse) Validate() error {
	return validation.ValidateStruct(&response,
		// ID is required and it must be non-zero
		validation.Field(&response.ID, validation.Required))
}
