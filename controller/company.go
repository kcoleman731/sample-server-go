package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kcoleman731/test-server/model"
)

type CompanyController struct {
	DB     *sql.DB
	Params map[string]string
}

// Performs the required actions form a GET HTTP request.
// Method with search for a company based a comapny_id
// parameter.
func (controller *CompanyController) Get(request *http.Request) HTTPResult {
	fmt.Printf("GET Request with params %v\n", controller.Params)

	// Query for company based on company_id
	company := &model.Company{}
	err := company.Find(controller.DB, controller.Params["company_id"])
	if err != nil {
		fmt.Printf("Failed to query for company with error: %v", err)
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}

	// If we have a nil company, the object doesn not exist.
	if company.Identifier == "" {
		errorString := fmt.Sprintf("Failed to find company with id %s", controller.Params["company_id"])
		err = errors.New(errorString)
		return HTTPResult{err, http.StatusNotFound, nil}
	}

	// We have a valid company, serialize and return
	fmt.Printf("Found company, serializing data %v\n", company)
	jsonData, err := json.Marshal(company)
	if err != nil {
		fmt.Printf("Failed to serialize company data with error: %v", err)
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}

	return HTTPResult{nil, http.StatusOK, jsonData}
}

func (controller *CompanyController) Post(request *http.Request) HTTPResult {
	// Parse Company JSON
	company, err := parseCompanyJSON(request)
	if err != nil {
		fmt.Printf("Failed to parse company JSON with error: %v", err)
		return HTTPResult{err, http.StatusBadRequest, nil}
	}

	// Persist Company Data
	fmt.Printf("Persisting company data: %v\n", company)
	err = company.Save(controller.DB)
	if err != nil {
		fmt.Printf("Failed to persist company data with error: %v", err)
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}

	// Serialize Company Data
	jsonData, err := json.Marshal(company)
	if err != nil {
		fmt.Printf("Failed to serialize company data with error: %v", err)
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}

	return HTTPResult{nil, http.StatusCreated, jsonData}
}

func (controller *CompanyController) Update(request *http.Request) HTTPResult {
	// Parse Company JSON
	company, err := parseCompanyJSON(request)
	if err != nil {
		fmt.Printf("Failed to parse company JSON with error: %v", err)
		return HTTPResult{err, http.StatusBadRequest, nil}
	}

	// Persist Company Data
	fmt.Printf("Persisting company update data: %v\n", company)
	err = company.Update(controller.DB)
	if err != nil {
		fmt.Printf("Failed to persist company update with error: %v", err)
		return HTTPResult{err, http.StatusNotFound, nil}
	}

	// Serialize Company Data
	jsonData, err := json.Marshal(company)
	if err != nil {
		fmt.Printf("Failed to serialize company data with error: %v", err)
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}
	return HTTPResult{nil, http.StatusOK, jsonData}

}

func (controller *CompanyController) Delete(request *http.Request) HTTPResult {
	return HTTPResult{nil, http.StatusInternalServerError, nil}
}

func parseCompanyJSON(r *http.Request) (*model.Company, error) {
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
