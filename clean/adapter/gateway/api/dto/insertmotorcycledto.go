// Package dto contains data transfer objects sent from client applications.
package dto

// MotorcycleDto contains the data for creating a new motorcycle in the repository.
type InsertMotorcycleDto struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Vin   string `json:"vin"`
}
