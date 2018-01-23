// Package interactor contains use cases, which contain the application specific business rules.
// Interactors encapsulate and implement all of the use cases of the system.  They orchestrate the
// flow of data to and from the entity, and can rely on their business rules to achieve the goals
// of the use case.  They do not have any dependencies, and are totally isolated from things like
// a database, UI or special frameworks, which exist in the outer rings.  They Will almost certainly
// require refactoring if details of the use case requirements change.
package interactor

import (
	"github.com/abitofhelp/motominderapi/clean/domain/contract"
	"github.com/go-ozzo/ozzo-validation"

	"github.com/pkg/errors"

	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
)

/*
TITLE
Get an unsorted list of motorcycles from the motorcycle repository.

DESCRIPTION
User accesses the system to get a list of motorcycles.

PRIMARY ACTOR
User

PRECONDITIONS
User is logged into system.
User possesses the necessary security authorizations to insert a motorcycle.
The network and configuration is working properly.

POSTCONDITIONS
User has received a list of motorcycles from the system, and the list can be empty.

MAIN SUCCESS SCENARIO
1. User selects "Get Motorcycles" from the menu.
2. System displays a view showing the unsorted list of motorcycles.
3. User clicks the "OK" button, and returns to the primary view.

EXTENSIONS
(3a) The user cannot log into the system.
       System displays an error message saying that authentication has failed,
	   and provides suggestions for resolving the issue.  The User clicks the
	   "OK" button, and returns to the login view.

(3b) The user does not possess the required authorization to get a list of motorcycles.
       System displays an error message saying that the user does possess the required
	   security authorizations.  It recommends contacting the
	   System Administrator.  The User clicks the "OK" button, and returns to the
	   primary view.
*/

// ListMotorcyclesInteractor is a use case for getting a list of motorcycles from the motorcycle repository.
type ListMotorcyclesInteractor struct {
	MotorcycleRepository contract.MotorcycleRepository
	AuthService          contract.AuthService
}

// NewListMotorcyclesInteractor creates a new instance of a ListMotorcyclesInteractor.
// Returns (nil, error) when there is an error, otherwise (ListMotorcyclesInteractor, nil).
func NewListMotorcyclesInteractor(motorcycleRepository contract.MotorcycleRepository, authService contract.AuthService) (*ListMotorcyclesInteractor, error) {

	interactor := &ListMotorcyclesInteractor{
		MotorcycleRepository: motorcycleRepository,
		AuthService:          authService,
	}

	// Validate the interactor
	err := interactor.Validate()
	if err != nil {
		return nil, err
	}

	// All okay
	return interactor, nil
}

// Validate verifies that a ListMotorcyclesInteractor's fields contain valid data.
// Returns nil if the ListMotorcyclesInteractor contains valid data, otherwise an error.
func (interactor ListMotorcyclesInteractor) Validate() error {
	return validation.ValidateStruct(&interactor,
		// MotorcycleRepository is required and cannot be null.
		validation.Field(&interactor.MotorcycleRepository, validation.Required),
		// AuthService is required and cannot be null.
		validation.Field(&interactor.AuthService, validation.Required))
}

// Handle processes the request message and generates the response message.  It is performing the use case.
// The request message is a dto containing the required data for completing the use case.
// On success, the method returns the (response message, nil), otherwise (nil, error).
func (interactor *ListMotorcyclesInteractor) Handle(requestMessage *request.ListMotorcyclesRequest) (*response.ListMotorcyclesResponse, error) {
	// Verify that the user has been properly authenticated.
	if !interactor.AuthService.IsAuthenticated() {
		return response.NewListMotorcyclesResponse(nil, errors.New("list operation failed due to not being authenticated"))
	}

	// Verify that the user has the necessary authorizations.
	if !interactor.AuthService.IsAuthorized(authorizationrole.AdminAuthorizationRole) {
		return response.NewListMotorcyclesResponse(nil, errors.New("list operation failed due to not being authorized, so please contact your system administrator"))
	}

	// Get the list of motorcycles from the repository.
	motorcycles, err := interactor.MotorcycleRepository.List()
	if err != nil {
		return response.NewListMotorcyclesResponse(nil, err)
	}

	// Return the successful response message.
	return response.NewListMotorcyclesResponse(motorcycles, nil)
}
