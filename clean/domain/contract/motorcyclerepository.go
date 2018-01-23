// Package contract contains contracts for entities and other objects.
package contract

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
)

// MotorcycleRepository defines the contract for its actions.
type MotorcycleRepository interface {
	FindByVin(vin string) (*entity.Motorcycle, error)
	ExistsByVin(vin string) (bool, error)

	List() ([]entity.Motorcycle, error)
	Insert(entity *entity.Motorcycle) (*entity.Motorcycle, error)
	Update(entity *entity.Motorcycle) (*entity.Motorcycle, error)
	Delete(id typedef.ID) error
	FindByID(id typedef.ID) (*entity.Motorcycle, error)
	Save() error
	Validate() error
	ExistsByID(id typedef.ID) (bool, error)
}
