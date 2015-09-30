package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcoleman731/test-server/model"
)

// ----------------------------------------------------------------------------
// Controller
// ----------------------------------------------------------------------------

// Controller models an HTTP controller.
type Controller struct {
	// Request models the HTTP request submitted to a controller.
	Request HTTPRequest
	// DB models a database connection to the underlying database
	// for teh controller.
	DB *sql.DB
}

// Controller interface models an HTTP request controller
// for a specific object model. The controller handles
// requests dispatched from a Handler, performs crud operations,
// and returns the result of the operation to the Handler.
type ControllerInterface interface {
	// Post notifies the controller that an HTTP POST request has
	// been made to its model's resource API.
	Post()
	// Get notifies the controller that an HTTP GET request has
	// been made to its model's resource API.
	Get()
	// PUT notifies the controller that an HTTP PUT request has
	// been made to its model's resource API.
	Update()
	// DELTE notifies the controller that an HTTP DELETE request has
	// been made to its model's resource API.
	Delete()
}

func NewController(request *http.Request, DB *sql.DB) (interface{}, error) {
	body, err := ParseJSON(request.Body)
	if err != nil {
		return nil, err
	}
	params := mux.Vars(request)
	httpRequest := HTTPRequest{params, body}
	return &Controller{httpRequest, DB}, err
}

// ----------------------------------------------------------------------------
// HTTP Helper Structs
// ----------------------------------------------------------------------------

// HTTPResult models the result of an HTTP request after it has been
// processed by a controller.
type HTTPResult struct {

	// Code represents the response code for the HTTP request.
	Code int

	// JSON data models the response body for the HTTP request.
	Data []byte

	// Error represents any error that may have occured during processing.
	Error error
}

// HTTPRequest models an HTTP request sumbitted to the controller
// for processing.
type HTTPRequest struct {
	// URLParams contains all URL parameters submitted with the request.
	Params map[string]string
	// Params models the body of the request as a hashmap.
	Body map[string]interface{}
}

// ----------------------------------------------------------------------------
// JSON Helpers
// ----------------------------------------------------------------------------

func ParseJSON(rc io.Reader) (map[string]interface{}, error) {
	// Read the request body
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	// Parse data into a generic interface. This will get us our JSON!
	var JSON interface{}
	err = json.Unmarshal(body, &JSON)
	if err != nil {
		return nil, err
	}

	// Create a map from the unmarshalled interface
	JSONMap := JSON.(map[string]interface{}) // This is a type assertion
	return JSONMap, nil
}

func ParseModleJSON(r *http.Request) (*model.Company, error) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Parse Company JSON
	var company *model.Company
	err = json.Unmarshal(body, &company)
	if err != nil {
		return nil, err
	}
	// Return the company object
	return company, nil
}
