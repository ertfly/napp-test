package main

import (
	"log"
	"napptest/business"
	"napptest/helpers"
	"os"
	"regexp"
	"strconv"

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
		*migration.Id = "MIGRATION"
		*migration.Value = "0"
		*migration.Description = "Number sequence migration"
		migration.Insert(helpers.DatabaseInstance())
	}

	sequence, _ := strconv.Atoi(*migration.Value)
	MigrationFiles(sequence)
}

func MigrationFiles(sequence int) {
	db := helpers.DatabaseInstance()
	check := true
	for check {
		_, err := os.Stat("./migrations/" + string(sequence) + ".sql")
		if err != nil {
			check = false
			return
		}

		sql, err := os.ReadFile("./migrations/" + string(sequence) + ".sql")
		if err != nil {
			log.Fatal(err)
		}
		res, err := db.Query(string(sql))
		if err != nil {
			panic(err.Error())
		}

		defer res.Close()
	}
}

func HelpersMigrationCheck(str string) bool {
	r, _ := regexp.Compile(`Table \'test\.configs\' doesn\'t exist`)
	return r.MatchString(str)
}
