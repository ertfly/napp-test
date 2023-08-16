package helpers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseInstance() *sql.DB {
	if db == nil {
		host := GetEnv("DB_HOST")
		port := GetEnv("DB_PORT")
		name := GetEnv("DB_NAME")
		user := GetEnv("DB_USER")
		pass := GetEnv("DB_PASS")

		db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+name)
		if err != nil {
			panic("Error connecting database: " + err.Error())
		}
		return db
	}

	return db
}
