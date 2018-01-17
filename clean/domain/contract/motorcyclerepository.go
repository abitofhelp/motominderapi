// Package contract contains contracts for entities and other objects.
package contract

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
)

// MotorcycleRepository defines the contract for its actions.
type MotorcycleRepository interface {
	List() ([]entity.Motorcycle, error)
	Insert(motorcycle *entity.Motorcycle) (*entity.Motorcycle, error)
	FindByVin(vin string) (*entity.Motorcycle, error)
	Update(motorcycle *entity.Motorcycle) (*entity.Motorcycle, error)
	Delete(motorcycle *entity.Motorcycle) error
	FindByID(id int) (*entity.Motorcycle, error)
	Save() error
	Validate() error
}
