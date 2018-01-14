// Package responsemodels contains the response models for each use case/interactor.
package responsemodels

import "github.com/go-ozzo/ozzo-validation"

// InsertMotorcycleResponseMessage is a simple dto containing the response data for the InsertMotorcycleInteractor.
type InsertMotorcycleResponseMessage struct {
	ID    int   `json:"id"`
	Error error `json:"error"`
}

// NewInsertMotorcycleResponseMessage creates a new instance of a MotorcycleResponseMessage.
// Returns (nil, error) when there is an error, otherwise (InsertMotorcycleResponseMessage, nil).
func NewInsertMotorcycleResponseMessage(id int, err error) (*InsertMotorcycleResponseMessage, error) {

	responseMessage := &InsertMotorcycleResponseMessage{
		ID:    id,
		Error: err,
	}

	err = responseMessage.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return responseMessage, nil
}

// Validate verifies that a InsertMotorcycleResponseMessage's fields contain valid data.
// Although it is possible that the same rules apply as for the Motorcycle entity, we
// will assume that different rules may be used with this dto.
// Returns nil if the InsertMotorcycleResponseMessage contains valid data, otherwise an error.
func (insertMotorcycleResponseMessage *InsertMotorcycleResponseMessage) Validate() error {
	return validation.ValidateStruct(nil)
}
