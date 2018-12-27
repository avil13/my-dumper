// Package main for creating dump
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

var env map[string]string
var currentDir string
var createEnv bool
var dumpAll bool
var importSQLFile string

var logFile *os.File

func main() {
	if logFile != nil {
		defer logFile.Close()
	}

	if createEnv {
		return
	}

	if importSQLFile != "" {
		Insert(importSQLFile)
		return
	}

	createDump()
}

func init() {
	flag.BoolVar(&createEnv, "make-env", false, "create .env boilerplate file")
	flag.BoolVar(&dumpAll, "all", false, "create all dump files")
	flag.StringVar(&importSQLFile, "import", "", "import dump file")
	flag.Parse()

	ex, err := os.Executable()
	checkErr(err, "Executeble")
	currentDir = filepath.Dir(ex)

	if createEnv {
		makeEnv()
		return
	}

	envFile := path.Join(currentDir, ".env")
	env, err = godotenv.Read(envFile)
	checkErr(err, "Error loading .env file variables")
	loggerInit()

	validKeys := []string{"TITLE", "DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD"}
	checkParams(&env, validKeys)

	env["DUMP_DIR"] = path.Join(currentDir, env["DUMP_DIR"])
	env["DUMP_SUB_DIR"] = path.Join(env["DUMP_DIR"], time.Now().Format("2006-01-02"))
}

// createDump function for create dump files
func createDump() {
	// create dump/date folder
	dumpDir(env["DUMP_SUB_DIR"])

	dump := GetDump()

	if dumpAll || env["DUMP_CREATE"] != "false" {
		MakeDumpFiles(dump, false)
	}

	if dumpAll || env["DUMP_INSERT"] != "false" {
		MakeDumpFiles(dump, true)
	}

	if !dumpAll && (env["DUMP_INSERT"] == "false" && env["DUMP_CREATE"] == "false") {
		fmt.Println("⚠️  None of the files are not created")
	}

	dirSize, err := DirSize(env["DUMP_DIR"])
	checkErr(err, "Dir size error")
	fmt.Println("Dumps total size:", SizeToString(dirSize))
}

// loggerInit write logs
func loggerInit() {
	if env["LOG"] != "true" {
		return
	}

	var logFileName string
	if env["LOG_ERROR_FILE"] == "" {
		logFileName = "error.log"
	} else {
		logFileName = env["LOG_ERROR_FILE"]
	}
	logFileName = path.Join(currentDir, logFileName)

	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// defer logFile.Close() // close in the main()

	log.SetOutput(logFile)
}
