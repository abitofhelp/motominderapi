// Package presenter performs the translation of a response message into a view model.
package presenter

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/viewmodel"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
	"github.com/go-ozzo/ozzo-validation"
	errors "github.com/pjebs/jsonerror"
)

// DeleteMotorcyclePresenter translates the response message from the DeleteMotorcycleInteractor to a view model.
type DeleteMotorcyclePresenter struct {
}

// NewDeleteMotorcyclePresenter creates a new instance of a DeleteMotorcyclePresenter.
// Returns (instance of DeleteMotorcyclePresenter, nil) on success, otherwise (nil, error).
func NewDeleteMotorcyclePresenter() (*DeleteMotorcyclePresenter, error) {

	presenter := &DeleteMotorcyclePresenter{}

	// All okay
	return presenter, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of DeleteMotorcycleViewModel, nil) on success, otherwise (nil, error)
func (presenter *DeleteMotorcyclePresenter) Handle(responseMessage *response.DeleteMotorcycleResponse) (*viewmodel.DeleteMotorcycleViewModel, error) {
	if responseMessage.Error != nil {
		return viewmodel.NewDeleteMotorcycleViewModel(responseMessage.ID, "Failed to delete the motorcycle.", responseMessage.Error)
	}

	return viewmodel.NewDeleteMotorcycleViewModel(responseMessage.ID, "Successfully delete the motorcycle.", responseMessage.Error)
}

// Validate verifies that a DeleteMotorcyclePresenter's fields contain valid data.
// Returns (an instance of DeleteMotorcyclePresenter, nil) on success, otherwise (nil, error)
func (presenter DeleteMotorcyclePresenter) Validate() error {
	err := validation.ValidateStruct(&presenter)

	if err != nil {
		return errors.New(enumeration.StatusInternalServerError, enumeration.StatusText(enumeration.StatusInternalServerError), err.Error())
	}

	return nil
}
