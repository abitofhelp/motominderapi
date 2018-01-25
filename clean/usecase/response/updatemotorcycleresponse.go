// Package response contains the response messages for the use cases.
package response

import (
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/operationstatus"
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// UpdateMotorcycleResponse is a simple dto containing the response data from the UpdateMotorcycleInteractor.
type UpdateMotorcycleResponse struct {
	// ID will be set to the value that was requested to be updated.
	ID     typedef.ID                      `json:"id"`
	Status operationstatus.OperationStatus `json:"operationStatus"`
	Error  error                           `json:"error"`
}

// NewUpdateMotorcycleResponse creates a new instance of a UpdateMotorcycleResponse.
// Returns (nil, error) when there is an error, otherwise (UpdateMotorcycleResponse, nil).
func NewUpdateMotorcycleResponse(id typedef.ID, status operationstatus.OperationStatus, err error) (*UpdateMotorcycleResponse, error) {

	// We return a (nil, error) only when validation of the response message fails, not for whether the
	// response message indicates failure.

	motorcycleResponse := &UpdateMotorcycleResponse{
		ID:     id,
		Status: status,
		Error:  err,
	}

	msgErr := motorcycleResponse.Validate()

	// If we have a response message with a failure and validation failed, we will wrap the original error with the validation error.
	if motorcycleResponse.Error != nil && msgErr != nil {
		return nil, errors.Wrap(motorcycleResponse.Error, msgErr.Error())
	}

	// If we have a response message that indicates success, but validation failed, we will return the validation error.
	if motorcycleResponse.Error == nil && msgErr != nil {
		return nil, msgErr
	}

	// If we have a response message that failed, but validation was successful, we will return response.
	if motorcycleResponse.Error != nil && msgErr == nil {
		return motorcycleResponse, nil
	}

	// Otherwise, all okay
	return motorcycleResponse, nil
}

// Validate verifies that a UpdateMotorcycleResponse's fields contain valid data.
// Returns nil if the UpdateMotorcycleResponse contains valid data, otherwise an error.
func (response UpdateMotorcycleResponse) Validate() error {
	return validation.ValidateStruct(&response,
		// ID is required and it must be non-zero
		validation.Field(&response.ID, validation.Required, validation.Min(1)))
}
