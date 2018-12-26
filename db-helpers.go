package main

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// connectionString -> user:pass@tcp(:3306)/test
func connectionString() string {
	var str string
	str = env["DB_USERNAME"] + ":" + env["DB_PASSWORD"] +
		"@(:" + env["DB_PORT"] + ")/" +
		env["DB_DATABASE"] + "?charset=utf8"
	return str
}

// getServerVersion ...
func getServerVersion(db *sql.DB) (string, error) {
	var serverVersion sql.NullString
	if err := db.QueryRow("SELECT version()").Scan(&serverVersion); err != nil {
		return "", err
	}
	return serverVersion.String, nil
}

// getTables ...
func getTables(db *sql.DB) ([]string, error) {
	tables := make([]string, 0)

	// Get table list
	rows, err := db.Query("SHOW TABLES")
	defer rows.Close()
	if err != nil {
		return tables, err
	}

	// Read result
	for rows.Next() {
		var table sql.NullString
		if err := rows.Scan(&table); err != nil {
			return tables, err
		}
		tables = append(tables, table.String)
	}
	return tables, rows.Err()
}

// createTableSQL ...
func createTableSQL(db *sql.DB, name string) (string, error) {
	// Get table creation SQL
	var tableReturn sql.NullString
	var tableSql sql.NullString
	err := db.QueryRow("SHOW CREATE TABLE "+name).Scan(&tableReturn, &tableSql)

	if err != nil {
		return "", err
	}
	if tableReturn.String != name {
		return "", errors.New("Returned table is not the same as requested table")
	}

	return tableSql.String, nil
}

// createTableValues
func createTableValues(db *sql.DB, name string) (string, error) {
	// Get Data
	rows, err := db.Query("SELECT * FROM " + name)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Get columns
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	if len(columns) == 0 {
		return "", errors.New("No columns in table " + name + ".")
	}

	// Read data
	dataText := make([]string, 0)
	for rows.Next() {
		// Init temp data storage
		data := make([]*sql.NullString, len(columns))
		ptrs := make([]interface{}, len(columns))
		for i := range data {
			ptrs[i] = &data[i]
		}

		// Read data
		if err := rows.Scan(ptrs...); err != nil {
			return "", err
		}

		dataStrings := make([]string, len(columns))

		for key, value := range data {
			if value != nil && value.Valid {
				dataStrings[key] = "'" + value.String + "'"
			} else {
				dataStrings[key] = "null"
			}
		}

		dataText = append(dataText, "("+strings.Join(dataStrings, ",")+")")
	}

	return strings.Join(dataText, ", "), rows.Err()
}

func createTable(db *sql.DB, name string) (*Table, error) {
	var err error
	t := &Table{Name: name}

	if t.SQL, err = createTableSQL(db, name); err != nil {
		return nil, err
	}

	if t.Values, err = createTableValues(db, name); err != nil {
		return nil, err
	}

	return t, nil
}
