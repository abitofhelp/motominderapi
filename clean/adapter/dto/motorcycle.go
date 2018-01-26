// Package dto contains data transfer objects sent to/from client applications.
package dto

import (
	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

// MotorcycleDto contains motorcycle information.
type MotorcycleDto struct {
	ID          typedef.ID `json:"id"`
	Make        string     `json:"make"`
	Model       string     `json:"model"`
	Year        int        `json:"year"`
	Vin         string     `json:"vin"`
	CreatedUtc  time.Time  `json:"createdUtc"`
	ModifiedUtc time.Time  `json:"modifiedUtc"`
}

func NewMotorcycleDto(motorcycle entity.Motorcycle) (*MotorcycleDto, error) {
	moto := &MotorcycleDto{
		ID:          motorcycle.ID,
		Make:        motorcycle.Make,
		Model:       motorcycle.Model,
		Year:        motorcycle.Year,
		Vin:         motorcycle.Vin,
		CreatedUtc:  motorcycle.CreatedUtc,
		ModifiedUtc: motorcycle.ModifiedUtc,
	}
	err := motorcycle.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return moto, nil
}

// Validate implemented Entity.Validate().  It verifies that a motorcycle's fields contain valid data that satisfies enterprise's common business rules.
// Returns nil if the motorcycle contains valid data, otherwise an error.
func (m MotorcycleDto) Validate() error {
	return validation.ValidateStruct(&m,
		// Make cannot be nil, cannot be empty, max length of 20, and not Ford (case insensitive)
		validation.Field(&m.Make, validation.Required, validation.Length(constant.MinMakeLength, constant.MaxMakeLength), validation.By(entity.IsInvalidManufacturer)),
		// Model cannot be nil, cannot be empty, and max length of 20
		validation.Field(&m.Model, validation.Required, validation.Length(constant.MinModelLength, constant.MaxModelLength)),
		// Year must be between 1999 and 2020, inclusive.
		validation.Field(&m.Year, validation.Required, validation.Min(constant.MinYear), validation.Max(constant.MaxYear)),
		// Vin cannot be nil, cannot be empty, and has a length of 17
		validation.Field(&m.Vin, validation.Required, validation.By(entity.Is17Characters)),
	)
}
