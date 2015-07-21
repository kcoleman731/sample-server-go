package main

import (
	"fmt"
	"net/http"

	"github.com/kcoleman731/test-server/controller"
	"github.com/kcoleman731/test-server/model"
)

// We setup our DB as a global variable
var Database = model.NewDatabase()
var Controllers = controller.SetupControllers()

func main() {

	// Log that we have started
	fmt.Printf("Listening on port 8080...\n")

	// Fire up the server
	http.ListenAndServe(":8080", RouteRequest())
}
