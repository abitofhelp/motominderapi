// Package presenter performs the translation of a response message into a view model.
package presenter

import (
	"fmt"
	"github.com/abitofhelp/motominderapi/clean/adapter/dto"
	"github.com/abitofhelp/motominderapi/clean/adapter/viewmodel"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
	"github.com/go-ozzo/ozzo-validation"
)

// GetMotorcyclePresenter translates the response message from the GetMotorcycleInteractor to a view model.
type GetMotorcyclePresenter struct {
}

// NewGetMotorcyclePresenter creates a new instance of a GetMotorcyclePresenter.
// Returns (instance of GetMotorcyclePresenter, nil) on success, otherwise (nil, error).
func NewGetMotorcyclePresenter() (*GetMotorcyclePresenter, error) {

	presenter := &GetMotorcyclePresenter{}

	// All okay
	return presenter, nil
}

// Handle performs the translation of the response message into a view model.
// Returns (instance of GetMotorcyclePresenter, nil) on success, otherwise (nil, error)
func (presenter *GetMotorcyclePresenter) Handle(responseMessage *response.GetMotorcycleResponse) (*viewmodel.GetMotorcycleViewModel, error) {
	if responseMessage.Error != nil {
		return viewmodel.NewGetMotorcycleViewModel(nil, "Failed to get the motorcycle.", responseMessage.Error)
	}

	motorcycleDto, err := dto.NewImmutableMotorcycleDto(*responseMessage.Motorcycle)
	if err != nil {
		return viewmodel.NewGetMotorcycleViewModel(nil, "Failed to create an immutable motorcycle.", err)
	}

	return viewmodel.NewGetMotorcycleViewModel(motorcycleDto, fmt.Sprintf("Successfully retrieved the motorcycle with ID %d.", motorcycleDto.ID), responseMessage.Error)
}

// Validate verifies that a GetMotorcyclePresenter's fields contain valid data.
// Returns (an instance of GetMotorcyclePresenter, nil) on success, otherwise (nil, error)
func (presenter GetMotorcyclePresenter) Validate() error {
	return validation.ValidateStruct(&presenter)
}
