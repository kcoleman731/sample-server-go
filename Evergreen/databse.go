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

type Database struct {
	User       string
	Password   string
	Name       string
	Connection *sql.DB
}

func NewDatabase(user string, password, string, name string) (*Database, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, name)
	connection, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Printf("Failed connecting to database with error\n")
		return nil, err
	}
	return &Database{user, password, name, connection}, nil
}

func (d *Database) Execute(sql string, args ...string) (sql.Result, error) {
	return d.Connection.Exec(sql, args)
}

func (d *Database) Query(sql string, db *sql.DB, args string) (*sql.Rows, error) {
	return d.Connection.Query(sql, args)
}
