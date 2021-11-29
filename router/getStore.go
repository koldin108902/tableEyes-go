package router

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type row struct {
	id           int
	ownerId      int
	name         string
	description  string
	category     string
	waitingState bool
	detail       string
	isOpen       bool
	nowPersonnel int
}

func connDb() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:#koldin13579@/tableEyes")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetStore(res http.ResponseWriter, req *http.Request) {
	db := connDb()

	query, err := db.Prepare("SELECT * FROM store")

	if err != nil {
		panic(err.Error())
	}

	rows, err := query.Query()
	dataes := []row{}

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var row row

		err := rows.Scan(
			&row.id, &row.ownerId, &row.name,
			&row.description, &row.category, &row.waitingState,
			&row.detail, &row.isOpen, &row.nowPersonnel)

		if err != nil {
			panic(err.Error())
		}

		dataes = append(dataes, row)
	}

	res.WriteHeader(200)
	fmt.Fprint(res, dataes)

	defer db.Close()
}
