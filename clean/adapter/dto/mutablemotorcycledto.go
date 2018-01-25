// Package dto contains data transfer objects sent to/from client applications.
package dto

// MutableMotorcycleDto contains the data that can be modified for a motorcycle.
type MutableMotorcycleDto struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Vin   string `json:"vin"`
}
