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

	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/operationstatus"
	"github.com/abitofhelp/motominderapi/clean/usecase/request"
	"github.com/abitofhelp/motominderapi/clean/usecase/response"
	"github.com/pkg/errors"
)

/*
TITLE
Delete an existing motorcycle from the motorcycle repository.

DESCRIPTION
User accesses the system to delete motorcycles.

PRIMARY ACTOR
User

PRECONDITIONS
User is logged into system.
User possesses the necessary security authorizations to delete a motorcycle.
A Motorcycle with the ID exists in the repository.
The network and configuration is working properly.

POSTCONDITIONS
User has deleted a motorcycle from the system, unless it didn't exist.

MAIN SUCCESS SCENARIO
1. User selects "Delete Motorcycle..." from the menu.
2. System displays a view in which the user selects a motorcycle to delete.
3. User click the "Submit" button.
4. System deletes the motorcycle from the motorcycle repository, and displays a confirmation message.
5. User clicks the "OK" button, and returns to the primary view.

EXTENSIONS
(3a) The user cannot log into the system.
       System displays an error message saying that authentication has failed,
	   and provides suggestions for resolving the issue.  The User clicks the
	   "OK" button, and returns to the login view.

(3b) The user does not possess the required authorization to delete a motorcycle.
       System displays an error message saying that the user does possess the required
	   security authorizations to delete a motorcycles.  It recommends contacting the
	   System Administrator.  The User clicks the "OK" button, and returns to the
	   primary view.

(3c) A motorcycle with the ID does not exist in the repository.
       System displays an error message indicating that a motorcycle with the
	   ID does not exist.  The User clicks the "OK" button, and
	   returns to the primary view.
*/

// DeleteMotorcycleInteractor is a use case for deleting a motorcycle from the motorcycle repository.
type DeleteMotorcycleInteractor struct {
	MotorcycleRepository contract.MotorcycleRepository
	AuthService          contract.AuthService
}

// NewDeleteMotorcycleInteractor creates a new instance of a DeleteMotorcycleInteractor.
// Returns (nil, error) when there is an error, otherwise (DeleteMotorcycleInteractor, nil).
func NewDeleteMotorcycleInteractor(motorcycleRepository contract.MotorcycleRepository, authService contract.AuthService) (*DeleteMotorcycleInteractor, error) {

	interactor := &DeleteMotorcycleInteractor{
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

// Validate verifies that a DeleteMotorcycleInteractor's fields contain valid data.
// Returns nil if the DeleteMotorcycleInteractor contains valid data, otherwise an error.
func (interactor DeleteMotorcycleInteractor) Validate() error {
	return validation.ValidateStruct(&interactor,
		// MotorcycleRepository is required and cannot be null.
		validation.Field(&interactor.MotorcycleRepository, validation.Required),
		// AuthService is required and cannot be null.
		validation.Field(&interactor.AuthService, validation.Required))
}

// Handle processes the request message and generates the response message.  It is performing the use case.
// The request message is a dto containing the required data for completing the use case.
// On success, the method returns the (response message, nil), otherwise (nil, error).
func (interactor *DeleteMotorcycleInteractor) Handle(requestMessage *request.DeleteMotorcycleRequest) (*response.DeleteMotorcycleResponse, error) {
	// Verify that the user has been properly authenticated.
	if !interactor.AuthService.IsAuthenticated() {
		return response.NewDeleteMotorcycleResponse(requestMessage.ID, operationstatus.NotAuthenticated, errors.New("delete operation failed due to not being authenticated"))
	}

	// Verify that the user has the necessary authorizations.
	if !interactor.AuthService.IsAuthorized(authorizationrole.AdminAuthorizationRole) {
		return response.NewDeleteMotorcycleResponse(requestMessage.ID, operationstatus.NotAuthorized, errors.New("delete operation failed due to not being authorized, so please contact your system administrator"))
	}

	// Delete the motorcycle with ID from the repository.
	status, err := interactor.MotorcycleRepository.Delete(requestMessage.ID)
	if err != nil {
		return response.NewDeleteMotorcycleResponse(requestMessage.ID, status, err)
	}

	// Save the changes.
	status, err = interactor.MotorcycleRepository.Save()
	if err != nil {
		return response.NewDeleteMotorcycleResponse(requestMessage.ID, status, err)
	}

	// Return the successful response message.
	return response.NewDeleteMotorcycleResponse(requestMessage.ID, operationstatus.Ok, nil)
}
