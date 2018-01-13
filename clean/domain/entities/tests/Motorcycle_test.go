// Package motorcycleTests implements unit tests for the Motorcycle entity.
package motorcycleTests

import (
	"testing"

	"github.com/abitofhelp/motominderapi/clean/domain/entities"
	"github.com/stretchr/testify/assert"
)

// TestMotorcycleMake_FordIsInvalid verifies that Ford is an invalid manufacturer.
func TestMotorcycleMake_FordIsInvalid(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Ford", "Falcon", 2006)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleMake_HondaIsValid verifies that Honda is a valid manufacturer.
func TestMotorcycleMake_HondaIsValid(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2006)

	// ASSERT
	assert.Nil(t, err)
}

// TestMotorcycleMake_NotEmpty verifies that the make cannot be empty.
func TestMotorcycleMake_NotEmpty(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("", "Falcon", 2006)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleMake_LengthLTE20 verifies that the make cannot exceed 20 characters.
func TestMotorcycleMake_LengthLTE20(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("0123456789012345678901234", "Falcon", 2006)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleModel_NotEmpty verifies that the make cannot be empty.
func TestMotorcycleModel_NotEmpty(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "", 2006)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleModel_LengthLTE20 verifies that the make cannot exceed 20 characters.
func TestMotorcycleModel_LengthLTE20(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "0123456789012345678901234", 2006)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleYear_LT1999 verifies that the year cannot be less than 1999.
func TestMotorcycleYear_LT1999(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 1998)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleYear_1999 verifies that 1999 is a valid year.
func TestMotorcycleYear_1999(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 1999)

	// ASSERT
	assert.Nil(t, err)
}

// TestMotorcycleYear_2020 verifies that 2020 is a valid year.
func TestMotorcycleYear_2020(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2020)

	// ASSERT
	assert.Nil(t, err)
}

// TestMotorcycleYear_GT2020 verifies that the year cannot be greater than 2020.
func TestMotorcycleYear_GT2020(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2021)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycle_ChangeFieldValueAndValidate_Successful verifies that changing a field with
// a correct value will pass validation.
func TestMotorcycle_ChangeFieldValueAndValidate_Successful(t *testing.T) {

	// ARRANGE
	motorcycle, err := entities.NewMotorcycle("Honda", "Shadow", 2006)
	motorcycle.Year = 2007

	// ACT
	err = motorcycle.Validate()

	// ASSERT
	assert.Nil(t, err)
}

// TestMotorcycle_ChangeFieldValueAndValidate_Failure verifies that changing a field with
// an invalid value will fail validation.
func TestMotorcycle_ChangeFieldValueAndValidate_Failure(t *testing.T) {

	// ARRANGE
	motorcycle, err := entities.NewMotorcycle("Honda", "Shadow", 2006)
	motorcycle.Year = 3000

	// ACT
	err = motorcycle.Validate()

	// ASSERT
	assert.NotNil(t, err)
}
