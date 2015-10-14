package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcoleman731/sample-server-go/controller"
)

type IntegrationHandler struct {
	DB *sql.DB
}

func (handler *IntegrationHandler) Index(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing Request to integration Controller\n")
	route := request.RequestURI
	fmt.Printf("Current Route %v\n", route)

	integrationController := controller.IntegrationController{DB: handler.DB}
	switch request.Method {
	case "GET":
		vars := mux.Vars(request)
		fmt.Printf("GET Request with var %v\n", vars)
		integrationController.Get(request)

	case "POST":
		integrationController.Post(request)

	case "UPDATE":
		integrationController.Update(request)

	case "DELETE":
		integrationController.Delete(request)
	}
}
