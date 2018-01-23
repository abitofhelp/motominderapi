// Package presenter performs the translation of a response message into a view model.
package presenter

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/viewmodel"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
	"github.com/go-ozzo/ozzo-validation"
	errors "github.com/pjebs/jsonerror"
)

// ListMotorcyclesPresenter translates the response message from the ListMotorcyclesInteractor to a view model.
type ListMotorcyclesPresenter struct {
}

// NewListMotorcyclesPresenter creates a new instance of a ListMotorcyclesPresenter.
// Returns (instance of ListMotorcyclesPresenter, nil) on success, otherwise (nil, error).
func NewListMotorcyclesPresenter() (*ListMotorcyclesPresenter, error) {

	presenter := &ListMotorcyclesPresenter{}

	// All okay
	return presenter, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of ListMotorcyclesPresenter, nil) on success, otherwise (nil, error)
func (presenter *ListMotorcyclesPresenter) Handle(responseMessage *response.ListMotorcyclesResponse) (*viewmodel.ListMotorcyclesViewModel, error) {
	if responseMessage.Error != nil {
		return viewmodel.NewListMotorcyclesViewModel(nil, "Failed to get the list of motorcycles.", responseMessage.Error)
	}

	return viewmodel.NewListMotorcyclesViewModel(responseMessage.Motorcycles, "Successfully retrieved the list of motorcycles.", responseMessage.Error)
}

// Validate verifies that a ListMotorcyclesPresenter's fields contain valid data.
// Returns (an instance of ListMotorcyclesPresenter, nil) on success, otherwise (nil, error)
func (presenter ListMotorcyclesPresenter) Validate() error {
	err := validation.ValidateStruct(&presenter)

	if err != nil {
		return errors.New(enumeration.StatusInternalServerError, enumeration.StatusText(enumeration.StatusInternalServerError), err.Error())
	}

	return nil
}
