// Package requestmessages contains the request messages for the use cases.
package requestmessages

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entities"
	"github.com/go-ozzo/ozzo-validation"
)

// InsertMotorcycleRequestMessage is a simple dto containing the required data for the InsertMotorcycleInteractor.
type InsertMotorcycleRequestMessage struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Vin   string `json:"vin"`
}

// NewInsertMotorcycleRequestMessage creates a new instance of a InsertMotorcycleRequestMessage.
// Returns (nil, error) when there is an error, otherwise (InsertMotorcycleRequestMessage, nil).
func NewInsertMotorcycleRequestMessage(make string, model string, year int, vin string) (*InsertMotorcycleRequestMessage, error) {

	requestMessage := &InsertMotorcycleRequestMessage{
		Make:  make,
		Model: model,
		Year:  year,
		Vin:   vin,
	}

	err := requestMessage.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return requestMessage, nil
}

// Validate verifies that a InsertMotorcycleRequestMessage's fields contain valid data.
// Although it is possible that the same rules apply as for the Motorcycle entity, we
// will assume that different rules may be used with this dto.
// Returns nil if the InsertMotorcycleRequestMessage contains valid data, otherwise an error.
func (insertMotorcycleRequestMessage InsertMotorcycleRequestMessage) Validate() error {
	return validation.ValidateStruct(&insertMotorcycleRequestMessage,
		// Make cannot be nil, cannot be empty, max length of 20, and not Ford (case insensitive)
		validation.Field(&insertMotorcycleRequestMessage.Make, validation.Required, validation.Length(1, 20)),
		// Model cannot be nil, cannot be empty, and max length of 20
		validation.Field(&insertMotorcycleRequestMessage.Model, validation.Required, validation.Length(1, 20)),
		// Year must be between 1999 and 2020, inclusive.
		validation.Field(&insertMotorcycleRequestMessage.Year, validation.Required, validation.Min(1999), validation.Max(2020)),
		// Vin cannot be nil, cannot be empty, and has a length of 17
		validation.Field(&insertMotorcycleRequestMessage.Vin, validation.Required, validation.By(entities.Is17Characters)),
	)
}
