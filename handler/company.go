package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcoleman731/sample-server-go/controller"
)

type CompanyHandler struct {
	DB *sql.DB
}

func (handler *CompanyHandler) Index(writter http.ResponseWriter, request *http.Request) {
	controller := controller.CompanyController{DB: handler.DB, Params: mux.Vars(request)}
	fmt.Printf("Routing Request to Company Controller\n")

	route := request.RequestURI
	fmt.Printf("Current Route %v\n", route)

	switch request.Method {
	case "GET":
		controller.Get(request)

	case "POST":
		controller.Post(request)

	case "UPDATE":
		controller.Update(request)

	case "DELETE":
		controller.Delete(request)
	}
}

func (handler *CompanyHandler) Products(writter http.ResponseWriter, request *http.Request) {

}

func (handler *CompanyHandler) Features(writter http.ResponseWriter, request *http.Request) {

}

func (handler *CompanyHandler) Integrations(writter http.ResponseWriter, request *http.Request) {

}

func (handler *CompanyHandler) Requests(writter http.ResponseWriter, request *http.Request) {

}
