// Package request contains the request messages for the use cases.
package request

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/go-ozzo/ozzo-validation"
	errors "github.com/pjebs/jsonerror"
)

// InsertMotorcycleRequest is a simple dto containing the required data for the InsertMotorcycleInteractor.
type InsertMotorcycleRequest struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Vin   string `json:"vin"`
}

// NewInsertMotorcycleRequest creates a new instance of a InsertMotorcycleRequest.
// Returns (nil, error) when there is an error, otherwise (InsertMotorcycleRequest, nil).
func NewInsertMotorcycleRequest(make string, model string, year int, vin string) (*InsertMotorcycleRequest, error) {

	motorcycleRequest := &InsertMotorcycleRequest{
		Make:  make,
		Model: model,
		Year:  year,
		Vin:   vin,
	}

	err := motorcycleRequest.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return motorcycleRequest, nil
}

// Validate verifies that a InsertMotorcycleRequest's fields contain valid data.
// Returns (an instance of InsertMotorcycleRequest, nil) on success, otherwise (nil, error)
func (request InsertMotorcycleRequest) Validate() error {
	err := validation.ValidateStruct(&request,
		// Make cannot be nil, cannot be empty, max length of 20, and not Ford (case insensitive)
		validation.Field(&request.Make, validation.Required, validation.Length(1, 20)),
		// Model cannot be nil, cannot be empty, and max length of 20
		validation.Field(&request.Model, validation.Required, validation.Length(1, 20)),
		// Year must be between 1999 and 2020, inclusive.
		validation.Field(&request.Year, validation.Required, validation.Min(1999), validation.Max(2020)),
		// Vin cannot be nil, cannot be empty, and has a length of 17
		validation.Field(&request.Vin, validation.Required, validation.By(entity.Is17Characters)),
	)

	if err != nil {
		return errors.New(enumeration.StatusInternalServerError, enumeration.StatusText(enumeration.StatusInternalServerError), err.Error())
	}

	return nil
}
