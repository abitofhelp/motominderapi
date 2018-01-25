// Package api contains the restful web service.
package api

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	// Motominder's entity packages
	"github.com/abitofhelp/motominderapi/clean/adapter/dto"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/adapter/presenter"
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
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
	return validation.ValidateStruct(&api,
		validation.Field(&api.Roles, validation.Required),
		validation.Field(&api.AuthService, validation.Required),
		validation.Field(&api.MotorcycleRepository, validation.Required),
		validation.Field(&api.Router, validation.Required))
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

	// Set up the handler to get a particular motorcycle from the repository.
	api.Router.GET("/api/motorcycles/:id", api.GetMotorcycleHandler)

	// Set up the handler to insert a new motorcycle into the repository.
	api.Router.POST("/api/motorcycles", api.InsertMotorcycleHandler)

	// Set up the handler to update a motorcycle in the repository.
	api.Router.PUT("/api/motorcycles/:id", api.UpdateMotorcycleHandler)

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
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err)
		return
	}

	listInteractor, err := interactor.NewListMotorcyclesInteractor(api.MotorcycleRepository, api.AuthService)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	listResponse, err := listInteractor.Handle(listRequest)
	if err != nil {
		w.WriteHeader(int(listResponse.Status))
		log.WithError(err)
		return
	}

	listPresenter, err := presenter.NewListMotorcyclesPresenter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	viewModel, err := listPresenter.Handle(listResponse)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Marshal provided viewModel into JSON structure
	uj, err := json.Marshal(viewModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Write content-type, status code, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// GetMotorcycleHandler gets a particular motorcycle from the repository.
func (api *Api) GetMotorcycleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err)
		return
	}

	// Create the motorcycleRequest, process it, and get the resulting view model or error.
	getRequest, err := request.NewGetMotorcycleRequest(typedef.ID(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err)
		return
	}

	getInteractor, err := interactor.NewGetMotorcycleInteractor(api.MotorcycleRepository, api.AuthService)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	getResponse, err := getInteractor.Handle(getRequest)
	if err != nil {
		w.WriteHeader(int(getResponse.Status))
		log.WithError(err)
		return
	}

	getPresenter, err := presenter.NewGetMotorcyclePresenter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	viewModel, err := getPresenter.Handle(getResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Marshal provided viewModel into JSON structure
	uj, err := json.Marshal(viewModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Write content-type, status code, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

// DeleteMotorcycleHandler removes a motorcycle from the repository.
func (api *Api) DeleteMotorcycleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err)
		return
	}

	// Create the motorcycleRequest, process it, and get the resulting view model or error.
	deleteRequest, err := request.NewDeleteMotorcycleRequest(typedef.ID(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err)
		return
	}

	deleteInteractor, err := interactor.NewDeleteMotorcycleInteractor(api.MotorcycleRepository, api.AuthService)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	deleteResponse, err := deleteInteractor.Handle(deleteRequest)
	if err != nil {
		w.WriteHeader(int(deleteResponse.Status))
		log.WithError(err)
		return
	}

	deletePresenter, err := presenter.NewDeleteMotorcyclePresenter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	viewModel, err := deletePresenter.Handle(deleteResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Marshal provided viewModel into JSON structure
	_, err = json.Marshal(viewModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Write status code
	w.WriteHeader(http.StatusNoContent)
}

// UpdateMotorcycleHandler updates an existing motorcycle in the repository.
func (api *Api) UpdateMotorcycleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err)
		return
	}

	// Stub a motorcycle to be populated from the body of the motorcycleRequest.
	motorcycle := &entity.Motorcycle{}

	// Populate the motorcycle from the motorcycleRequest body.
	json.NewDecoder(r.Body).Decode(&motorcycle)

	// Create the motorcycleRequest, process it, and get the resulting view model or error.
	motorcycleRequest, err := request.NewUpdateMotorcycleRequest(typedef.ID(id), motorcycle) //typedef.ID(id), motorcycle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err)
		return
	}

	motorcycleInteractor, err := interactor.NewUpdateMotorcycleInteractor(api.MotorcycleRepository, api.AuthService)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	updateResponse, err := motorcycleInteractor.Handle(motorcycleRequest)
	if err != nil {
		w.WriteHeader(int(updateResponse.Status))
		log.WithError(err)
		return
	}

	motorcyclePresenter, err := presenter.NewUpdateMotorcyclePresenter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	viewModel, err := motorcyclePresenter.Handle(updateResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Marshal provided viewModel into JSON structure
	_, err = json.Marshal(viewModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Write status code
	w.WriteHeader(http.StatusNoContent)
}

// InsertMotorcycleHandler adds a new motorcycle to the repository.
func (api *Api) InsertMotorcycleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Stub a motorcycle to be populated from the body of the motorcycleRequest.
	motorcycleDto := dto.MutableMotorcycleDto{}

	// Populate the motorcycle from the motorcycleRequest body.
	json.NewDecoder(r.Body).Decode(&motorcycleDto)

	// Create the motorcycleRequest, process it, and get the resulting view model or error.
	motorcycleRequest, err := request.NewInsertMotorcycleRequest(motorcycleDto.Make, motorcycleDto.Model, motorcycleDto.Year, motorcycleDto.Vin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err)
		return
	}

	motorcycleInteractor, err := interactor.NewInsertMotorcycleInteractor(api.MotorcycleRepository, api.AuthService)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	insertResponse, err := motorcycleInteractor.Handle(motorcycleRequest)
	if err != nil {
		w.WriteHeader(int(insertResponse.Status))
		log.WithError(err)
		return
	}

	motorcyclePresenter, err := presenter.NewInsertMotorcyclePresenter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	viewModel, err := motorcyclePresenter.Handle(insertResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Marshal provided viewModel into JSON structure
	uj, err := json.Marshal(viewModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithError(err)
		return
	}

	// Write content-type, status code, payload
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "/api/motorcycles/"+fmt.Sprintf("%d", insertResponse.ID))
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

// init configures the API for use.
func (api *Api) init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
