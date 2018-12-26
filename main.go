// Package main for creating dump
package main

import (
	"flag"
	"path"
	"time"

	"github.com/joho/godotenv"
)

var env map[string]string
var createEnv bool

func main() {
	if createEnv {
		return
	}
	// create dump/date folder
	dumpDir(env["DUMP_SUB_DIR"])

	dump := GetDump()

	MakeDumpFiles(dump, false)
	MakeDumpFiles(dump, true)
}

func init() {
	flag.BoolVar(&createEnv, "make-env", false, "create .env boilerplate file")
	flag.Parse()

	if createEnv {
		makeEnv()
		return
	}

	err := godotenv.Load()
	if err != nil {
		envLoadErrorMessage("Error loading .env file")
	}
	checkErr(err, "Error loading .env file")

	env, err = godotenv.Read()
	if err != nil {
		envLoadErrorMessage("Error loading .env file variables")
	}

	validKeys := []string{"TITLE", "DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD"}
	checkParams(&env, validKeys)

	if env["DEBUG"] != "false" {
		env["DUMP_SUB_DIR"] = path.Join(env["DUMP_DIR"], time.Now().Format("2006-01-02_15:04"))
	} else {
		env["DUMP_SUB_DIR"] = path.Join(env["DUMP_DIR"], time.Now().Format("2006-01-02"))
	}
}
