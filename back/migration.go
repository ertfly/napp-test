package main

import (
	"fmt"
	"log"
	"napptest/business"
	"napptest/helpers"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func Migration() {
	migration, err := business.ConfigsById("MIGRATION")
	if err != nil {
		if !HelpersMigrationCheck(err.Error()) {
			panic(err.Error())
		}
	}

	migration, err = business.ConfigsById("MIGRATION")
	if err != nil {
		panic(err.Error())
	}

	if migration.HasNew() {
		fmt.Println("insert")
	} else {
		fmt.Println("update")
	}

	fmt.Println(migration.Description)
}

func MigrationFiles() {
	entries, err := os.ReadDir("./migrations")
	if err != nil {
		log.Fatal(err)
	}

	db := helpers.DatabaseInstance()
	for _, e := range entries {
		sql, err := os.ReadFile("./migrations/" + e.Name())
		if err != nil {
			log.Fatal(err)
		}
		res, err := db.Query(string(sql))
		if err != nil {
			panic(err.Error())
		}

		defer res.Close()
	}

	defer db.Close()
}

func HelpersMigrationCheck(str string) bool {
	r, _ := regexp.Compile(`Table \'test\.configs\' doesn\'t exist`)
	if r.MatchString(str) {
		MigrationFiles()
		return true
	}

	return false
}
