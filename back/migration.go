package main

import (
	"fmt"
	"log"
	"napptest/business"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func Migration() {
	migration, err := business.ConfigsById("MIGRATION")
	if err != nil {
		HelpersMigrationCheck(err.Error())
		return
	}

	fmt.Println(migration.Description)
}

func MigrationFiles() {
	entries, err := os.ReadDir("./migrations")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}

func HelpersMigrationCheck(str string) {
	r, _ := regexp.Compile(`Table \'test\.configs\' doesn\'t exist`)
	if r.MatchString(str) {
		MigrationFiles()
	}
}
