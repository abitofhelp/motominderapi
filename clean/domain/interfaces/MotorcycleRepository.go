// Package interfaces contains contracts for entities.
package interfaces

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entities"
)

// MotorcycleRepository defines the contract for its actions.
type MotorcycleRepository interface {
	List() ([]entities.Motorcycle, error)
	Insert(motorcycle *entities.Motorcycle) error
	//Delete(motorcycle entities.Motorcycle) error
	//FindById(id uint64) entities.Motorcycle
	//Save(motorcycle entities.Motorcycle) error
}
