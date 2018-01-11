// Package repositories contains implementations of data repositories.
package repositories

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entities"

	"github.com/go-ozzo/ozzo-validation"
)

// MotorcycleRepository provides CRUD operations against a collection of motorcycles.
type MotorcycleRepository struct {
	Motorcycles []entities.Motorcycle `json:"motorcycles"`
}

// NewMotorcycleRepository creates a new instance of a MotorcycleRepository.
// Returns 'nil, error' when there is an error, otherwise a 'MotorcycleRepository, nil'.
func NewMotorcycleRepository() (*MotorcycleRepository, error) {
	motorcycleRepository := &MotorcycleRepository{}
	err := motorcycleRepository.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return motorcycleRepository, nil
}

// Validate tests that a motorcycle repository is valid.
// Returns nil on success, otherwise and error.
func (repo MotorcycleRepository) Validate() error {
	return validation.ValidateStruct(&repo)
}

// List gets the list of motorcycles in the repository.
// Returns the list of motorcycles, or an error.
func (repo MotorcycleRepository) List() ([]entities.Motorcycle, error) {
	return repo.Motorcycles, nil
}

// Insert adds a motorcycle to the repository.
// Returns nil on success, otherwise and error.
func (repo *MotorcycleRepository) Insert(motorcycle *entities.Motorcycle) error {
	repo.Motorcycles = append(repo.Motorcycles, *motorcycle)
	return nil
}

/*

func Delete(repo interfaces.IMotorcycleRepository, motorcycle entities.Motorcycle) error {
	return nil
}

func FindById(repo interfaces.IMotorcycleRepository, id uint64) entities.Motorcycle {
	moto := entities.Motorcycle{
		Id:    123,
		Make:  "KTM",
		Model: "350 EXC-F",
		Year:  2018,
	}
numbers = append(numbers, 1)
	return moto
}

func Save(repo interfaces.IMotorcycleRepository, motorcycle entities.Motorcycle) error {
	return nil
}
*/
