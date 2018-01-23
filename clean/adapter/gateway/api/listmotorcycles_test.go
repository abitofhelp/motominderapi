// Package api contains the restful web service.
package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// TestApi_ListMotorcycles_Empty verifies a successful response with no motorcycles.
func TestApi_ListMotorcycles_Empty(t *testing.T) {

	// ARRANGE

	// Configure the application...
	roles := map[enumeration.AuthorizationRole]bool{
		enumeration.AdminAuthorizationRole: true,
	}

	authService, _ := security.NewAuthService(true, roles)
	motorcycleRepository, _ := repository.NewMotorcycleRepository()
	router := httprouter.New()

	// Create an instance of the API web service.
	ourApi, err := NewApi(roles, authService, motorcycleRepository, router)
	if err != nil {
		println("Failed to create an instance of the API web service: &s", err.Error())
		return
	}

	// An http handler wrapper around httprouter's handler.  It permits us to use
	// the test server and httpExpected.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ourApi.ListMotorcyclesHandler(w, r, httprouter.Params{})
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// ACT
	resp, err := http.Get(server.URL)

	// ASSERT
	assert.True(t, resp.StatusCode == http.StatusOK)
}

// TestApi_ListMotorcycles_NotEmpty verifies a successful response with a list of motorcycles.
func TestApi_ListMotorcycles_NotEmpty(t *testing.T) {

	// ARRANGE

	// Configure the application...
	roles := map[enumeration.AuthorizationRole]bool{
		enumeration.AdminAuthorizationRole: true,
	}

	authService, _ := security.NewAuthService(true, roles)
	motorcycleRepository, _ := repository.NewMotorcycleRepository()
	router := httprouter.New()

	// Create an instance of the API web service.
	ourApi, err := NewApi(roles, authService, motorcycleRepository, router)
	if err != nil {
		println("Failed to create an instance of the API web service: &s", err.Error())
		return
	}

	// Insert two motorcycles into the repository using the Api.
	motorcycle, _ := entity.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")
	InsertMotorcycle(ourApi, motorcycle)
	motorcycle, _ = entity.NewMotorcycle("Honda", "Shadow", 2010, "01234567890199999")
	InsertMotorcycle(ourApi, motorcycle)

	// ACT
	resp, err := GetMotorcycles(ourApi)

	// ASSERT
	assert.True(t, resp.StatusCode == http.StatusOK)
}

// GetMotorcycles retrieves the list of motorcycles from the repository using the RESTful API.
// Returns (*response, nil) on success, otherwise (nil, error).
func GetMotorcycles(ourApi *Api) (*http.Response, error) {
	// An http handler wrapper around httprouter's handler.  It permits us to use
	// the test server.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ourApi.ListMotorcyclesHandler(w, r, httprouter.Params{})
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
