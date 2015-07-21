package model

import (
	"database/sql"
	"fmt"
)

const FeatureResource string = "FeatureResource"

type Feature struct {
	Name        string
	Description string
}

func SaveFeature(feature Feature, db *sql.DB) (sql.Result, error) {
	sql := "INSERT INTO features(name, description) VALUES($1,$2);"
	args := ""
	result, err := Execute(sql, db, args)
	if err != nil {
		fmt.Printf("This is my Error - %+v\n", err)
	} else {
		fmt.Printf("Feature Saved!")
	}
	return result, err
}
