// Package main for creating dump
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
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

	if env["DUMP_CREATE"] != "false" {
		MakeDumpFiles(dump, false)
	}

	if env["DUMP_INSERT"] != "false" {
		MakeDumpFiles(dump, true)
	}

	if env["DUMP_INSERT"] == "false" && env["DUMP_CREATE"] == "false" {
		fmt.Println("⚠️  None of the files are not created")
	}

	dirSize, err := DirSize(env["DUMP_DIR"])
	checkErr(err, "Dir size error")
	fmt.Println("Dumps total size:", SizeToString(dirSize))
}

func init() {
	flag.BoolVar(&createEnv, "make-env", false, "create .env boilerplate file")
	flag.Parse()

	if createEnv {
		makeEnv()
		return
	}

	ex, err := os.Executable()
	checkErr(err, "Executeble")

	envFile := path.Join(filepath.Dir(ex), ".env")

	env, err = godotenv.Read(envFile)
	checkErr(err, "Error loading .env file variables")

	validKeys := []string{"TITLE", "DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD"}
	checkParams(&env, validKeys)

	env["DUMP_SUB_DIR"] = path.Join(env["DUMP_DIR"], time.Now().Format("2006-01-02"))
}
