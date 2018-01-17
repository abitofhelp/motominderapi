// Package responsemessage contains the response messages for the use cases.
package responsemessage

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// InsertMotorcycleResponseMessage is a simple dto containing the response data from the InsertMotorcycleInteractor.
type InsertMotorcycleResponseMessage struct {
	ID    int   `json:"id"`
	Error error `json:"error"`
}

// NewInsertMotorcycleResponseMessage creates a new instance of a MotorcycleResponseMessage.
// Returns (nil, error) when there is an error, otherwise (InsertMotorcycleResponseMessage, nil).
func NewInsertMotorcycleResponseMessage(id int, err error) (*InsertMotorcycleResponseMessage, error) {

	// We return a (nil, error) only when validation of the response message fails, not for whether the
	// response message indicates failure.

	responseMessage := &InsertMotorcycleResponseMessage{
		ID:    id,
		Error: err,
	}

	msgErr := responseMessage.Validate()

	// If were have a response message with a failure and validation failed, we will wrap the original error with the validation error.
	if responseMessage.Error != nil && msgErr != nil {
		return nil, errors.Wrap(msgErr, responseMessage.Error.Error())
	}

	// If were have a response message that indicates success, but validation failed, we will return the validation error.
	if responseMessage.Error == nil && msgErr != nil {
		return nil, msgErr
	}

	// If were have a response message that failed, but validation was successful, we will return response.
	if responseMessage.Error != nil && msgErr == nil {
		return responseMessage, nil
	}

	// Otherwise, all okay
	return responseMessage, nil
}

// Validate verifies that a InsertMotorcycleResponseMessage's fields contain valid data.
// Returns nil if the InsertMotorcycleResponseMessage contains valid data, otherwise an error.
func (insertMotorcycleResponseMessage InsertMotorcycleResponseMessage) Validate() error {
	return validation.ValidateStruct(&insertMotorcycleResponseMessage,
		// ID is required and it must be non-zero
		validation.Field(&insertMotorcycleResponseMessage.ID, validation.Required))
}
