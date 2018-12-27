package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Insert insert file to DB
func Insert(filePath string) {
	// Open connection to database
	db, err := sql.Open("mysql", connectionString())
	checkErr(err, "Connection error")
	defer db.Close()

	file, err := ioutil.ReadFile(filePath)
	checkErr(err, "Can't read file: "+filePath)

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		request = strings.Trim(request, "\n")
		request = strings.Trim(request, " ")
		if request != "" {
			// result, err := db.Exec(request)
			_, err := db.Exec(request)
			checkErr(err, "Can't exec SQL \n\n"+request+"\n")
		}
	}

	fmt.Println("Executed file: " + filePath)
}
