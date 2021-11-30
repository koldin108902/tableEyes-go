package account

import (
	"fmt"
	"net/http"
	createConn "web/config"
)

func Login(res http.ResponseWriter, req *http.Request) bool {
	db := createConn.ConnDb()

	//body값 읽기
	len := req.ContentLength
	reqTocken := make([]byte, len)

	query, err := db.Query("SELECT token FROM account")

	if err != nil {
		fmt.Printf("err(login db connect part): %v\n", err)
		panic(err.Error())
	}

	for query.Next() {
		var tocken string
		query.Scan(&tocken)

		if tocken == string(reqTocken) {
			return true
		}
	}

	return false
}
