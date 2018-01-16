package api

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"

	// Motominder's entities packages
	"github.com/abitofhelp/motominderapi/clean/adapters/gateways/repositories"
	"github.com/abitofhelp/motominderapi/clean/adapters/gateways/security"
	"github.com/abitofhelp/motominderapi/clean/adapters/presenters"
	"github.com/abitofhelp/motominderapi/clean/configuration/web/api/dtos"
	"github.com/abitofhelp/motominderapi/clean/domain/enumerations"
	"github.com/abitofhelp/motominderapi/clean/usecases/interactors"
	"github.com/abitofhelp/motominderapi/clean/usecases/requestmessages"
	"github.com/go-ozzo/ozzo-validation"
)

// Api is a web service.
type Api struct {
	Roles                map[enumerations.AuthorizationRole]bool
	AuthService          *security.AuthService
	MotorcycleRepository *repositories.MotorcycleRepository
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
func NewApi() (*Api, error) {

	roles := map[enumerations.AuthorizationRole]bool{
		enumerations.AdminAuthorizationRole: true,
	}

	authService, _ := security.NewAuthService(true, roles)
	repo, _ := repositories.NewMotorcycleRepository()
	router := httprouter.New()

	api := &Api{
		Roles:                roles,
		AuthService:          authService,
		MotorcycleRepository: repo,
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

		// Stub a motorcycle to be populated from the body of the request.
		dto := dtos.InsertMotorcycleDto{}

		// Populate the motorcycle from the request body.
		json.NewDecoder(r.Body).Decode(&dto)

		// Create the request, process it, and get the resulting view model or error.
		request, err := requestmessages.NewInsertMotorcycleRequestMessage(dto.Make, dto.Model, dto.Year, dto.Vin)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			return
		}

		interactor, err := interactors.NewInsertMotorcycleInteractor(api.MotorcycleRepository, api.AuthService)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}

		response, err := interactor.Handle(request)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}

		presenter, err := presenters.NewInsertMotorcycleResponseMessagePresenter()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}

		viewModel, err := presenter.Handle(response)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			return
		}

		// Marshal provided interface into JSON structure
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
