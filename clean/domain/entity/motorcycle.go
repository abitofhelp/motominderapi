// Package entity contains the domain entities.
package entity

import (
	"errors"
	"strings"

	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

// Motorcycle is an entity
type Motorcycle struct {
	ID          int       `json:"id"`
	Make        string    `json:"make"`
	Model       string    `json:"model"`
	Year        int       `json:"year"`
	Vin         string    `json:"vin"`
	CreatedUtc  time.Time `json:"createdUtc"`
	ModifiedUtc time.Time `json:"modifiedUtc"`
	//rowVersion  []byte    `json:"rowVersion"`
}

// IsInvalidManufacturer verifies that a motorcycle's make is not an invalid manufacturer.
// Returns nil if the manufacturer is valid, otherwise an error.
func IsInvalidManufacturer(value interface{}) error {
	s, _ := value.(string)

	// Test for invalid manufacturers
	if strings.EqualFold(s, "Ford") {
		return errors.New("cannot be Ford")
	}
	return nil
}

// Is17Characters verifies that a string has 17 characters.
// Returns nil if the string does not contain 17 characters, otherwise an error.
func Is17Characters(value interface{}) error {
	s, _ := value.(string)

	if len(s) != constant.VinLength {
		return errors.New("must contain 17 characters")
	}

	return nil
}

// Validate verifies that a motorcycle's fields contain valid data that satisfies enterprise's common business rules.
// Returns nil if the motorcycle contains valid data, otherwise an error.
func (m Motorcycle) Validate() error {
	return validation.ValidateStruct(&m,
		// Make cannot be nil, cannot be empty, max length of 20, and not Ford (case insensitive)
		validation.Field(&m.Make, validation.Required, validation.Length(constant.MinMakeLength, constant.MaxMakeLength), validation.By(IsInvalidManufacturer)),
		// Model cannot be nil, cannot be empty, and max length of 20
		validation.Field(&m.Model, validation.Required, validation.Length(constant.MinModelLength, constant.MaxModelLength)),
		// Year must be between 1999 and 2020, inclusive.
		validation.Field(&m.Year, validation.Required, validation.Min(constant.MinYear), validation.Max(constant.MaxYear)),
		// Vin cannot be nil, cannot be empty, and has a length of 17
		validation.Field(&m.Vin, validation.Required, validation.By(Is17Characters)),
	)
}

// NewMotorcycle creates a new instance of a Motorcycle.
// Returns (nil, error) when there is an error, otherwise (motorcycle, nil).
func NewMotorcycle(make string, model string, year int, vin string) (*Motorcycle, error) {

	motorcycle := &Motorcycle{
		ID:    0,
		Make:  make,
		Model: model,
		Year:  year,
		Vin:   vin,
		// CreatedUtc: Set when an instance is created in the repository.
		// ModifiedUtc: Set when an instance is updated in the repository.
	}

	err := motorcycle.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return motorcycle, nil
}
