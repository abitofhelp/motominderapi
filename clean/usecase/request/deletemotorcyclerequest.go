// Package request contains the request messages for the use cases.
package request

import (
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/go-ozzo/ozzo-validation"
	errors "github.com/pjebs/jsonerror"
)

// DeleteMotorcycleRequest is a simple dto containing the required data for the DeleteMotorcycleInteractor.
type DeleteMotorcycleRequest struct {
	ID int `json:"id"`
}

// NewDeleteMotorcycleRequest creates a new instance of a DeleteMotorcycleRequest.
// Returns (nil, error) when there is an error, otherwise (DeleteMotorcycleRequest, nil).
func NewDeleteMotorcycleRequest(id int) (*DeleteMotorcycleRequest, error) {

	motorcycleRequest := &DeleteMotorcycleRequest{
		ID: id,
	}

	err := motorcycleRequest.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return motorcycleRequest, nil
}

// Validate verifies that a DeleteMotorcycleRequest's fields contain valid data.
// Returns (an instance of DeleteMotorcycleRequest, nil) on success, otherwise (nil, error)
func (request DeleteMotorcycleRequest) Validate() error {
	err := validation.ValidateStruct(&request,
		// ID is required and it must be greater than 0.
		validation.Field(&request.ID, validation.Required, validation.Min(1)))

	if err != nil {
		return errors.New(enumeration.StatusInternalServerError, enumeration.StatusText(enumeration.StatusInternalServerError), err.Error())
	}

	return nil
}
