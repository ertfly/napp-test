package business

import "napptest/helpers"

const TABLE = "configs"

func ConfigsById(id string) (ConfigsEntity, error) {
	db := helpers.DatabaseInstance()
	row := ConfigsEntity{}
	res, err := db.Query("SELECT * FROM " + TABLE + " limit 1")
	if err != nil {
		return row, err
	}

	for res.Next() {
		res.Scan(&row.Id, &row.Value, &row.Description)
		return row, nil
	}

	return row, nil
}

type ConfigsEntity struct {
	Id          int
	Value       string
	Description string
}
