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

	"errors"
	"github.com/abitofhelp/motominderapi/clean/domain/constant"
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/abitofhelp/motominderapi/clean/usecase/requestmessage"
	"github.com/abitofhelp/motominderapi/clean/usecase/responsemessage"
)

/*
TITLE
Insert a new motorcycle make, model, and year into the motorcycle repository.

DESCRIPTION
User accesses the system to add a new motorcycle make, model, and year to it.

PRIMARY ACTOR
User

PRECONDITIONS
User is logged into system.
User possesses the necessary security authorizations to insert a motorcycle.
A motorcycle of the same make, model, and year does not already exist.
The network and configuration is working properly.

POSTCONDITIONS
User has inserted a new motorcycle make, model, and year into the system,
unless it already exists.

MAIN SUCCESS SCENARIO
1. User selects "Add Motorcycle..." from the menu.
2. System displays a view in which the user enters the make, model, and year of the motorcycle.
3. User click the "Submit" button.
4. System inserts the motorcycle into the motorcycle repository, and displays a confirmation message.
5. User clicks the "OK" button, and returns to the primary view.

EXTENSIONS
(3a) The user cannot log into the system.
       System displays an error message saying that authentication has failed,
	   and provides suggestions for resolving the issue.  The User clicks the
	   "OK" button, and returns to the login view.

(3b) The user does not possess the required authorization to insert a motorcycle.
       System displays an error message saying that the user does possess the required
	   security authorizations to insert a motorcycle.  It recommends contacting the
	   System Administrator.  The User clicks the "OK" button, and returns to the
	   primary view.

(3c) A motorcycle with the same make, model, and year already exists.
       System displays an error message indicating that a motorcycle with the same
	   make, model, and year already exists.  The User clicks the "OK" button, and
	   returns to the primary view.

(3d) Inserting the motorcycle into the repository failed due to configuration issues.
       System displays an error message saying that the insertion of the motorcycle was
	   unsuccessful due to configuration issues.  The user can click "Retry" or "Cancel".
	   Cancel will return to the primary view.  If the configuration is operating properly, System returns to (4), otherwise System goes to (3d).
*/

// InsertMotorcycleInteractor is a use case for adding a motorcycle to the motorcycle repository.
type InsertMotorcycleInteractor struct {
	MotorcycleRepository contract.MotorcycleRepository
	AuthService          contract.AuthService
}

// NewInsertMotorcycleInteractor creates a new instance of a InsertMotorcycleInteractor.
// Returns (nil, error) when there is an error, otherwise (InsertMotorcycleInteractor, nil).
func NewInsertMotorcycleInteractor(motorcycleRepository contract.MotorcycleRepository, authService contract.AuthService) (*InsertMotorcycleInteractor, error) {

	interactor := &InsertMotorcycleInteractor{
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

// Validate verifies that a InsertMotorcycleInteractor's fields contain valid data.
// Returns nil if the InsertMotorcycleInteractor contains valid data, otherwise an error.
func (insertMotorcycleInteractor InsertMotorcycleInteractor) Validate() error {
	return validation.ValidateStruct(&insertMotorcycleInteractor,
		// MotorcycleRepository is required and cannot be null.
		validation.Field(&insertMotorcycleInteractor.MotorcycleRepository, validation.Required),
		// AuthService is required and cannot be null.
		validation.Field(&insertMotorcycleInteractor.AuthService, validation.Required))
}

// Handle processes the request message and generates the response message.  It is performing the use case.
// The request message is a dto containing the required data for completing the use case.
// On success, the method returns the (response message, nil), otherwise (nil, error).
func (insertMotorcycleInteractor *InsertMotorcycleInteractor) Handle(requestMessage *requestmessage.InsertMotorcycleRequestMessage) (*responsemessage.InsertMotorcycleResponseMessage, error) {
	// Verify that the user has been properly authenticated.
	if !insertMotorcycleInteractor.AuthService.IsAuthenticated() {
		return responsemessage.NewInsertMotorcycleResponseMessage(constant.InvalidEntityID, errors.New("insert operation failed due to not being authenticated, so please contact your system administrator"))
	}

	// Verify that the user has the necessary authorizations.
	if !insertMotorcycleInteractor.AuthService.IsAuthorized(enumeration.AdminAuthorizationRole) {
		return responsemessage.NewInsertMotorcycleResponseMessage(constant.InvalidEntityID, errors.New("insert operation failed due to not having the required user authorization roles, so please contact your system administrator"))
	}

	// Verify that a motorcycle with the same vin does not exist.
	moto, err := insertMotorcycleInteractor.MotorcycleRepository.FindByVin(requestMessage.Vin)
	if moto != nil {
		return responsemessage.NewInsertMotorcycleResponseMessage(constant.InvalidEntityID, errors.New("insert operation failed due to a motorcycle with the same VIN already existing in the repository"))
	}

	// Create a new Motorcycle entity.
	motorcycle, err := entity.NewMotorcycle(requestMessage.Make, requestMessage.Model, requestMessage.Year, requestMessage.Vin)
	if err != nil {
		return responsemessage.NewInsertMotorcycleResponseMessage(constant.InvalidEntityID, err)
	}

	// Insert the new motorcycle entity into the repository.
	motorcycle, err = insertMotorcycleInteractor.MotorcycleRepository.Insert(motorcycle)
	if err != nil {
		return responsemessage.NewInsertMotorcycleResponseMessage(constant.InvalidEntityID, err)
	}

	// Save the changes.
	err = insertMotorcycleInteractor.MotorcycleRepository.Save()
	if err != nil {
		return responsemessage.NewInsertMotorcycleResponseMessage(constant.InvalidEntityID, err)
	}

	// Return the successful response message.
	return responsemessage.NewInsertMotorcycleResponseMessage(motorcycle.ID, nil)
}
