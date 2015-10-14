package model

import (
	"database/sql"
	"fmt"
	"reflect"
)

const modelTag = "model"

type Model struct {
	Identifier string `model:"identitifer"`
	Type       string
	tableName  string
	db         *sql.DB
}

type ModelStore interface {
	// Create persists all attributes of a model to the underlying database.
	Create()

	// Delete pruges the record of the model and its attributes from the
	// underlying database.
	Delete()

	// Description prints a simple, human readable description of the model.
	Description()

	// Find queries the underlying database for model records whose
	// attributes match those supplies in the `values` parameter.
	//
	// `Values` - A map of values for which a query will be performed.
	//
	// `Limit` - Used to specify the limit of records returned.
	Find(values map[string]interface{}, limit int)

	// Hash prints a human readable hash representation of the model.
	Hash()

	// Update persists all values contained in the `value` parameter to the
	// underlying database for the given model. If `value` is empty, the
	// all model attributes will be updated.
	//
	// `Values` - A map of values to be updated for the model.
	Update(values map[string]interface{})

	// SQLCollums returns an array containing all of the database collumns
	// for the model.
	//
	//
	dbCollums()

	dbValues([]string)
}

func (m *Model) Create() error {
	collums := m.dbCollums()
	values := m.dbValues(collums)

	sql := Insert(m.tableName).Collums(collums).Values(values)
	rows, err := database.Query(sql, values...)
	if err != nil {
		return err
	}
	rowIdentifier, err := DatabaseIdentifier(rows)
	if err != nil {
		return err
	}
	m.Identifier = rowIdentifier
	return err
}

func Find(values map[string]interface{}, limit int) {
	sql := "SELECT * FROM companies WHERE company_id = $1"
	fmt.Printf("Querying for company by id %v\n", identifer)
	rows, err := database.Query(sql, identifer)
	if err != nil {
		fmt.Printf("Failed finding company with error  - %+v\n", err)
	}
	*company = CompanyFromQuery(rows)
	return err
}

func (m *Model) dbCollums() (string, string) {
	// Capture the type of struct we are.
	modelStruct := reflect.TypeOf(m)
	sqlCollums = make([]string)

	// Itterate through all feilds to get DB tags.
	for i := 0; i < modelStruct.NumFeild(); i++ {
		structFeild = object.Field(i)
		tagValue = structFeild.Tag.Get(modelTag)
		if tagValue != "" {
			sqlCollmns.append(tagValue)
		}
	}
	return sqlCollmsn
}

func (m *Model) dbValues() (string, string) {
	// Capture the type of struct we are.
	modelStruct := reflect.TypeOf(m)
	sqlCollums = make([]string)

	// Itterate through all feilds to get DB tags.
	for i := 0; i < modelStruct.NumFeild(); i++ {
		structFeild = object.Field(i)
		tagValue = structFeild.Tag.Get(modelTag)
		if tagValue != "" {
			sqlCollmns.append(tagValue)
		}
	}
	return sqlCollmsn
}

//------------
// SQL Helpers
//------------

func Select(values []string) string {
	if values.Count > 0 {
		return fmt.Sprintf("SELECT * ")
	} else {
		return fmt.Sprintf("SELECT * ")
	}
}

func From(value string) string {
	return fmt.Sprintf("FROM %v", value)
}

func Insert(table string) string {
	return fmt.Sprintf("INSERT INTO %v", m.Name)
}

func Collums(values []string) string {
	return fmt.Sprintf("(%v)", collums)
}

func Values(values []string) string {
	return fmt.Sprintf("VALUES(%v)", values)
}

func Return(value string) {
	return fmt.Sprintf("RETURNING %v", value)
}
