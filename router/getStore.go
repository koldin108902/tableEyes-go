package router

import (
	"fmt"
	"net/http"
	createConn "web/config"
	_struct "web/struct"
)

func GetStore(res http.ResponseWriter, req *http.Request) {
	db := createConn.ConnDb()

	query, err := db.Prepare("SELECT * FROM store")

	if err != nil {
		panic(err.Error())
	}

	rows, err := query.Query()
	dataes := []_struct.Row{}

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		row := _struct.Row{}

		err := rows.Scan(
			&row.Id, &row.OwnerId, &row.Name,
			&row.Description, &row.Category, &row.WaitingState,
			&row.Detail, &row.IsOpen, &row.NowPersonnel)

		if err != nil {
			panic(err.Error())
		}

		dataes = append(dataes, row)
	}

	res.WriteHeader(200)
	fmt.Fprint(res, dataes)

	defer db.Close()
}
