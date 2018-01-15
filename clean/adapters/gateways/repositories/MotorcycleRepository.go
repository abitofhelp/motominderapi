// Package repositories contains implementations of data repositories.
package repositories

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entities"

	"errors"
	"github.com/abitofhelp/motominderapi/clean/domain/constants"
	"github.com/go-ozzo/ozzo-validation"
	"sort"
	"time"
)

// nextID is the next primary key ID value for an object being inserted into the repository.
var nextID = 0

// MotorcycleRepository provides CRUD operations against a collection of motorcycles.
type MotorcycleRepository struct {
	// These items are unordered.
	Motorcycles []entities.Motorcycle `json:"motorcycles"`
}

// NewMotorcycleRepository creates a new instance of a MotorcycleRepository.
// Returns (nil, error) when there is an error, otherwise a (MotorcycleRepository, nil).
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
// Returns nil on success, otherwise an error.
func (repo MotorcycleRepository) Validate() error {
	return validation.ValidateStruct(&repo)
}

// List gets the unordered list of motorcycles in the repository.
// Returns the (list of motorcycles, nil) or an (nil, error).
func (repo MotorcycleRepository) List() ([]entities.Motorcycle, error) {
	return repo.Motorcycles, nil
}

// Insert adds a motorcycle to the repository.
// Do not permit duplicate ID values.
// Returns the (new motorcycle, nil) on success, otherwise (nil, error).
func (repo *MotorcycleRepository) Insert(motorcycle *entities.Motorcycle) (*entities.Motorcycle, error) {

	// Determine whether the motorcycle already exists in the repository.
	i, err := repo.findByID(motorcycle.ID)
	if i != constants.InvalidEntityID {
		return nil, errors.New("cannot insert this motorcycle because the ID already exists")
	}

	// Save the time when this entity was created in the repository.
	motorcycle.ID = repo.getNextID()
	motorcycle.CreatedUtc = time.Now().UTC()

	// Validate the object
	err = motorcycle.Validate()
	if err != nil {
		return nil, err
	}

	repo.Motorcycles = append(repo.Motorcycles, *motorcycle)

	return motorcycle, nil
}

// Update replaces a motorcycle an existing motorcycle in the repository.
// If the motorcycle does not exist, an error is returned.
// Returns (updated motorcycle, nil) on success, otherwise an (nil, error).
func (repo *MotorcycleRepository) Update(motorcycle *entities.Motorcycle) (*entities.Motorcycle, error) {
	// Find the motorcycle, so it can be updated in the repository.
	i, _ := repo.findByID(motorcycle.ID)
	if i == constants.InvalidEntityID {
		return nil, errors.New("cannot update a motorcycle that does not exist")
	}

	// Save the time when this entity was updated in the repository.
	motorcycle.ModifiedUtc = time.Now().UTC()

	// Validate the object
	err := motorcycle.Validate()
	if err != nil {
		return nil, err
	}

	repo.Motorcycles[i] = *motorcycle

	return motorcycle, nil

}

// findByID a motorcycle in the repository using its primary key, ID.
// Returns its (index, nil) on success, otherwise (index of -1, error).
func (repo *MotorcycleRepository) findByID(id int) (int, error) {
	// Sort the list of motorcycles by id and find the index to the motorcycle.
	// The result is the slice index for the single element or -1.
	i := sort.Search(len(repo.Motorcycles), func(i int) bool {
		return repo.Motorcycles[i].ID >= id
	})

	if i < len(repo.Motorcycles) && repo.Motorcycles[i].ID == id {
		// Found the motorcycle
		return i, nil
	}

	// Motorcycle was not found.
	return -1, errors.New("motorcycle was not found using its ID")
}

//FindByID a motorcycle in the repository using its primary key, ID.
// Returns (motorcycle, nil) on success, otherwise (nil, error).
func (repo *MotorcycleRepository) FindByID(id int) (*entities.Motorcycle, error) {

	// Try to find the index for the motorcycle in the repository.
	i, err := repo.findByID(id)

	if err != nil {
		return nil, err
	}

	// Motorcycle was found.
	return &repo.Motorcycles[i], nil
}

// FindByVin a motorcycle in the repository using its VIN.
// Returns (motorcycle, nil) on success, otherwise (nil, error).
func (repo *MotorcycleRepository) FindByVin(vin string) (*entities.Motorcycle, error) {
	// Sort the list of motorcycles by VIN and search.
	i := sort.Search(len(repo.Motorcycles), func(i int) bool {
		return repo.Motorcycles[i].Vin >= vin
	})

	if i < len(repo.Motorcycles) && repo.Motorcycles[i].Vin == vin {
		// Found the motorcycle
		return &repo.Motorcycles[i], nil
	}

	// Motorcycle was not found.
	return nil, errors.New("motorcycle was not found using its VIN")
}

// Delete an existing motorcycle from the repository.
// If the motorcycle does not exist, an error is returned.
// Returns nil on success, otherwise an error.
func (repo *MotorcycleRepository) Delete(motorcycle *entities.Motorcycle) error {
	// Find the motorcycle, so it can be updated in the repository.
	i, _ := repo.findByID(motorcycle.ID)
	if i == constants.InvalidEntityID {
		return errors.New("cannot delete a motorcycle that does not exist")
	}

	repo.Motorcycles = repo.removeAtIndex(i)

	return nil
}

// removeAtIndex deletes the motorcycle at the specified index.
// This is an internal method.
// Returns the updated list of motorcycles in the repository.
func (repo *MotorcycleRepository) removeAtIndex(index int) []entities.Motorcycle {
	return append(repo.Motorcycles[:index], repo.Motorcycles[index+1:]...)
}

// Save all of the changes to the repository (assuming some kind of unit of work/dbContext).
// Returns nil on success, otherwise an error.
func (repo *MotorcycleRepository) Save() error {
	return nil
}

// GetNextID determines the next primary key ID value when an item is inserted into the repository.
// Returns the next ID.
func (repo *MotorcycleRepository) getNextID() int {
	nextID = nextID + 1
	return nextID
}
