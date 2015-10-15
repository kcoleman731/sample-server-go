package model

import "database/sql"

const modelTag = "model"

type Model struct {
	Identifier string `model:"identitifer"`
	Type       string `model:"type"`
	tableName  string `model:"table_name"`
	db         *sql.DB
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
	Find(values map[string]interface{}, limit int)

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
