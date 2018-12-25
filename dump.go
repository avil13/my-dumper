package main

import (
	"database/sql"
	"fmt"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// Table ...
type Table struct {
	Name   string
	SQL    string
	Values string
}

// Dump ...
type Dump struct {
	Tables []*Table
}

const createTableTemplate = `
{{range .Tables}}
--
-- Table structure for table {{ .Name }}
--
{{.SQL}}
{{end}}
`

var reportCreate = template.Must(template.New("create").Parse(createTableTemplate))

// Funcs(template.FuncMap{"dayAgo": daysAgo}).

// GetDump - make dump
func GetDump() {
	// Open connection to database
	db, err := sql.Open("mysql", connectionString())
	checkErr(err, "Connection error")
	defer db.Close()

	tables, err := getTables(db)
	checkErr(err, "Not found tables")

	data := Dump{
		Tables: make([]*Table, 0),
	}

	for _, tableName := range tables {
		table, err := createTable(db, tableName)
		checkErr(err, "Can't create `"+tableName+"`")
		data.Tables = append(data.Tables, table)
	}

	fmt.Println(data.Tables[0].SQL)
}

// GetCreateSqlTables
// func GetCreateSqlTables(data *Dump) (string, error) {
// 	reportCreate.
// }
