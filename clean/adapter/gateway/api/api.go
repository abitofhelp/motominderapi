// Package api contains the restful web service.
package api

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"

	// Motominder's entity packages
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/api/dto"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/adapter/presenter"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/abitofhelp/motominderapi/clean/usecase/interactor"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/go-ozzo/ozzo-validation"
)

// Api is a web service.
type Api struct {
	Roles                map[enumeration.AuthorizationRole]bool
	AuthService          *security.AuthService
	MotorcycleRepository *repository.MotorcycleRepository
	Router               *httprouter.Router
}

// Validate verifies that a api's fields contain valid data.
// Returns nil on success, otherwise error.
func (api Api) Validate() error {
	return validation.ValidateStruct(&api,
		validation.Field(&api.Roles, validation.Required),
		validation.Field(&api.AuthService, validation.Required),
		validation.Field(&api.MotorcycleRepository, validation.Required),
		validation.Field(&api.Router, validation.Required))
}

// NewApi creates a new instance of an Api.
// Returns (an instance of APi, nil), otherwise (nil, error)
func NewApi(roles map[enumeration.AuthorizationRole]bool, authService *security.AuthService, motorcycleRepository *repository.MotorcycleRepository, router *httprouter.Router) (*Api, error) {

	api := &Api{
		Roles:                roles,
		AuthService:          authService,
		MotorcycleRepository: motorcycleRepository,
		Router:               router,
	}

	err := api.Validate()
	if err != nil {
		return nil, err
	}

	// Configure the router.
	api.configureRouter()

	// All okay
	return api, nil
}

func (api *Api) configureRouter() error {
	// Create a new motorcycle in the repository.
	api.Router.POST("/motorcycles", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// Stub a motorcycle to be populated from the body of the motorcycleRequest.
		motorcycleDto := dto.InsertMotorcycleDto{}

		// Populate the motorcycle from the motorcycleRequest body.
		json.NewDecoder(r.Body).Decode(&motorcycleDto)

		// Create the motorcycleRequest, process it, and get the resulting view model or error.
		motorcycleRequest, err := request.NewInsertMotorcycleRequestMessage(motorcycleDto.Make, motorcycleDto.Model, motorcycleDto.Year, motorcycleDto.Vin)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			return
		}

		motorcycleInteractor, err := interactor.NewInsertMotorcycleInteractor(api.MotorcycleRepository, api.AuthService)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}

		response, err := motorcycleInteractor.Handle(motorcycleRequest)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}

		motorcyclePresenter, err := presenter.NewInsertMotorcyclePresenter()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}

		viewModel, err := motorcyclePresenter.Handle(response)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}

		// Marshal provided contract into JSON structure
		uj, _ := json.Marshal(viewModel)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", uj)
	})

	/*
		// Get the motorcycles from the repository.
		api.Router.GET("/motorcycles", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			// Marshal moto into JSON structure
			jsonMoto, err := json.Marshal(moto)
			if err == nil {
				// Write content-type, statuscode, payload
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
				fmt.Fprintf(w, "%s", jsonMoto)
			}
		})
	*/

	return nil
}

// Start launches the web service.
func (api *Api) Start() error {
	println("Starting the API server...")
	log.Fatal(http.ListenAndServe(":8080", api.Router))
	return nil
}

// Stop terminates the web service.
func (api *Api) Stop() error {
	println("Stopping the API server...")
	return nil
}
