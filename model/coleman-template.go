package main

var header string = `// Code generated by ModelQ
// {{.TableName}}.go contains model for the database table [{{.DbName}}.{{.TableName}}]

package {{.PkgName}}

import (
	"encoding/json"
	"encoding/gob"
	"fmt"
	"strings"
	"github.com/mijia/modelq/gmq"
	"database/sql"
	{{if .ImportTime}}"time"{{end}}
)
`
var modelStruct string = `type {{.Name}} struct {
	{{range .Fields}}{{.Name}} {{.Type}} {{.JsonMeta}}{{if .Comment}} // {{.Comment}}{{end}}
	{{end}}
}
`

var objApi string = `
func (obj *{{.Name}}) Create() ({{.Name}}, error) {
	{{if .PrimaryFields}}{{ call .PrimaryFields.FormatFilters .Name }}
	if result, err := {{.Name}}Objs.Select().Where(filter).One(dbtx); err != nil {
		return obj, err
	} else {
		return result, nil
	}{{else}}return 0, gmq.ErrNoPrimaryKeyDefined{{end}}
}

func (obj *{{.Name}}) Create() error {
	query := &evergreen.Query{
		Action:  evergreen.INSERT,
		Table:   obj.TableName,
		Collums: DBCollums(obj),
		Values:  DBValues(obj),
	}
	rows, err := obj.db.Query(query)
	if err != nil {
		return err
	}

	identifier, err := evergreen.DatabaseIdentifier(rows)
	if err != nil {
		return err
	}
	obj.ID = rowIdentifier.(int64)
	return err
}

func (obj *{{.Name}}) Find(params map[string]interface{}, limit int) error {
	query := &evergreen.Query{
		Action: evergreen.SELECT,
		Table:  obj.TableName,
		Where:  params,
	}
	rows, err := obj.db.Query(query)
	if err != nil {
		return err
	}
	*obj = ObjectsFromRows(rows)
	return err
}

func (obj *{{.Name}}) Update(params map[string]interface{}) error {
	query := &evergreen.Query{
		Action: evergreen.UPDATE,
		Table:  obj.TableName,
		Where:  params,
	}
	_, err := obj.db.Query(query)
	if err != nil {
		return err
	}
	return err
}

func (obj *{{.Name}}) Delete() error {
	query := &evergreen.Query{
		Action: evergreen.DELETE,
		Table:  obj.TableName,
		Where: map[string]interface{}{
			"database_identifier": obj.ID,
		},
	}
	_, err := obj.db.Query(query)
	if err != nil {
		return err
	}
	return err
}

func ObjectsFromRows(rows *sql.Rows) {{.Name}} {
	object := {{.Name}}{}
	if rows.Next() {
		err := rows.Scan(&company.ID, &company.Name, &company.Funding, &company.Website)
		if err != nil {
			fmt.Printf("Failed getting database identifier with error - %+v\n", err)
		}
	}
	return company
}`
