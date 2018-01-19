// Package dto contains data transfer objects sent from client applications.
package dto

import "time"

// MotorcycleDto contains the motorcycle data that will be returned to a client.
type MotorcycleDto struct {
	ID          int       `json:"id"`
	Make        string    `json:"make"`
	Model       string    `json:"model"`
	Year        int       `json:"year"`
	Vin         string    `json:"vin"`
	CreatedUtc  time.Time `json:"createdUtc"`
	ModifiedUtc time.Time `json:"modifiedUtc"`
}
