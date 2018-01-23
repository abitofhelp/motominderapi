// Package api contains the restful web service.
package api

import (
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

// TestApi_DeleteMotorcycle verifies a successful response after inserting a motorcycle.
func TestApi_DeleteMotorcycle(t *testing.T) {

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

	// Insert a motorcycle into the repository so it can be deleted.
	motorcycle, _ := entity.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")
	insertResponse, _ := InsertMotorcycle(ourApi, motorcycle)

	//  Get the response data for the insertion.
	insertionViewModel := viewmodel.InsertMotorcycleViewModel{}
	json.NewDecoder(insertResponse.Body).Decode(&insertionViewModel)

	// ACT
	resp, err := DeleteMotorcycle(ourApi, insertionViewModel.ID)

	// ASSERT
	assert.True(t, resp.StatusCode == 204)
}

// DeleteMotorcycle deletes a motorcycle from the repository using the RESTful API.
// Returns (*response, nil) on success, otherwise (nil, error).
func DeleteMotorcycle(ourApi *Api, id typedef.ID) (*http.Response, error) {

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
	url := server.URL + "/" + idText
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
