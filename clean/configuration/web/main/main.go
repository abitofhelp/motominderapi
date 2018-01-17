// Package main is the entrypoint for the API web service.
package main

import (
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/api"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/repository"
	"github.com/abitofhelp/motominderapi/clean/adapter/gateway/security"
	"github.com/abitofhelp/motominderapi/clean/domain/enumeration"
	"github.com/julienschmidt/httprouter"
)

// Main is the entry point for the API web service.
func main() {

	// Configure the application...
	roles := map[enumeration.AuthorizationRole]bool{
		enumeration.AdminAuthorizationRole: true,
	}

	authService, _ := security.NewAuthService(true, roles)
	motorcycleRepository, _ := repository.NewMotorcycleRepository()
	router := httprouter.New()

	// Create an instance of the API web service.
	ourApi, err := api.NewApi(roles, authService, motorcycleRepository, router)
	if err != nil {
		println("Failed to create an instance of the API web service: &s", err.Error())
		return
	}

	// Start the API web service.
	err = ourApi.Start()

	if err != nil {
		println("API has had a failure, and is exiting: &s", err.Error())
		return
	}

	println("API is exiting after normal processing.")
}
