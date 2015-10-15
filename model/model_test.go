package model

import (
	"fmt"
	"os"
	"testing"
)

// Perform Setup Here
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestExtractingStructValues(t *testing.T) {
	company := Company{Name: "name", Funding: "taste"}
	company.Model = Model{Type: "Company"}
	values := company.DBValues(nil)
	fmt.Printf("Got Values %v", values)
}
