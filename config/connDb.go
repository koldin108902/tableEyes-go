package config

import (
	"database/sql"
	"encoding/json"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type dbConfig struct {
	driver   string `json : "driver"`
	user     string `json : "user"`
	password string `json : "password"`
	database string `json : "dbName`
}

func configOpen() dbConfig {
	data, err := os.Open("./admin/db.json")
	var config dbConfig
	defer data.Close()

	if err != nil {
		panic(err.Error())
	}

	jsonParser := json.NewDecoder(data)
	jsonParser.Decode(&config)

	return config
}

func ConnDb() (db *sql.DB) {
	config := configOpen()

	db, err := sql.Open(config.driver, config.user+":"+config.password+"@/"+config.database)

	if err != nil {
		panic(err.Error())
	}

	return db
}
