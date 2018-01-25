// Package api contains the restful web service.
package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/adapter/viewmodel"
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/abitofhelp/motominderapi/clean/domain/typedef"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// TestApi_UpdateMotorcycle verifies a successful response after updating a motorcycle.
func TestApi_UpdateMotorcycle(t *testing.T) {

	// ARRANGE

	// Configure the application...
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
	}

	authService, _ := security.NewAuthService(true, roles)
	repos, _ := repository.NewMotorcycleRepository()
	router := httprouter.New()

	// Create an instance of the API web service.
	ourApi, err := NewApi(roles, authService, repos, router)
	if err != nil {
		println("Failed to create an instance of the API web service: &s", err.Error())
		return
	}

	// Insert a motorcycle into the repository so it can be updated.
	motorcycle, _ := entity.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")
	insertResponse, _ := InsertMotorcycle(ourApi, motorcycle)

	//  Get the response data for the insertion.
	insertionViewModel := viewmodel.InsertMotorcycleViewModel{}
	json.NewDecoder(insertResponse.Body).Decode(&insertionViewModel)
	// Load the fully populated repository entity.
	motorcycle, _, _ = repos.FindByID(insertionViewModel.ID)

	// Change its VIN.
	motorcycle.Vin = "65432109876543210"

	// ACT
	resp, err := UpdateMotorcycle(ourApi, motorcycle.ID, *motorcycle)

	// ASSERT
	assert.True(t, resp.StatusCode == 204)
}

// UpdateMotorcycle updates a motorcycle in the repository using the RESTful API.
// Returns (*response, nil) on success, otherwise (nil, error).
func UpdateMotorcycle(ourApi *Api, id typedef.ID, motorcycle entity.Motorcycle) (*http.Response, error) {

	idText := strconv.Itoa(int(id))

	// An http handler wrapper around httprouter's handler.  It permits us to use
	// the test server.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ourApi.DeleteMotorcycleHandler(w, r, httprouter.Params{httprouter.Param{
			Key:   "id",
			Value: idText,
		}})
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}

	motorcycleJson, _ := json.Marshal(motorcycle)
	url := server.URL + "/" + idText
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(motorcycleJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(motorcycleJson)))

	return client.Do(req)
}
