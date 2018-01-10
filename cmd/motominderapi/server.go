package main

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"

	// Motominder's entities packages
	"github.com/abitofhelp/motominderapi/clean/domain/entities"
)

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// Get a Motorcycle resource
	router.GET("/motorcycles", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// A motorcycle
		moto := entities.Motorcycle{
			Id:    123,
			Make:  "KTM",
			Model: "350 EXC-F",
			Year:  2018,
		}

		// Marshal moto into JSON structure
		jsonMoto, err := json.Marshal(moto)
		if err == nil {
			// Write content-type, statuscode, payload
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", jsonMoto)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
