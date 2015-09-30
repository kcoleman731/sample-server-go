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
		return HTTPResult{http.StatusInternalServerError, nil, err}
	}

	// Serialize Company Data
	companyData, err := json.Marshal(c.Company)
	if err != nil {
		return HTTPResult{http.StatusInternalServerError, nil, err}
	}

	return HTTPResult{http.StatusCreated, companyData, err}
}

// Get processes HTTP GET requests made to the /companies API.
func (c *CompanyController) Get() HTTPResult {
	companyID := c.Request.Body["id"].(string) // Add model method to get JSON tag from struct attribute

	// Query for company based on company_id
	company := &model.Company{}
	err := company.Find(c.DB, companyID)
	if err != nil {
		return HTTPResult{http.StatusInternalServerError, nil, err}
	}

	// If we have a nil company, the object doesn not exist.
	if company.Identifier == "" {
		return HTTPResult{http.StatusNotFound, nil, err}
	}

	// We have a valid company, serialize and return
	companyData, err := json.Marshal(company)
	if err != nil {
		return HTTPResult{http.StatusInternalServerError, nil, err}
	}

	return HTTPResult{http.StatusOK, companyData, nil}
}

// Update processes HTTP PUT requests made to the /companies API.
func (c *CompanyController) Update() HTTPResult {
	// Persist Company Data
	err := c.Company.Update(c.DB)
	if err != nil {
		return HTTPResult{http.StatusNotFound, nil, err}
	}

	// Serialize Company Data
	companyData, err := json.Marshal(c.Company)
	if err != nil {
		return HTTPResult{http.StatusInternalServerError, nil, err}
	}
	return HTTPResult{http.StatusOK, companyData, nil}
}

// Delete processes HTTP DELETE requests made to the /companies API.
func (c *CompanyController) Delete() HTTPResult {
	// Persist Company Data
	err := c.Company.Delete(c.DB)
	if err != nil {
		return HTTPResult{http.StatusNotFound, nil, err}
	}
	return HTTPResult{http.StatusOK, nil, nil}
}
