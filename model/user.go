package model

import (
	"database/sql"
	"fmt"
)

const UserResource string = "UserResource"

type User struct {
	FirstName string
	LastName  string
	Email     string
}

func SaveUser(user User, db *sql.DB) (sql.Result, error) {
	err := db.Ping()
	result, err := db.Exec("INSERT INTO companies(name, website) VALUES($1,$2);", user.FirstName, user.LastName)
	if err != nil {
		fmt.Printf("This is my Error - %+v\n", err)
	} else {
		fmt.Printf("Object Saved!")
	}
	return result, err
}
