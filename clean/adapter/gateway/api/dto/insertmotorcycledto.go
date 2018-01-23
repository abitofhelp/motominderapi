// Package dto contains data transfer objects sent to/from client applications.
package dto

// InsertMotorcycleDto contains the data for creating a new motorcycle in the repository and is sent from the client.
type InsertMotorcycleDto struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Vin   string `json:"vin"`
}
