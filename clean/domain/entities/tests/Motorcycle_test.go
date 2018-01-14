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
	_, err := entities.NewMotorcycle("Ford", "Falcon", 2006, "01234567890123456")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleMake_HondaIsValid verifies that Honda is a valid manufacturer.
func TestMotorcycleMake_HondaIsValid(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")

	// ASSERT
	assert.Nil(t, err)
}

// TestMotorcycleMake_NotEmpty verifies that the make cannot be empty.
func TestMotorcycleMake_NotEmpty(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("", "Falcon", 2006, "01234567890123456")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleMake_LengthLTE20 verifies that the make cannot exceed 20 characters.
func TestMotorcycleMake_LengthLTE20(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("0123456789012345678901234", "Falcon", 2006, "01234567890123456")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleModel_NotEmpty verifies that the make cannot be empty.
func TestMotorcycleModel_NotEmpty(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "", 2006, "01234567890123456")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleModel_LengthLTE20 verifies that the make cannot exceed 20 characters.
func TestMotorcycleModel_LengthLTE20(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "0123456789012345678901234", 2006, "01234567890123456")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleYear_LT1999 verifies that the year cannot be less than 1999.
func TestMotorcycleYear_LT1999(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 1998, "01234567890123456")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleYear_1999 verifies that 1999 is a valid year.
func TestMotorcycleYear_1999(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 1999, "01234567890123456")

	// ASSERT
	assert.Nil(t, err)
}

// TestMotorcycleYear_2020 verifies that 2020 is a valid year.
func TestMotorcycleYear_2020(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2020, "01234567890123456")

	// ASSERT
	assert.Nil(t, err)
}

// TestMotorcycleYear_GT2020 verifies that the year cannot be greater than 2020.
func TestMotorcycleYear_GT2020(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2021, "01234567890123456")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycle_ChangeFieldValueAndValidate_Successful verifies that changing a field with
// a correct value will pass validation.
func TestMotorcycle_ChangeFieldValueAndValidate_Successful(t *testing.T) {

	// ARRANGE
	motorcycle, err := entities.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")
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
	motorcycle, err := entities.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")
	motorcycle.Year = 3000

	// ACT
	err = motorcycle.Validate()

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleVin_GT17Characters verifies that the VIN's length cannot be more than 17 characters.
func TestMotorcycleVin_GT17Characters(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2021, "012345678901234567")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleVin_LT17Characters verifies that the VIN's length cannot be less than 17 characters.
func TestMotorcycleVin_LT17Characters(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2021, "0123456789012345")

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleVin_17Characters verifies that the VIN's length is 17 characters.
func TestMotorcycleVin_17Characters(t *testing.T) {

	// ARRANGE

	// ACT
	_, err := entities.NewMotorcycle("Honda", "Shadow", 2021, "01234567890123456")

	// ASSERT
	assert.NotNil(t, err)
}
