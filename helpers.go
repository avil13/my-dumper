package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// checkErr - checking error
func checkErr(err error, message string) {
	if err != nil {
		color.Set(color.FgHiRed)
		fmt.Println(message)
		log.Fatal(err)
		// panic(err)
		color.Unset()
	}
}

// checkParams - check func params
func checkParams(data *map[string]string, params []string) {
	var errList string

	for _, v := range params {
		if _, ok := (*data)[v]; ok != true {
			errList += "Invalid env parameter '" + v + "'\n"
		}
	}
	if errList != "" {
		envLoadErrorMessage(errList)
	}
}

func envLoadErrorMessage(msg string) {
	log.Fatalf(
		"We have a few problems with 'env' variables \n"+
			"To create the '.env' file, specify the option --make-env\n\n"+
			"%s", msg)
}

// dumpDir - make dump folder
func dumpDir(dir string) (err error) {
	if dir == "" {
		log.Fatalf("Invalid .env parameter %s ", dir)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	return
}

// makeEnv - create .env file
func makeEnv() {
	tmpl := `# Examle .env vars
DB_HOST=mysql
DB_PORT=3306
DB_DATABASE=homestead
DB_USERNAME=homestead
DB_PASSWORD=secret

DUMP_DIR=dumps

DEBUG=false
	`
	file, err := os.Create(".env")
	checkErr(err, "Can't create file '.env'")
	defer file.Close()

	_, err = file.WriteString(tmpl)
	checkErr(err, "Can't write file '.env'")

	fmt.Println(".env file created")
}
