package interfaces

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entities"
)

type IMotorcycleRepository interface {
	Insert(motorcycle entities.Motorcycle) error
	Delete(motorcycle entities.Motorcycle) error
	FindById(id uint64) entities.Motorcycle
	Save(motorcycle entities.Motorcycle) error
}
