// Package main for creating dump
package main

import (
	"github.com/joho/godotenv"
)

var env map[string]string

func main() {

}

func init() {
	err := godotenv.Load()
	checkErr(err, "Error loading .env file")

	env, err = godotenv.Read()
	checkErr(err, "Error loading .env file variables")

	// dumpDir(env["DUMP_DIR"])
	// GetDump()
	// fmt.Println(env)
	validKeys := []string{"DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD"}
	checkParams(&env, validKeys)
}
