package main

import (
	"fmt"
	"net/http"

	"github.com/kcoleman731/sample-server-go/controller"
	"github.com/kcoleman731/sample-server-go/model"
)

// We setup our DB as a global variable
var Database = model.NewDatabase()

var Controllers = controller.SetupControllers()

func main() {

	// Log that we have started
	fmt.Printf("Listening on port 8080...\n")

	// Fire up the server
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("Failed to parse form with error: %v\n", err)
		} else {
			fmt.Printf("Form: %v\n", r.Form)
			fmt.Printf("Post Form: %v\n", r.PostForm)
			fmt.Printf("Body: %v\n\n", r.Body)
			fmt.Printf("URL: %v\n\n", r.URL)
		}
	})
	http.ListenAndServe(":8080", nil)
}
