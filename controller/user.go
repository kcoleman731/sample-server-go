package controller

import (
	"database/sql"
	"net/http"
)

type UserController struct {
	DB *sql.DB
}

func (controller *UserController) Post(request *http.Request) {

}

func (controller *UserController) Get(request *http.Request) {

}

func (controller *UserController) Update(request *http.Request) {

}

func (controller *UserController) Delete(request *http.Request) {

}
