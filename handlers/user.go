package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcoleman731/test-server/controller"
)

type UserHandler struct {
	DB *sql.DB
}

func (handler *UserHandler) Index(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing Request to user Controller\n")
	route := request.RequestURI
	fmt.Printf("Current Route %v\n", route)

	userController := controller.UserController{DB: handler.DB}
	switch request.Method {
	case "GET":
		vars := mux.Vars(request)
		fmt.Printf("GET Request with var %v\n", vars)
		userController.Get(request)

	case "POST":
		userController.Post(request)

	case "UPDATE":
		userController.Update(request)

	case "DELETE":
		userController.Delete(request)
	}
}
