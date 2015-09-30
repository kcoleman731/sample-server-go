package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/kcoleman731/test-server/controller"
)

type CompanyHandler struct {
	DB *sql.DB
}

func (handler *CompanyHandler) Index(writter http.ResponseWriter, request *http.Request) {
	companyController, err := NewCompanyController(request, handler.DB)
	if err != nil {

	}
	switch request.Method {
	case "GET":
		writeResponse(writter, companyController.Get())

	case "POST":
		writeResponse(writter, companyController.Post())

	case "PUT":
		writeResponse(writter, companyController.Update())

	case "DELETE":
		writeResponse(writter, companyController.Delete())
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

func NewCompanyController(request *http.Request, DB *sql.DB) (controller.CompanyController, error) {
	genericController, err := controller.NewController(request, DB)
	if err != nil {

	}
	return genericController.(controller.CompanyController), err
}
