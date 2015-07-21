package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcoleman731/test-server/controller"
)

type CompanyHandler struct {
	DB *sql.DB
}

func (handler *CompanyHandler) Index(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing Request to Company Controller\n")

	route := request.RequestURI
	fmt.Printf("Current Route %v\n", route)

	controller := controller.CompanyController{DB: handler.DB}
	switch request.Method {
	case "GET":
		fmt.Printf("GET Request with parameters %v\n", mux.Vars(request))
		controller.Get(request)

	case "POST":
		controller.Post(request)

	case "UPDATE":
		controller.Update(request)

	case "DELETE":
		ontroller.Delete(request)
	}
}

func (handler *CompanyHandler) Products(request *http.Request) {

}

func (handler *CompanyHandler) Features(request *http.Request) {

}

func (handler *CompanyHandler) Integrations(request *http.Request) {

}
