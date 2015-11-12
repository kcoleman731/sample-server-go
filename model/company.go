package model

import (
	"database/sql"
	"fmt"

	"github.com/kcoleman731/evergreen"
)

const CompanyResource string = "CompanyResource"

type Company struct {
	Model
	Name    string `json:"name"`
	Funding string `json:"funding"`
	Website string `json:"website"`
}

func (c *Company) Create() error {
	query := &evergreen.Query{
		Action:  evergreen.INSERT,
		Table:   c.TableName,
		Collums: DBCollums(c),
		Values:  DBValues(c),
	}
	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}

	rowIdentifier, err := evergreen.DatabaseIdentifier(rows)
	if err != nil {
		return err
	}
	c.ID = rowIdentifier.(int64)
	return err
}

func (c *Company) Find(params map[string]interface{}, limit int) error {
	query := &evergreen.Query{
		Action: evergreen.SELECT,
		Table:  c.TableName,
		Where:  params,
	}
	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	*c = ObjectsFromRows(rows)
	return err
}

func (c *Company) Update(params map[string]interface{}) error {
	query := &evergreen.Query{
		Action: evergreen.UPDATE,
		Table:  c.TableName,
		Where:  params,
	}
	_, err := c.db.Query(query)
	if err != nil {
		return err
	}
	return err
}

func (c *Company) Delete() error {
	query := &evergreen.Query{
		Action: evergreen.DELETE,
		Table:  c.TableName,
		Where: map[string]interface{}{
			"database_identifier": c.ID,
		},
	}
	_, err := c.db.Query(query)
	if err != nil {
		return err
	}
	return err
}

//-----------------
// Database Helpers
//-----------------

func ObjectsFromRows(rows *sql.Rows) Company {
	company := Company{}
	if rows.Next() {
		err := rows.Scan(&company.ID, &company.Name, &company.Funding, &company.Website)
		if err != nil {
			fmt.Printf("Failed getting database identifier with error - %+v\n", err)
		}
	}
	return company
}
