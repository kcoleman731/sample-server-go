package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcoleman731/sample-server-go/controller"
)

type FeatureHandler struct {
	DB *sql.DB
}

func (handler *FeatureHandler) Index(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Routing Request to feature Controller\n")
	route := request.RequestURI
	fmt.Printf("Current Route %v\n", route)

	featureController := controller.FeatureController{DB: handler.DB}
	switch request.Method {
	case "GET":
		vars := mux.Vars(request)
		fmt.Printf("GET Request with var %v\n", vars)
		featureController.Get(request)

	case "POST":
		featureController.Post(request)

	case "UPDATE":
		featureController.Update(request)

	case "DELETE":
		featureController.Delete(request)
	}
}
