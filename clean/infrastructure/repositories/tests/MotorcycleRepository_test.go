// Package motorcycleRepositoryTests implements unit tests for the MotorcycleRepository.
package motorcycleRepositoryTests

import (
	"testing"

	"github.com/abitofhelp/motominderapi/clean/domain/entities"
	"github.com/abitofhelp/motominderapi/clean/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

// TestMotorcycleRepository_List verifies that an empty list of motorcycles is returned.
func TestMotorcycleRepository_ListEmpty(t *testing.T) {

	// ARRANGE

	// ACT
	repo, _ := repositories.NewMotorcycleRepository()

	// ASSERT
	assert.True(t, len(repo.Motorcycles) == 0)
}
// TestMotorcycleRepository_ListOfOne verifies that a list with one motorcycle is returned.
func TestMotorcycleRepository_ListOfOne(t *testing.T) {

	// ARRANGE

	// ACT
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle(1, "Honda", "Shadow", 2006)
	repo.Insert(motorcycle)

	// ASSERT
	assert.True(t, len(repo.Motorcycles) == 1)
}

// TestMotorcycleRepository_Insert verifies that an insert is successful.
func TestMotorcycleRepository_Insert(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle(1, "Honda", "Shadow", 2006)

	// ACT
	moto, _ := repo.Insert(motorcycle)

	// ASSERT
	assert.True(t, len(repo.Motorcycles) == 1)
	assert.True(t, *moto == repo.Motorcycles[0])
}

// TestMotorcycleRepository_FindByID verifies that an insert is successful.
func TestMotorcycleRepository_FindByID(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle(1, "Honda", "Shadow", 2006)
	moto, _:= repo.Insert(motorcycle)

	// ACT
	_, err := repo.FindByID(moto.ID)
	println(err)

	// ASSERT
	//assert.NotNil(t, foundMoto)
	//assert.True(t, moto.ID == foundMoto.ID)
}


// TestMotorcycleRepository_Update verifies that an update is successful.
func TestMotorcycleRepository_Update(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle(1, "Honda", "Shadow", 2006)
	moto, _ := repo.Insert(motorcycle)
	moto.Make = "Harley Davidson"

	// ACT
	repo.Update(moto)

	// ASSERT
	assert.True(t, repo.Motorcycles[0].Make == "Harley Davidson")
}