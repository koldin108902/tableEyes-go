package config

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"dbName"`
}

func configOpen() DBConfig {
	var config DBConfig                                       //config struct
	data, err := os.Open("F:/files/golang/web/admin/db.json") //jsonFile memory
	if err != nil {
		panic(err.Error())
	}
	defer data.Close()

	byteVal, _ := ioutil.ReadAll(data) //byte jsonFile
	json.Unmarshal(byteVal, &config)

	return config
}

var Db = connDb()

func connDb() (db *sql.DB) {
	config := configOpen()

	db, err := sql.Open(config.Driver, config.User+":"+config.Password+"@/"+config.Database)

	if err != nil {
		panic(err.Error())
	}

	return db
}
