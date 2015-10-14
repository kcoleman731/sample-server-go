package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcoleman731/sample-server-go/controller"
)

type RequestHandler struct {
	DB *sql.DB
}

func (handler *RequestHandler) Index(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing Request to request Controller\n")
	route := request.RequestURI
	fmt.Printf("Current Route %v\n", route)

	requestController := controller.RequestController{DB: handler.DB}
	switch request.Method {
	case "GET":
		vars := mux.Vars(request)
		fmt.Printf("GET Request with var %v\n", vars)
		requestController.Get(request)

	case "POST":
		requestController.Post(request)

	case "UPDATE":
		requestController.Update(request)

	case "DELETE":
		requestController.Delete(request)
	}
}
