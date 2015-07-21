package main

import (
	"github.com/gorilla/mux"
	"github.com/kcoleman731/test-server/handlers"
)

func RouteRequest() *mux.Router {

	// Setup out router
	r := mux.NewRouter()

	// Setup Handlers
	companyHandler := handlers.CompanyHandler{DB: Database}
	productHandler := handlers.ProductHandler{DB: Database}
	featureHandler := handlers.FeatureHandler{DB: Database}
	userHandler := handlers.UserHandler{DB: Database}
	requestHandler := handlers.RequestHandler{DB: Database}
	integrationHandler := handlers.IntegrationHandler{DB: Database}

	//---------------
	// Company Routes
	//---------------

	r.HandleFunc("/companies", companyHandler.Index)
	r.HandleFunc("/companies/{company_id}", companyHandler.Index)

	r.HandleFunc("/companies/{company_id}/products", companyHandler.Products)
	r.HandleFunc("/companies/{company_id}/requests", companyHandler.Requests)
	r.HandleFunc("/companies/{company_id}/integrations", companyHandler.Integrations)

	//---------------
	// Product Routes
	//---------------

	r.HandleFunc("/products", productHandler.Index)
	r.HandleFunc("/products/{product_id}", productHandler.Index)

	r.HandleFunc("/products/{product_id}/features", productHandler.Index)
	r.HandleFunc("/products/{product_id}/requests", productHandler.Index)

	//---------------
	// Feature Routes
	//---------------

	r.HandleFunc("/features", featureHandler.Index)
	r.HandleFunc("/features/{feature_id}", featureHandler.Index)

	r.HandleFunc("/features/{feature_id}/requests", featureHandler.Index)

	//---------------
	// User Routes
	//---------------

	r.HandleFunc("/user", userHandler.Index)
	r.HandleFunc("/user/{user_id}", userHandler.Index)

	//---------------
	// Request Routes
	//---------------

	r.HandleFunc("/requests", requestHandler.Index)
	r.HandleFunc("/request/{request_id}", requestHandler.Index)

	//-------------------
	// Integration Routes
	//-------------------

	r.HandleFunc("/integrations", integrationHandler.Index)
	r.HandleFunc("/integrations/{integration_id}", integrationHandler.Index)

	return r
	//http.Handle("/", r)
}
