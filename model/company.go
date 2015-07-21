package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

const CompanyResource string = "CompanyResource"

type Company struct {
	Identifier string      `json:"company_id"`
	Name       string      `json:"name"`
	Funding    string      `json:"funding"`
	Website    string      `json:"website"`
	Created    interface{} `json:"created"`
}

func (company *Company) Find(database *sql.DB, identifer string) error {
	sql := "SELECT * FROM companies WHERE company_id = $1"
	fmt.Printf("Querying for company by id %v\n", identifer)
	rows, err := database.Query(sql, identifer)
	if err != nil {
		fmt.Printf("Failed finding company with error  - %+v\n", err)
	}
	*company = CompanyFromQuery(rows)
	return err
}

func (company *Company) Save(database *sql.DB) error {
	sql := "INSERT INTO companies(name, website, funding) VALUES($1,$2,$3) RETURNING company_id;"
	rows, err := database.Query(sql, company.Name, company.Website, company.Funding)
	if err != nil {
		fmt.Printf("Failed persisting company with error  - %+v\n", err)
	}
	rowIdentifier, err := DatabaseIdentifier(rows)
	if err != nil {
		fmt.Printf("No Company ID compadre %v\n", rowIdentifier)
	}
	company.Identifier = rowIdentifier
	return err
}

func (company *Company) Update(database *sql.DB) error {
	companyID, err := strconv.Atoi(company.Identifier)
	sql := "UPDATE companies SET name = $1, website = $2, funding = $3 WHERE company_id = $4 RETURNING company_id;"
	rows, err := database.Query(sql, company.Name, company.Website, company.Funding, companyID)
	if err != nil {
		fmt.Printf("Failed persisting company update with error  - %+v\n", err)
	}

	rowIdentifier, err := DatabaseIdentifier(rows)
	if err != nil {
		fmt.Printf("No Company ID compadre %v\n", rowIdentifier)
	}
	return err
}

func (company *Company) Delete(database *sql.DB) error {
	sql := "INSERT INTO companies(name, website) VALUES($1,$2);"
	rows, err := database.Query(sql, company.Name, company.Website, company.Funding)
	if err != nil {
		fmt.Printf("Failed persisting company with error  - %+v\n", err)
	}
	rowIdentifier, err := DatabaseIdentifier(rows)
	if err != nil {
		fmt.Printf("No Company ID compadre %v\n", rowIdentifier)
	}
	return err
}

func DatabaseIdentifier(rows *sql.Rows) (string, error) {
	var err error
	var companyID string
	for rows.Next() {
		err = rows.Scan(&companyID)
		if err != nil {
			fmt.Printf("Failed getting database identifier with error - %+v\n", err)
		}
	}
	if companyID == "" {
		errorString := fmt.Sprintf("Failed to get Database identifier with error: %v\n", err)
		err = errors.New(errorString)
	}
	return companyID, err
}

func CompanyFromQuery(rows *sql.Rows) Company {
	company := Company{}
	if rows.Next() {
		err := rows.Scan(&company.Identifier, &company.Name, &company.Funding, &company.Website, &company.Created)
		if err != nil {
			fmt.Printf("Failed getting database identifier with error - %+v\n", err)
		}
	}
	// TODO - Fix this shit
	company.Created = "created"
	return company
}
