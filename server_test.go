package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kcoleman731/sample-server-go/controller"

	//"github.com/kcoleman731/test-server/controller"
)

var (
	server       *httptest.Server
	reader       io.Reader //Ignore this for now
	companiesURL string
)

func init() {
	server = httptest.NewServer(RouteRequest()) //Creating new server with the server handlers
}

//---------------
// POST Companies
//---------------

func TestCreateCompany(t *testing.T) {
	companiesURL = fmt.Sprintf("%s/companies", server.URL)
	companyJSON := `{"name": "test company", "funding" : "test funding", "website" : "www.test.com"}`

	reader = strings.NewReader(companyJSON)                       //Convert string to reader
	request, err := http.NewRequest("POST", companiesURL, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 201 {
		t.Errorf("Expected 201 status code but got: %d, %v", res.StatusCode, res.Body)
	}
}

func TestCreateCompanyWithNoParams(t *testing.T) {
	companiesURL = fmt.Sprintf("%s/companies", server.URL)
	request, err := http.NewRequest("POST", companiesURL, nil) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 400 {
		t.Errorf("Expected 400 status code but got: %d", res.StatusCode)
	}
}

//--------------
// GET Companies
//--------------

func TestGetCompany(t *testing.T) {
	companyID := createTestCompany(t)
	companiesURL = fmt.Sprintf("%s/companies/%v", server.URL, companyID)
	request, err := http.NewRequest("GET", companiesURL, nil)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected 200 status code but got: %d", res.StatusCode)
	}
}

func TestGetCompanyWithNoParams(t *testing.T) {
	companiesURL = fmt.Sprintf("%s/companies/500", server.URL)
	request, err := http.NewRequest("GET", companiesURL, nil)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 404 {
		t.Errorf("Expected 404 status code but got: %d", res.StatusCode)
	}
}

func TestGetCompanyCompanyThatDoesNotExist(t *testing.T) {
	companiesURL = fmt.Sprintf("%s/companies/500", server.URL)
	request, err := http.NewRequest("GET", companiesURL, nil)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 404 {
		t.Errorf("Expected 404 status code but got: %d", res.StatusCode)
	}
}

//---------------
// POST Companies
//---------------

func TestUpdateCompanyFunding(t *testing.T) {
	companyID := createTestCompany(t)
	companyJSON := `{"name": "test company", "funding" : "new test funding", "website" : "www.test.com"}`

	companiesURL = fmt.Sprintf("%s/companies/%v", server.URL, companyID)
	reader = strings.NewReader(companyJSON)
	request, err := http.NewRequest("PUT", companiesURL, reader)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	JSON, err := controller.ParseJSON(res.Body)
	if err != nil {
		t.Errorf("Failed to parse response JSON with error: %v", err)
	}

	companyFunding := JSON["funding"]
	if companyFunding != "new test funding" {
		t.Errorf("Expected Funding to be new test funding but got %v: %d", companyFunding, res.StatusCode)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 status code but got: %d", res.StatusCode)
	}
}

func TestUpdateCompanyWithNoParams(t *testing.T) {
	companyID := createTestCompany(t)
	companiesURL = fmt.Sprintf("%s/companies/%v", server.URL, companyID)
	request, err := http.NewRequest("PUT", companiesURL, nil)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode != 400 {
		t.Errorf("Expected 400 status code but got: %d", res.StatusCode)
	}
}

func TestUpdateCompanyThatDoesntExist(t *testing.T) {
	companyJSON := `{"name": "test company", "funding" : "new test funding", "website" : "www.test.com"}`

	companiesURL = fmt.Sprintf("%s/companies/%v", server.URL, 500)
	reader = strings.NewReader(companyJSON)
	request, err := http.NewRequest("PUT", companiesURL, reader)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 404 {
		t.Errorf("Expected 404 status code but got: %d", res.StatusCode)
	}
}

//----------------
// Factory Methods
//----------------

func createTestCompany(t *testing.T) string {
	companiesURL = fmt.Sprintf("%s/companies", server.URL)
	companyJSON := `{"name": "test company", "funding" : "test funding", "website" : "www.test.com"}`

	reader = strings.NewReader(companyJSON)                       //Convert string to reader
	request, err := http.NewRequest("POST", companiesURL, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 201 {
		t.Errorf("Expected 201 status code but got: %d", res.StatusCode)
	}
	JSON, err := controller.ParseJSON(res.Body)
	if err != nil {
		t.Errorf("Failed to parse response JSON with error: %v", err)
	}
	companyID := JSON["company_id"]
	return companyID.(string)
}
