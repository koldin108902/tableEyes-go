package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/account"
	createConn "web/config"
	structure "web/struct"
)

func AddStore(res http.ResponseWriter, req *http.Request) {
	//db connnect
	db := createConn.Db
	defer db.Close()

	//login tocken check
	_, isLogin := account.Login(req)

	//로그인이 안되어 있으면
	if !isLogin {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprint(res, "")
		return
	}

	var body structure.Row
	json.NewDecoder(req.Body).Decode(&body) //body parse

	//insert query
	query := `INSERT INTO store (ownerId, name, description, category, waitingState, detail) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, body.OwnerId, body.Name, body.Description, body.Category, body.WaitingState, body.Detail)

	//inserting error handle
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "error! during inserting")

		fmt.Println(err)
		return
	}

	//response
	res.WriteHeader(201)
	fmt.Fprint(res, "created")
}
