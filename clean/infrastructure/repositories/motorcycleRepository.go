package repositories

import (
	"github.com/abitofhelp/motominderapi/clean/domain/entities"
	"github.com/abitofhelp/motominderapi/clean/domain/interfaces"
)

func Insert(repo interfaces.IMotorcycleRepository, motorcycle entities.Motorcycle) error {
	return nil
}

func Delete(repo interfaces.IMotorcycleRepository, motorcycle entities.Motorcycle) error {
	return nil
}

func FindById(repo interfaces.IMotorcycleRepository, id uint64) entities.Motorcycle {
	moto := entities.Motorcycle{
		Id:    123,
		Make:  "KTM",
		Model: "350 EXC-F",
		Year:  "2018",
	}

	return moto
}

func Save(repo interfaces.IMotorcycleRepository, motorcycle entities.Motorcycle) error {
	return nil
}
