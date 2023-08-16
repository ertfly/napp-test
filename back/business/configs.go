package business

import (
	"database/sql"
	"log"
	"napptest/helpers"

	_ "github.com/go-sql-driver/mysql"
)

func ConfigsNew() Configs {
	row := Configs{}
	return row
}

func ConfigsTable() string {
	return "configs"
}

func ConfigsById(id string) (Configs, error) {
	db := helpers.DatabaseInstance()
	row := Configs{}
	res, err := db.Query("SELECT * FROM configs WHERE `id` = ? limit 1", id)
	if err != nil {
		return row, err
	}
	defer db.Close()

	for res.Next() {
		res.Scan(&row.Id, &row.Value, &row.Description)
		return row, nil
	}

	return row, nil
}

type Configs struct {
	Id          *string
	Value       *string
	Description *string
}

func (row Configs) Insert(db *sql.DB) {
	_, err := db.Exec("INSERT INTO configs ( `id`, `value`, `description` ) values (?, ?, ?)", row.Id, row.Value, row.Description)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()
}

func (row Configs) Update(db *sql.DB) {
	_, err := db.Exec("UPDATE configs SET `value` = ?, `description` = ? WHERE `id` = ?", row.Value, row.Description, row.Id)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()
}

func (row Configs) HasNew() bool {
	return row.Id == nil
}
