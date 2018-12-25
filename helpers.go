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
	for _, v := range params {
		if _, ok := (*data)[v]; ok != true {
			log.Fatalf("Invalid env parameter %s ", v)
		}
	}
}

// dumpDir - make dump folder
func dumpDir(dir string) (err error) {
	if dir == "" {
		dir = "dumps"
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	return
}
