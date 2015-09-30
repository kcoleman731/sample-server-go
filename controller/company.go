package controller

import (
	"encoding/json"
	"net/http"

	"github.com/kcoleman731/test-server/model"
)

// CompanyController provides for processing HTTP requests made to the /companies
// API.
type CompanyController struct {
	Controller
	Company model.Company
}

// Post processes HTTP POST requests made to the /companies API.
func (c *CompanyController) Post() HTTPResult {
	// Persist Company Data
	err := c.Company.Save(c.DB)
	if err != nil {
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}

	// Serialize Company Data
	jsonData, err := json.Marshal(c.Company)
	if err != nil {
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}

	return HTTPResult{nil, http.StatusCreated, jsonData}
}

// Get processes HTTP GET requests made to the /companies API.
func (c *CompanyController) Get() HTTPResult {
	company := &model.Company{}

	// Query for company based on company_id
	companyID := c.Request.Params["id"].(string)
	err := company.Find(c.DB, companyID)
	if err != nil {
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}

	// If we have a nil company, the object doesn not exist.
	if company.Identifier == "" {
		return HTTPResult{err, http.StatusNotFound, nil}
	}

	// We have a valid company, serialize and return
	jsonData, err := json.Marshal(company)
	if err != nil {
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}

	return HTTPResult{nil, http.StatusOK, jsonData}
}

// Update processes HTTP PUT requests made to the /companies API.
func (c *CompanyController) Update() HTTPResult {
	// Persist Company Data
	err := c.Company.Update(c.DB)
	if err != nil {
		return HTTPResult{err, http.StatusNotFound, nil}
	}

	// Serialize Company Data
	jsonData, err := json.Marshal(c.Company)
	if err != nil {
		return HTTPResult{err, http.StatusInternalServerError, nil}
	}
	return HTTPResult{nil, http.StatusOK, jsonData}

}

// Delete processes HTTP DELETE requests made to the /companies API.
func (c *CompanyController) Delete() HTTPResult {
	return HTTPResult{nil, http.StatusInternalServerError, nil}
}
