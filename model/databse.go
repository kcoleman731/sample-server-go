package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	databaseUser     = "kevincoleman"
	databasePassword = ""
	databaseName     = "kevincoleman"
)

type Taste struct {
}

func NewDatabase() *sql.DB {
	fmt.Printf("Opening Database Connection...\n")

	// Open a Database Connection for the givein credentials
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", databaseUser, databasePassword, databaseName)
	database, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	fmt.Printf("Connected to database...\n")

	// Open a Database Connection for the givein credentials
	//---------------------------------------
	// TODO: Create database tables if needed
	//---------------------------------------

	return database
}

func Execute(sql string, db *sql.DB, args ...string) (sql.Result, error) {
	return db.Exec(sql, args)
}

func Query(sql string, db *sql.DB, args string) (*sql.Rows, error) {
	return db.Query(sql, args)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
