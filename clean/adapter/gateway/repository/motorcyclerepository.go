// Package repository contains implementations of data repositories.
package repository

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/operationstatus"

	"fmt"
	"sort"
	"time"

	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/go-ozzo/ozzo-validation"
	errors "github.com/pjebs/jsonerror"
)

// MotorcycleRepository provides CRUD operations against a collection of motorcycles.
type MotorcycleRepository struct {
	// NextID is the next primary key ID value for an object being inserted into the repository.
	NextID int `json:"nextId"`

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
	err := validation.ValidateStruct(&repo,
		// Motorcycles can be empty, but not nil
		validation.Field(&repo.Motorcycles, validation.NotNil))

	if err != nil {
		return errors.New(operationstatus.StatusInternalServerError, operationstatus.StatusText(operationstatus.StatusInternalServerError), err.Error())
	}

	return nil
}

// List gets the unordered list of motorcycles in the repository.
// Returns the (list of motorcycles, nil) or an (nil, error).
func (repo *MotorcycleRepository) List() ([]entity.Motorcycle, error) {
	if repo.Motorcycles == nil {
		return nil, errors.New(operationstatus.StatusInternalServerError, "list of motorcycles is nil", "create an instance of []entity.Motorcycle")
	}
	return repo.Motorcycles, nil
}

// Insert adds a motorcycle to the repository.
// Do not permit duplicate ID or VIN values.
// Returns the (new motorcycle, nil) on success, otherwise (nil, error).
func (repo *MotorcycleRepository) Insert(motorcycle *entity.Motorcycle) (*entity.Motorcycle, error) {
	// Determine whether the motorcycle already exists in the repository.
	i, _ := repo.findByID(motorcycle.ID)
	if i != constant.InvalidEntityID {
		return nil, errors.New(operationstatus.StatusFound, operationstatus.StatusText(operationstatus.StatusFound), fmt.Sprintf("cannot insert the motorcycle with ID %d because it already exists", motorcycle.ID))
	}

	i, _ = repo.findByVin(motorcycle.Vin)
	if i != constant.InvalidEntityID {
		return nil, errors.New(operationstatus.StatusFound, operationstatus.StatusText(operationstatus.StatusFound), fmt.Sprintf("cannot insert the motorcycle with VIN %s because it already exists", motorcycle.ID))
	}

	// Save the time when this entity was created in the repository.
	motorcycle.ID = repo.getNextID()
	motorcycle.CreatedUtc = time.Now().UTC()

	// Validate the object
	err := motorcycle.Validate()
	if err != nil {
		return nil, err
	}

	repo.Motorcycles = append(repo.Motorcycles, *motorcycle)

	return motorcycle, nil
}

// Update replaces a motorcycle an existing motorcycle in the repository.
// If the motorcycle does not exist, an error is returned.
// Returns (updated motorcycle, nil) on success, otherwise an (nil, error).
func (repo *MotorcycleRepository) Update(motorcycle *entity.Motorcycle) (*entity.Motorcycle, error) {
	// Find the motorcycle, so it can be updated in the repository.
	i, _ := repo.findByID(motorcycle.ID)
	if i == constant.InvalidEntityID {
		return nil, errors.New(operationstatus.StatusFound, operationstatus.StatusText(operationstatus.StatusFound), fmt.Sprintf("cannot update the motorcycle with ID %d because it does not exist", motorcycle.ID))
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
	if repo.Motorcycles == nil {
		return constant.InvalidEntityID, errors.New(operationstatus.StatusInternalServerError, "list of motorcycles is nil", "create an instance of []entity.Motorcycle")
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

//FindByID a motorcycle in the repository using its primary key, ID.
// Returns (motorcycle, nil) on success, otherwise (nil, error).
func (repo *MotorcycleRepository) FindByID(id int) (*entity.Motorcycle, error) {

	// Try to find the index for the motorcycle in the repository.
	i, err := repo.findByID(id)

	if err != nil {
		return nil, err
	}

	if i == constant.InvalidEntityID {
		return nil, nil
	}

	// Motorcycle was found.
	return &repo.Motorcycles[i], nil
}

// findByVin a motorcycle in the repository using its VIN.
// Returns its (index, nil) on success, otherwise (index of -1, error).
func (repo *MotorcycleRepository) findByVin(vin string) (int, error) {
	if repo.Motorcycles == nil {
		return constant.InvalidEntityID, errors.New(operationstatus.StatusInternalServerError, "list of motorcycles is nil", "create an instance of []entity.Motorcycle")
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
// Returns (motorcycle, nil) on success, otherwise (nil, error).
func (repo *MotorcycleRepository) FindByVin(vin string) (*entity.Motorcycle, error) {
	// Try to find the index for the motorcycle in the repository.
	i, err := repo.findByVin(vin)

	if err != nil {
		return nil, err
	}

	if i == constant.InvalidEntityID {
		return nil, nil
	}

	// Motorcycle was found.
	return &repo.Motorcycles[i], nil
}

// Delete an existing motorcycle from the repository.
// If the motorcycle does not exist, an error is returned.
// Returns nil on success, otherwise an error.
func (repo *MotorcycleRepository) Delete(id int) error {

	// Find the motorcycle, so it can be updated in the repository.
	i, _ := repo.findByID(id)
	if i == constant.InvalidEntityID {
		return errors.New(operationstatus.StatusNotFound, operationstatus.StatusText(operationstatus.StatusNotFound), fmt.Sprintf("cannot delete the motorcycle with ID %d because it was not found", id))

	}

	repo.Motorcycles = repo.removeAtIndex(i)

	return nil
}

// removeAtIndex deletes the motorcycle at the specified index.
// This is an internal method.
// Returns the updated list of motorcycles in the repository.
func (repo *MotorcycleRepository) removeAtIndex(index int) []entity.Motorcycle {
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
	repo.NextID = repo.NextID + 1
	return repo.NextID
}
