// Package api contains the restful web service.
package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"bytes"

	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/domain/entity"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration/authorizationrole"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// TestApi_InsertMotorcycle verifies a successful response after inserting a motorcycle.
func TestApi_InsertMotorcycle(t *testing.T) {

	// ARRANGE

	// Configure the application...
	roles := map[authorizationrole.AuthorizationRole]bool{
		authorizationrole.AdminAuthorizationRole: true,
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

	motorcycle, _ := entity.NewMotorcycle("Honda", "Shadow", 2006, "01234567890123456")

	// ACT
	resp, err := InsertMotorcycle(ourApi, motorcycle)

	// ASSERT
	assert.True(t, resp.StatusCode == 201)
	assert.Equal(t, "/api/motorcycles/1", resp.Header.Get("Location"))
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

// InsertMotorcycle inserts a motorcycle into the repository using the RESTful API.
// Returns (*response, nil) on success, otherwise (nil, error).
func InsertMotorcycle(ourApi *Api, motorcycle *entity.Motorcycle) (*http.Response, error) {
	// An http handler wrapper around httprouter's handler.  It permits us to use
	// the test server.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ourApi.PostMotorcycleHandler(w, r, httprouter.Params{})
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}

	motorcycleJson, _ := json.Marshal(motorcycle)

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(motorcycleJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(motorcycleJson)))

	return client.Do(req)
}
