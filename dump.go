package main

import (
	"database/sql"
	"log"
	"os"
	"path"
	"text/template"
	"time"

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
	DateTime string
	Tables   []*Table
}

const createTableTemplate = `
-- -----------------------
-- Date: {{ .DateTime }}
-- -----------------------

{{ range .Tables }}
--
-- Table structure for table "{{ .Name }}"
--
{{.SQL}}
{{end}}
`

const insertTableTemplate = `
-- -----------------------
-- Date: {{ .DateTime }}
-- -----------------------

{{ range .Tables }}
--
-- Dumping data for table {{ .Name }}
--
LOCK TABLES {{ .Name }} WRITE;
/*!40000 ALTER TABLE {{ .Name }} DISABLE KEYS */;
{{ if .Values }}
INSERT INTO {{ .Name }} VALUES {{ .Values }};
{{ end }}
/*!40000 ALTER TABLE {{ .Name }} ENABLE KEYS */;
UNLOCK TABLES;

{{ end }}

`

var reportCreate = template.Must(template.New("create").Parse(createTableTemplate))
var reportInsert = template.Must(template.New("insert").Parse(insertTableTemplate))

// Funcs(template.FuncMap{"dayAgo": daysAgo}).

// GetDump - make dump
func GetDump() *Dump {
	// Open connection to database
	db, err := sql.Open("mysql", connectionString())
	checkErr(err, "Connection error")
	defer db.Close()

	tables, err := getTables(db)
	checkErr(err, "Not found tables")

	data := Dump{
		DateTime: time.Now().Format("2006-01-02 15:04"),
		Tables:   make([]*Table, 0),
	}

	for _, tableName := range tables {
		table, err := createTable(db, tableName)

		checkErr(err, "Can't create `"+tableName+"`")

		data.Tables = append(data.Tables, table)
	}

	// fmt.Println(data.Tables[0].SQL)
	return &data
}

// MakeDumpFiles ...
func MakeDumpFiles(data *Dump, isInsert bool) {
	// file
	var fileName string
	if isInsert {
		fileName = "insert.sql"
	} else {
		fileName = "create.sql"
	}

	p := path.Join(env["DUMP_SUB_DIR"], fileName)

	file, err := os.Create(p)
	checkErr(err, "Can't crate file "+fileName)
	defer file.Close()

	// template
	var report *template.Template

	if isInsert {
		report = reportInsert
	} else {
		report = reportCreate
	}

	if err := report.Execute(file, data); err != nil {
		log.Fatal(err)
	}
}
