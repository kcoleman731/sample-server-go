package model

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/kcoleman731/evergreen"
)

const CompanyResource string = "CompanyResource"

type Company struct {
	Model
	ModelInterface
	Name    string `json:"name"`
	Funding string `json:"funding"`
	Website string `json:"website"`
	Created string `json:"created"`
}

func (c *Company) Create() error {
	query := evergreen.Query.Insert(c.tableName).Collums(c.DBCollums())
	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	rowIdentifier, err := evergreen.DatabaseIdentifier(rows)
	if err != nil {
		return err
	}
	c.Identifier = rowIdentifier
	return err
}

func (c *Company) Find(values map[string]interface{}, limit int) error {
	query := evergreen.Query.Select("*").From(c.tableName).Where("company_id").Values(1)
	rows, err := database.Query(query)
	if err != nil {
		return err
	}
	*company = CompanyFromQuery(rows)
	return err
}

func (c *Company) Update(u map[string]interface{}) error {
	query := evergreen.Query.Insert(c.tableName).Collums(Keys(u)).Values(Values(u)).Return("company_id")
	rows, err := database.Query(query)
	if err != nil {
		err
	}
	return err
}

func (c *Company) Delete() error {
	query := evergreen.Query.Delete(c.tableName).Collmns("database_identifier").Values("values")
	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	return err
}

//-----------------
// Database Helpers
//-----------------

func (c *Company) DBCollums() []string {
	// Capture the type of struct we are.
	modelStruct := reflect.TypeOf(c)
	// count := modelStruct.NumField()

	// Itterate through all feilds to get DB tags.
	sqlCollums := []string{}
	for i := 0; i < modelStruct.NumField(); i++ {
		field := modelStruct.Field(i)
		tagValue := field.Tag.Get(modelTag)
		if tagValue != "" {
			sqlCollums = append(sqlCollums, tagValue)
		}
	}
	return sqlCollums
}

func (c *Company) DBValues([]string) []interface{} {
	// Capture the type of struct we are.
	modelStruct := reflect.TypeOf(*c)
	sqlValues := []interface{}{}

	// Itterate through all feilds to get DB tags.
	for i := 0; i < modelStruct.NumField(); i++ {
		value := reflect.ValueOf(*c).Field(i)
		if &value != nil {
			sqlValues = append(sqlValues, value)
		}
	}
	return sqlValues
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

func Keys(m map[string]interface{}) []string {
	keys := []string{}
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}

func Values(m map[string]interface{}) []string {
	values := []interface{}{}
	for _, v := range mymap {
		values = append(values, v)
	}
	return values
}
