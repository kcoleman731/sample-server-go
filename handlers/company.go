package handlers

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
	fmt.Printf("Routing %v Request to Company Controller\n", request.Method)

	controller := controller.CompanyController{DB: handler.DB, Params: mux.Vars(request)}
	switch request.Method {
	case "GET":
		writeResponse(writter, controller.Get(request))

	case "POST":
		writeResponse(writter, controller.Post(request))

	case "PUT":
		writeResponse(writter, controller.Update(request))

	case "DELETE":
		writeResponse(writter, controller.Delete(request))
	}
}

func (handler *CompanyHandler) Products(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing %v Request to Company -> Products Controller\n", request.Method)
}

func (handler *CompanyHandler) Requests(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing %v Request to Company -> Requests Controller\n", request.Method)
}

func (handler *CompanyHandler) Integrations(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing %v Request to Company -> Integrations Controller\n", request.Method)
}

func writeResponse(writter http.ResponseWriter, result controller.HTTPResult) {
	if result.Error != nil {
		http.Error(writter, result.Error.Error(), result.Code)
	} else {
		writter.WriteHeader(result.Code)
		writter.Write(result.JSONData)
	}
}
