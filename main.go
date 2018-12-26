// Package main for creating dump
package main

import (
	"path"
	"time"

	"github.com/joho/godotenv"
)

var env map[string]string

func main() {
	// create dump/date folder
	dumpDir(env["DUMP_SUB_DIR"])

	dump := GetDump()

	MakeDumpFiles(dump, false)
	MakeDumpFiles(dump, true)
}

func init() {
	err := godotenv.Load()
	checkErr(err, "Error loading .env file")

	env, err = godotenv.Read()
	checkErr(err, "Error loading .env file variables")

	validKeys := []string{"DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD"}
	checkParams(&env, validKeys)

	// env["DUMP_SUB_DIR"] = path.Join(env["DUMP_DIR"], time.Now().Format("2006-01-02_15:04"))
	env["DUMP_SUB_DIR"] = path.Join(env["DUMP_DIR"], time.Now().Format("2006-01-02"))
}
