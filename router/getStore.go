package router

import (
	"fmt"
	"net/http"
	"web/config"
	"web/structure"
)

func GetStore(res http.ResponseWriter, req *http.Request) {
	db := config.Db

	//get db data
	rows, err := db.Query("SELECT * FROM store")
	dataes := []structure.Row{}

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		row := structure.Row{}

		//각 데이터 지역변수에 담기
		err := rows.Scan(
			&row.Id, &row.OwnerId, &row.Name,
			&row.Description, &row.Category, &row.WaitingState,
			&row.Detail, &row.IsOpen, &row.NowPersonnel)

		if err != nil {
			panic(err.Error())
		}

		dataes = append(dataes, row) //전역 변수에 row값을 담음
	}

	res.WriteHeader(200)
	fmt.Fprint(res, dataes)

	defer db.Close()
}
