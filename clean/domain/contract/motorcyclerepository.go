// Package contract contains contracts for entities and other objects.
package contract

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/operationstatus"
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
)

// MotorcycleRepository defines the contract for its actions.
type MotorcycleRepository interface {
	FindByVin(vin string) (*entity.Motorcycle, operationstatus.OperationStatus, error)
	ExistsByVin(vin string) (bool, operationstatus.OperationStatus, error)
	ExistsByID(id typedef.ID) (bool, operationstatus.OperationStatus, error)

	List() ([]entity.Motorcycle, operationstatus.OperationStatus, error)
	Insert(motorcycle *entity.Motorcycle) (*entity.Motorcycle, operationstatus.OperationStatus, error)
	Update(id typedef.ID, motorcycle *entity.Motorcycle) (*entity.Motorcycle, operationstatus.OperationStatus, error)
	Delete(id typedef.ID) (operationstatus.OperationStatus, error)
	FindByID(id typedef.ID) (*entity.Motorcycle, operationstatus.OperationStatus, error)
	Save() (operationstatus.OperationStatus, error)
	Validate() error
}
