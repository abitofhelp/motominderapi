// Package repository contains implementations of data repositories.
package repository

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entity"

	"fmt"
	"github.com/pkg/errors"
	"sort"
	"time"

	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/operationstatus"
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
	"github.com/go-ozzo/ozzo-validation"
)

// MotorcycleRepository provides CRUD operations against a collection of motorcycles.
type MotorcycleRepository struct {
	// NextID is the next primary key ID value for an object being inserted into the repository.
	NextID typedef.ID `json:"nextId"`

	// These items are unordered.
	Motorcycles []entity.Motorcycle `json:"motorcycles"`
}

// NewMotorcycleRepository creates a new instance of a MotorcycleRepository.
// Returns (nil, error) when there is an error, otherwise a (MotorcycleRepository, nil).
func NewMotorcycleRepository() (*MotorcycleRepository, error) {
	motorcycleRepository := &MotorcycleRepository{

		// nextID is the next primary key ID value for an object being inserted into the repository.
		NextID: 0,

		// Ensure that we create an empty slice rather than the default for []entity.Motorcycle, which is a null pointer.
		Motorcycles: make([]entity.Motorcycle, 0),
	}
	err := motorcycleRepository.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return motorcycleRepository, nil
}

// Validate test that a motorcycle repository is valid.
// Returns nil on success, otherwise an error.
func (repo MotorcycleRepository) Validate() error {
	return validation.ValidateStruct(&repo,
		// Motorcycles can be empty, but not nil
		validation.Field(&repo.Motorcycles, validation.NotNil))
}

// List gets the unordered list of motorcycles in the repository.
// Returns the (list of motorcycles, Ok, nil), otherwise a (nil, operationStatus, error).
func (repo *MotorcycleRepository) List() ([]entity.Motorcycle, operationstatus.OperationStatus, error) {
	if repo.Motorcycles == nil {
		return nil, operationstatus.InternalError, errors.New("list of motorcycles is nil, so create an instance of []entity.Motorcycle")
	}
	return repo.Motorcycles, operationstatus.Ok, nil
}

// ExistsByVin determines whether a motorcycle with the VIN exists in the repository.
// Returns (true, Ok, nil) for found, (false, Ok, nil) for not found, otherwise (false, operationStatus, error).
func (repo *MotorcycleRepository) ExistsByVin(vin string) (bool, operationstatus.OperationStatus, error) {
	// Determine whether a motorcycle with the VIN already exists in the repository.
	moto, status, err := repo.FindByVin(vin)
	if err != nil {
		return false, status, err
	}

	// Found
	if moto != nil {
		return true, status, nil
	}

	// Not Found
	return false, status, nil
}

// ExistsByID determines whether a motorcycle with the ID exists in the repository.
// Returns (true, Ok, nil) for found, (false, Ok, nil) for not found, otherwise (false, operationStatus, error).
func (repo *MotorcycleRepository) ExistsByID(id typedef.ID) (bool, operationstatus.OperationStatus, error) {
	// Determine whether a motorcycle with the VIN already exists in the repository.
	moto, status, err := repo.FindByID(id)
	if err != nil {
		return false, status, err
	}

	// Found
	if moto != nil {
		return true, status, nil
	}

	// Not Found
	return false, status, nil
}

// Insert adds a motorcycle to the repository.
// Does not permit duplicate VIN values.
// Returns the (new motorcycle, Ok, nil) on success, otherwise an (nil, operationStatus, error).
func (repo *MotorcycleRepository) Insert(motorcycle *entity.Motorcycle) (*entity.Motorcycle, operationstatus.OperationStatus, error) {
	exists, status, err := repo.ExistsByVin(motorcycle.Vin)
	if err != nil {
		return nil, status, err
	}

	if exists {
		return nil, status, fmt.Errorf("cannot insert the motorcycle with VIN %s because the VIN already exists in the repository", motorcycle.Vin)
	}

	// Assign the ID to the new motorcycle, and save the time when this entity was created in the repository.
	motorcycle.ID = repo.getNextID()
	motorcycle.CreatedUtc = time.Now().UTC()

	// Validate the object
	err = motorcycle.Validate()
	if err != nil {
		return nil, operationstatus.InternalError, err
	}

	repo.Motorcycles = append(repo.Motorcycles, *motorcycle)

	return motorcycle, operationstatus.Ok, nil
}

// Update replaces an existing motorcycle in the repository.
// If the motorcycle does not exist, an error is returned.
// Does not permit duplicate VIN values.
// Returns (updated motorcycle, Ok, nil) on success, otherwise an (nil, operationStatus, error).
func (repo *MotorcycleRepository) Update(id typedef.ID, motorcycle *entity.Motorcycle) (*entity.Motorcycle, operationstatus.OperationStatus, error) {
	moto, status, err := repo.FindByID(id)
	if err != nil {
		return nil, status, err
	}

	if moto == nil {
		return nil, status, fmt.Errorf("cannot update the motorcycle with ID %d because it doesn't exist in the repository", id)
	}

	// Update all fields...
	moto.ID = motorcycle.ID
	moto.Make = motorcycle.Make
	moto.Model = motorcycle.Model
	moto.Year = motorcycle.Year
	moto.Vin = motorcycle.Vin
	moto.CreatedUtc = motorcycle.CreatedUtc

	// Save the time when this entity was updated in the repository.
	moto.ModifiedUtc = time.Now().UTC()

	// Validate the object
	err = moto.Validate()
	if err != nil {
		return nil, operationstatus.InternalError, err
	}

	// We are updating the actual object in the slice, so the slice doesn't need to be manipulated.

	return moto, operationstatus.Ok, nil
}

// findByID a motorcycle in the repository using its primary key, ID.
// Returns its (index, nil) on found, (index of -1, nil) for not found, otherwise (index of -1, error).
func (repo *MotorcycleRepository) findByID(id typedef.ID) (int, error) {
	if repo.Motorcycles == nil {
		return constant.InvalidEntityID, errors.New("list of motorcycles is nil")
	}

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
	return constant.InvalidEntityID, nil
}

// FindByID a motorcycle in the repository using its primary key, ID.
// Returns (motorcycle, nil) on found, (nil, nil) for not found,, otherwise (nil, error).
func (repo *MotorcycleRepository) FindByID(id typedef.ID) (*entity.Motorcycle, operationstatus.OperationStatus, error) {

	// Try to find the index for the motorcycle in the repository.
	i, err := repo.findByID(id)

	if err != nil {
		return nil, operationstatus.InternalError, err
	}

	// Not Found
	if i == constant.InvalidEntityID {
		return nil, operationstatus.NotFound, nil
	}

	// Motorcycle was found.
	return &repo.Motorcycles[i], operationstatus.Ok, nil
}

// findByVin a motorcycle in the repository using its VIN.
// Returns its (index, nil) on found, (index of -1, nil) for not found, otherwise (index of -1, error).
func (repo *MotorcycleRepository) findByVin(vin string) (int, error) {
	if repo.Motorcycles == nil {
		return constant.InvalidEntityID, errors.New("list of motorcycles is nil, so create an instance of []entity.Motorcycle")
	}

	// Sort the list of motorcycles by id and find the index to the motorcycle.
	// The result is the slice index for the single element or -1.
	i := sort.Search(len(repo.Motorcycles), func(i int) bool {
		return repo.Motorcycles[i].Vin >= vin
	})

	if i < len(repo.Motorcycles) && repo.Motorcycles[i].Vin == vin {
		// Found the motorcycle
		return i, nil
	}

	// Motorcycle was not found.
	return constant.InvalidEntityID, nil
}

// FindByVin a motorcycle in the repository using its VIN.
// Returns (motorcycle, Ok, nil) on success, otherwise an (nil, operationStatus, error).
func (repo *MotorcycleRepository) FindByVin(vin string) (*entity.Motorcycle, operationstatus.OperationStatus, error) {
	// Try to find the index for the motorcycle in the repository.
	i, err := repo.findByVin(vin)

	if err != nil {
		return nil, operationstatus.InternalError, err
	}

	if i == constant.InvalidEntityID {
		return nil, operationstatus.NotFound, nil
	}

	// Motorcycle was found.
	return &repo.Motorcycles[i], operationstatus.Found, nil
}

// Delete an existing motorcycle from the repository.
// If the motorcycle does not exist, an error is returned.
// Returns (Ok, nil) on success, otherwise an (operationStatus, error).
func (repo *MotorcycleRepository) Delete(id typedef.ID) (operationstatus.OperationStatus, error) {

	i, err := repo.findByID(id)
	if err != nil {
		return operationstatus.InternalError, err
	}

	if i == constant.InvalidEntityID {
		return operationstatus.NotFound, fmt.Errorf("cannot delete the motorcycle with ID %d because it was not found", id)
	}

	repo.Motorcycles = repo.removeAtIndex(i)

	return operationstatus.Ok, nil
}

// removeAtIndex deletes the motorcycle at the specified index.
// This is an internal method.
// Returns the updated list of motorcycles in the repository.
func (repo *MotorcycleRepository) removeAtIndex(index int) []entity.Motorcycle {
	return append(repo.Motorcycles[:index], repo.Motorcycles[index+1:]...)
}

// Save all of the changes to the repository (assuming some kind of unit of work/dbContext).
// Returns nil on success, otherwise an error.
func (repo *MotorcycleRepository) Save() (operationstatus.OperationStatus, error) {
	return operationstatus.Ok, nil
}

// GetNextID determines the next primary key ID value when an item is inserted into the repository.
// Returns the next ID.
func (repo *MotorcycleRepository) getNextID() typedef.ID {
	repo.NextID = repo.NextID + 1
	return repo.NextID
}
