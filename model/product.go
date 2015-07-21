package model

import (
	"database/sql"
	"fmt"
)

const ProductResource string = "ProductResource"

type Product struct {
	Name        string
	Description string
}

func SaveProduct(product Product, db *sql.DB) (sql.Result, error) {
	err := db.Ping()
	result, err := db.Exec("INSERT INTO product(name, description) VALUES($1,$2);", product.Name, product.Description)
	if err != nil {
		fmt.Printf("This is my Error - %+v\n", err)
	} else {
		fmt.Printf("Product Saved!")
	}
	return result, err
}
