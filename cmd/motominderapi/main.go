// Package main is the entrypoint for the API web service.
package main

import (
	"github.com/abitofhelp/motominderapi/clean/configuration/web/api"
)

// Main is the entry point for the API web service.
func main() {

	// Create an instance of the API web service.
	ourApi, err := api.NewApi()
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
