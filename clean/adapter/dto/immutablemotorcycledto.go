// Package dto contains data transfer objects sent to/from client applications.
package dto

import (
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
	"time"
)

// ImmutableMotorcycleDto contains read-only motorcycle information.
type ImmutableMotorcycleDto struct {
	ID          typedef.ID `json:"id"`
	Make        string     `json:"make"`
	Model       string     `json:"model"`
	Year        int        `json:"year"`
	Vin         string     `json:"vin"`
	CreatedUtc  time.Time  `json:"createdUtc"`
	ModifiedUtc time.Time  `json:"modifiedUtc"`
}