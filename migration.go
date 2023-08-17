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
		HelpersMigrationCheck(err.Error())
		MigrationFiles(1)
		return
	}

	migration, err = business.ConfigsById("MIGRATION")
	if err != nil {
		panic(err.Error())
	}

	sequence, _ := strconv.Atoi(*migration.Value)
	MigrationFiles((sequence + 1))
}

func MigrationFiles(sequenceInitial int) {
	sequence := sequenceInitial
	db := helpers.DatabaseInstance()
	check := true
	for check {
		_, err := os.Stat("./migrations/" + strconv.FormatInt(int64(sequence), 10) + ".sql")
		if err != nil {
			check = false
			break
		}

		sql, err := os.ReadFile("./migrations/" + strconv.FormatInt(int64(sequence), 10) + ".sql")
		if err != nil {
			log.Fatal(err)
		}
		res, err := db.Query(string(sql))
		if err != nil {
			panic(err.Error())
		}

		defer res.Close()

		sequence++
	}
	defer db.Close()

	sequence--

	migration, err := business.ConfigsById("MIGRATION")
	if err != nil {
		panic(err.Error())
	}

	strSequence := strconv.FormatInt(int64(sequence), 10)
	migration.Value = &strSequence

	if migration.HasNew() {
		id := "MIGRATION"
		description := "Number sequence migration"
		migration.Id = &id
		migration.Description = &description
		migration.Insert(helpers.DatabaseInstance())
	} else {
		migration.Update(helpers.DatabaseInstance())
	}

}

func HelpersMigrationCheck(str string) bool {
	r, _ := regexp.Compile(`Table \'test\.configs\' doesn\'t exist`)
	return r.MatchString(str)
}
