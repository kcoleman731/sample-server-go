package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcoleman731/sample-server-go/controller"
)

type ProductHandler struct {
	DB *sql.DB
}

func (handler *ProductHandler) Index(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing Request to Product Controller\n")
	route := request.RequestURI
	fmt.Printf("Current Route %v\n", route)

	productController := controller.ProductController{DB: handler.DB}
	switch request.Method {
	case "GET":
		vars := mux.Vars(request)
		fmt.Printf("GET Request with var %v\n", vars)
		productController.Get(request)

	case "POST":
		productController.Post(request)

	case "UPDATE":
		productController.Update(request)

	case "DELETE":
		productController.Delete(request)
	}
}
