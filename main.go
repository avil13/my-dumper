// Package main for creating dump
package main

import (
	"flag"
	"path"
	"time"

	"github.com/joho/godotenv"
)

var env map[string]string
var createEnv *bool

func main() {
	if *createEnv {
		return
	}
	// create dump/date folder
	dumpDir(env["DUMP_SUB_DIR"])

	dump := GetDump()

	MakeDumpFiles(dump, false)
	MakeDumpFiles(dump, true)
}

func init() {
	createEnv = flag.Bool("--make-env", false, "create .env boilerplate file")

	if *createEnv {
		makeEnv()
		return
	}

	err := godotenv.Load()
	checkErr(err, "Error loading .env file")

	env, err = godotenv.Read()
	checkErr(err, "Error loading .env file variables")

	validKeys := []string{"DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD"}
	checkParams(&env, validKeys)

	// env["DUMP_SUB_DIR"] = path.Join(env["DUMP_DIR"], time.Now().Format("2006-01-02_15:04"))
	env["DUMP_SUB_DIR"] = path.Join(env["DUMP_DIR"], time.Now().Format("2006-01-02"))
}
