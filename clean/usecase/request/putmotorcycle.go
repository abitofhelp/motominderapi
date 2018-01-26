// Package request contains the request messages for the use cases.
package request

import (
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
	"github.com/go-ozzo/ozzo-validation"

	"github.com/abitofhelp/motominderapi/clean/domain/entity"
)

// UpdateMotorcycleRequest is a simple dto containing the required data for the UpdateMotorcycleInteractor.
type UpdateMotorcycleRequest struct {
	ID         typedef.ID         `json:"id"`
	Motorcycle *entity.Motorcycle `json:"motorcycle"`
}

// NewUpdateMotorcycleRequest creates a new instance of a UpdateMotorcycleRequest.
// Returns (nil, error) when there is an error, otherwise (UpdateMotorcycleRequest, nil).
func NewUpdateMotorcycleRequest(id typedef.ID, motorcycle *entity.Motorcycle) (*UpdateMotorcycleRequest, error) {

	motorcycleRequest := &UpdateMotorcycleRequest{
		ID:         id,
		Motorcycle: motorcycle,
	}

	err := motorcycleRequest.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return motorcycleRequest, nil
}

// Validate verifies that a UpdateMotorcycleRequest's fields contain valid data.
// Returns (an instance of UpdateMotorcycleRequest, nil) on success, otherwise (nil, error)
func (request UpdateMotorcycleRequest) Validate() error {
	return validation.ValidateStruct(&request,
		// ID is required and it must be greater than 0.
		validation.Field(&request.ID, validation.Required, validation.Min(1)),
		// Make cannot be nil, cannot be empty, max length of 20, and not Ford (case insensitive)
		validation.Field(&request.Motorcycle, validation.Required))

}
