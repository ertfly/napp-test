package business

import (
	"database/sql"
	"log"
	"napptest/helpers"

	_ "github.com/go-sql-driver/mysql"
)

func ProductsNew() Configs {
	row := Configs{}
	return row
}

func ProductsById(id string) (Products, error) {
	db := helpers.DatabaseInstance()
	row := Products{}
	res, err := db.Query("SELECT * FROM products WHERE `id` = ? limit 1", id)
	if err != nil {
		return row, err
	}
	defer db.Close()

	for res.Next() {
		res.Scan(&row.Id, &row.Sku, &row.Name, &row.PriceUnit, &row.CreatedAt, &row.UpdatedAt, &row.Trash)
		return row, nil
	}

	return row, nil
}

func ProductsBySku(sku string) (Products, error) {
	db := helpers.DatabaseInstance()
	row := Products{}
	res, err := db.Query("SELECT * FROM products WHERE `sku` = ? and trash = 0 limit 1", sku)
	if err != nil {
		return row, err
	}
	defer db.Close()

	for res.Next() {
		res.Scan(&row.Id, &row.Sku, &row.Name, &row.PriceUnit, &row.CreatedAt, &row.UpdatedAt, &row.Trash)
		return row, nil
	}

	return row, nil
}

type Products struct {
	Id        *int64
	Sku       *string
	Name      *string
	PriceUnit *float64
	CreatedAt *string
	UpdatedAt *string
	Trash     *bool
}

func (row *Products) Insert(db *sql.DB) {
	trash := false
	row.Trash = &trash
	res, err := db.Exec("INSERT INTO products ( `sku`, `name`, `price_unit`, `created_at`, `updated_at`, `trash` ) values (?, ?, ?, ?, ?, ?)", row.Sku, row.Name, row.PriceUnit, row.CreatedAt, row.UpdatedAt, row.Trash)
	if err != nil {
		log.Fatalln(err.Error())
	}

	lastId, _ := res.LastInsertId()
	row.Id = &lastId

	defer db.Close()
}

func (row Products) Update(db *sql.DB) {
	_, err := db.Exec("UPDATE products SET sku=?, name=?, price_unit=?, created_at=?, updated_at=?, trash=? WHERE `id` = ?", row.Sku, row.Name, row.PriceUnit, row.CreatedAt, row.UpdatedAt, row.Trash, row.Id)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()
}

func (row Products) HasNew() bool {
	return row.Id == nil
}
