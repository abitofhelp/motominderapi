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

// TestMotorcycleRepository_Insert verifies that an insert is successful.
func TestMotorcycleRepository_Insert(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle(1, "Honda", "Shadow", 2006)

	// ACT
	repo.Insert(motorcycle)

	// ASSERT
	assert.True(t, len(repo.Motorcycles) == 1)
}
