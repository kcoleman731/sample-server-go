package model

import (
	"reflect"

	"github.com/kcoleman731/evergreen"
)

const modelTag = "model"

type Model struct {
	ID        int64  `model:"identitifer"`
	Type      string `model:"type"`
	TableName string `model:"table_name"`
	db        evergreen.Database
}

type ModelInterface interface {

	// Creates and persists all attributes of a model to the underlying database.
	Create()

	// Delete pruges the record of the model and its attributes from the
	// underlying database.
	Delete()

	// Update persists all values contained in the `value` parameter to the
	// underlying database for the given model. If `value` is empty, the
	// all model attributes will be updated.
	//
	// `Values` - A map of values to be updated for the model.
	Update(values map[string]interface{})

	// Find queries the underlying database for model records whose
	// attributes match those supplies in the `values` parameter.
	//
	// `Values` - A map of values for which a query will be performed.
	//
	// `Limit` - Used to specify the limit of records returned.
	Find(params map[string]interface{}, limit int)

	// Description prints a simple, human readable description of the model.
	Description()

	// Hash prints a human readable hash representation of the model.
	Hash()

	// SQLCollums returns an array containing all of the database collumns
	// for the model.
	DBCollums()

	// SQLCollums returns an array containing all of the database collumns
	// for the model.
	DBValues([]string)
}

func DBCollums(i interface{}) []string {
	// Capture the type of struct we are.
	modelStruct := reflect.TypeOf(i)
	// count := modelStruct.NumField()

	// Itterate through all feilds to get DB tags.
	sqlCollums := []string{}
	for c := 0; c < modelStruct.NumField(); c++ {
		field := modelStruct.Field(c)
		tagValue := field.Tag.Get(modelTag)
		if tagValue != "" {
			sqlCollums = append(sqlCollums, tagValue)
		}
	}
	return sqlCollums
}

func DBValues(i interface{}) []interface{} {
	// Capture the type of struct we are.
	modelStruct := reflect.TypeOf(i)
	sqlValues := []interface{}{}

	// Itterate through all feilds to get DB tags.
	for v := 0; v < modelStruct.NumField(); v++ {
		value := reflect.ValueOf(i).Field(v)
		if &value != nil {
			sqlValues = append(sqlValues, value)
		}
	}
	return sqlValues
}
