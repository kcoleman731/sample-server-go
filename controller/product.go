package controller

import (
	"database/sql"
	"net/http"
)

type ProductController struct {
	DB *sql.DB
}

func (controller *ProductController) Post(request *http.Request) {

}

func (controller *ProductController) Get(request *http.Request) {

}

func (controller *ProductController) Update(request *http.Request) {

}

func (controller *ProductController) Delete(request *http.Request) {

}
