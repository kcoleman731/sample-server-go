package controller

import (
	"database/sql"
	"net/http"
)

type IntegrationController struct {
	DB *sql.DB
}

func (controller *IntegrationController) Post(request *http.Request) {

}

func (controller *IntegrationController) Get(request *http.Request) {

}

func (controller *IntegrationController) Update(request *http.Request) {

}

func (controller *IntegrationController) Delete(request *http.Request) {

}
