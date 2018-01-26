// Package request contains the request messages for the use cases.
package request

import (
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
	"github.com/go-ozzo/ozzo-validation"
)

// GetMotorcycleRequest is a simple dto containing the required data for the GetMotorcycleInteractor.
type GetMotorcycleRequest struct {
	ID typedef.ID `json:"id"`
}

// NewGetMotorcycleRequest creates a new instance of a GetMotorcycleRequest.
// Returns (nil, error) when there is an error, otherwise (GetMotorcycleRequest, nil).
func NewGetMotorcycleRequest(id typedef.ID) (*GetMotorcycleRequest, error) {

	motorcycleRequest := &GetMotorcycleRequest{
		ID: id,
	}

	err := motorcycleRequest.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return motorcycleRequest, nil
}

// Validate verifies that a GetMotorcycleRequest's fields contain valid data.
// Returns (an instance of GetMotorcycleRequest, nil) on success, otherwise (nil, error)
func (request GetMotorcycleRequest) Validate() error {
	return validation.ValidateStruct(&request,
		// ID is required and it must be greater than 0.
		validation.Field(&request.ID, validation.Required, validation.Min(1)))
}
