package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/config"
)

type Tocken struct {
	Tocken string
}

func Login(req *http.Request) (string, bool) {
	db := config.Db

	//body값 읽기
	var body Tocken
	var reqTocken string

	json.NewDecoder(req.Body).Decode(&body)

	reqTocken = body.Tocken //user tocken

	//get db
	query, err := db.Query("SELECT token FROM account")

	if err != nil {
		fmt.Printf("err(login db connect part): %v\n", err)
		panic(err.Error())
	}

	for query.Next() { //각 row값
		var tocken string
		query.Scan(&tocken)

		if tocken == reqTocken { //로그인 기록이 있는 토큰일시
			return "error", true
		}
	}

	return reqTocken, false
}
