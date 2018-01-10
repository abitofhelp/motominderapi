// Package motorcycle_tests implements unit tests for the Motorcycle entity.
package entities

import (
	"errors"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
)

// Motorcycle is an entity
type Motorcycle struct {
	Id    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	//createdUtc  time.Time `json:"createdUtc"`
	//deletedUtc  time.Time `json:"deletedUtc"`
	//modifiedUtc time.Time `json:"modifiedUtc"`
	//rowVersion  []byte    `json:"rowVersion"`
}

// VALIDATION OF BUSINESS RULES
// Assumes that after changing values of a motorcycle,
// Validate() will be invoked.

// IsValidManufacturer verifies that a motorcycle's make is not an invalid manufacturer.
// Returns nil if the manufacturer is valid, otherwise an error.
func IsInvalidManufacturer(value interface{}) error {
	s, _ := value.(string)

	// Test for invalid manufacturers
	if strings.EqualFold(s, "Ford") {
		return errors.New("cannot be Ford")
	}
	return nil
}

// Validate verifies that a motorcycle's fields contain valid data that satisfies business rules.
// Returns nil if the motorcycle contains valid data, otherwise an error.
func (m Motorcycle) Validate() error {
	return validation.ValidateStruct(&m,
		// Id must be non-zero and a positive value.
		validation.Field(&m.Id, validation.Required, validation.Min(1)),
		// Make cannot be nil, cannot be empty, max length of 20, and not Ford (case insensitive)
		validation.Field(&m.Make, validation.Required, validation.NotNil, validation.Length(1, 20), validation.By(IsInvalidManufacturer)),
		// Model cannot be nil, cannot be empty, and max length of 20
		validation.Field(&m.Model, validation.Required, validation.NotNil, validation.Length(1, 20)),
		// Year must be between 1999 and 2020, inclusive.
		validation.Field(&m.Year, validation.Required, validation.Min(1999), validation.Max(2020)),
	)
}

// Creates a new motorcycle
// Returns nil, error when there is an error, otherwise motorcycle, nil.
func NewMotorcycle(id int, make string, model string, year int) (*Motorcycle, error) {

	motorcycle := &Motorcycle{
		Id:    id,
		Make:  make,
		Model: model,
		Year:  year,
	}
	err := motorcycle.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return motorcycle, nil
}
