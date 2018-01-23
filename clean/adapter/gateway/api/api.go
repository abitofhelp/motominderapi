// Package api contains the restful web service.
package api

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	errors "github.com/pjebs/jsonerror"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	// Motominder's entity packages
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/api/dto"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/adapter/presenter"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/operationstatus"
	"github.com/abitofhelp/motominderapi/clean/usecase/interactor"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/go-ozzo/ozzo-validation"
)

// Api is a web service.
type Api struct {
	Roles                map[authorizationrole.AuthorizationRole]bool
	AuthService          *security.AuthService
	MotorcycleRepository *repository.MotorcycleRepository
	Router               *httprouter.Router
}

// Validate verifies that a api's fields contain valid data.
// Returns nil on success, otherwise error.
func (api Api) Validate() error {
	err := validation.ValidateStruct(&api,
		validation.Field(&api.Roles, validation.Required),
		validation.Field(&api.AuthService, validation.Required),
		validation.Field(&api.MotorcycleRepository, validation.Required),
		validation.Field(&api.Router, validation.Required))

	if err != nil {
		return errors.New(operationstatus.StatusInternalServerError, operationstatus.StatusText(operationstatus.StatusInternalServerError), err.Error())
	}

	return nil
}

// NewApi creates a new instance of an Api.
// Returns (an instance of APi, nil), otherwise (nil, error)
func NewApi(roles map[authorizationrole.AuthorizationRole]bool, authService *security.AuthService, motorcycleRepository *repository.MotorcycleRepository, router *httprouter.Router) (*Api, error) {

	api := &Api{
		Roles:                roles,
		AuthService:          authService,
		MotorcycleRepository: motorcycleRepository,
		Router:               router,
	}

	// Initialize logging
	api.init()

	err := api.Validate()
	if err != nil {
		return nil, err
	}

	// Configure the router.
	api.configureRouter()

	// All okay
	return api, nil
}

// configureRouter sets up the router's handlers and endpoints.
// Returns nil on success, otherwise error.
func (api *Api) configureRouter() error {

	// Set up the handler to get a list of motorcycles from the repository.
	api.Router.GET("/api/motorcycles", api.ListMotorcyclesHandler)

	// Set up the handler to insert a new motorcycle into the repository.
	api.Router.POST("/api/motorcycles", api.InsertMotorcycleHandler)

	// Set up the handler to delete a motorcycle from the repository.
	api.Router.DELETE("/api/motorcycles/:id", api.DeleteMotorcycleHandler)

	return nil
}

// Start launches the web service.
// Returns nil on success, otherwise error.
func (api *Api) Start() error {
	println("Starting the API server...")
	log.Fatal(http.ListenAndServe(":8080", api.Router))
	return nil
}

// Stop terminates the web service.
// Returns nil on success, otherwise error.
func (api *Api) Stop() error {
	println("Stopping the API server...")
	return nil
}

// ListMotorcyclesHandler processes requests to get a list of motorcycles from the repository.
func (api *Api) ListMotorcyclesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Create the listRequest, process it, and get the resulting view model or error.
	listRequest, err := request.NewListMotorcyclesRequest()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	listInteractor, err := interactor.NewListMotorcyclesInteractor(api.MotorcycleRepository, api.AuthService)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	listResponse, err := listInteractor.Handle(listRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	listPresenter, err := presenter.NewListMotorcyclesPresenter()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	viewModel, err := listPresenter.Handle(listResponse)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	// Marshal provided contract into JSON structure
	uj, err := json.Marshal(viewModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	// Write content-type, status code, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)

}

// DeleteMotorcycleHandler removes a motorcycle from the respository.
func (api *Api) DeleteMotorcycleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	// Create the motorcycleRequest, process it, and get the resulting view model or error.
	deleteRequest, err := request.NewDeleteMotorcycleRequest(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	deleteInteractor, err := interactor.NewDeleteMotorcycleInteractor(api.MotorcycleRepository, api.AuthService)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	deleteResponse, err := deleteInteractor.Handle(deleteRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	deletePresenter, err := presenter.NewDeleteMotorcyclePresenter()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	deleteViewModel, err := deletePresenter.Handle(deleteResponse)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	// Marshal provided contract into JSON structure
	uj, _ := json.Marshal(deleteViewModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	// Write content-type, status code, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "%s", uj)
}

// InsertMotorcycleHandler adds a new motorcycle to the repository.
func (api *Api) InsertMotorcycleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Stub a motorcycle to be populated from the body of the motorcycleRequest.
	motorcycleDto := dto.InsertMotorcycleDto{}

	// Populate the motorcycle from the motorcycleRequest body.
	json.NewDecoder(r.Body).Decode(&motorcycleDto)

	// Create the motorcycleRequest, process it, and get the resulting view model or error.
	motorcycleRequest, err := request.NewInsertMotorcycleRequest(motorcycleDto.Make, motorcycleDto.Model, motorcycleDto.Year, motorcycleDto.Vin)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	motorcycleInteractor, err := interactor.NewInsertMotorcycleInteractor(api.MotorcycleRepository, api.AuthService)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	response, err := motorcycleInteractor.Handle(motorcycleRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	motorcyclePresenter, err := presenter.NewInsertMotorcyclePresenter()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	viewModel, err := motorcyclePresenter.Handle(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	// Marshal provided contract into JSON structure
	uj, _ := json.Marshal(viewModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(operationstatus.ToHttpStatus(err.(errors.JE).Code))
		log.WithError(err)
		return
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "/api/motorcycles/"+fmt.Sprintf("%d", response.ID))
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func (api *Api) init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
