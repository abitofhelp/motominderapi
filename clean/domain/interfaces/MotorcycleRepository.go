// Package interfaces contains contracts for entities.
package interfaces

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entities"
)

// MotorcycleRepository defines the contract for its actions.
type MotorcycleRepository interface {
	List() ([]entities.Motorcycle, error)
	Insert(motorcycle *entities.Motorcycle) error
	Find(motorcycle *entities.Motorcycle) (*entities.Motorcycle, error)
	Update(motorcycle *entities.Motorcycle) error
	Delete(motorcycle entities.Motorcycle) error
	FindByID(id int) (*entities.Motorcycle, error)
	Save() error
}
