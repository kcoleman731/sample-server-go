package controller

import (
	"database/sql"
	"net/http"
)

type RequestController struct {
	DB *sql.DB
}

func (controller *RequestController) Post(request *http.Request) {

}

func (controller *RequestController) Get(request *http.Request) {

}

func (controller *RequestController) Update(request *http.Request) {

}

func (controller *RequestController) Delete(request *http.Request) {

}
