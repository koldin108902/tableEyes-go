package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	var config DBConfig
	data, err := os.Open("F:/files/golang/web/admin/db.json")
	if err != nil {
		panic(err.Error())
	}
	defer data.Close()

	byteVal, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteVal, &config)

	fmt.Println(string(config.Database))
	fmt.Print(config)

	return config
}

func ConnDb() (db *sql.DB) {
	config := configOpen()

	fmt.Println(config.Driver, config.User+":"+config.Password+"@/"+config.Database)
	db, err := sql.Open(config.Driver, config.User+":"+config.Password+"@/"+config.Database)

	if err != nil {
		panic(err.Error())
	}

	return db
}
