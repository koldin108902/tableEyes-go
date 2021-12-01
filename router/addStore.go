package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/account"
	createConn "web/config"
	_struct "web/struct"
)

func AddStore(res http.ResponseWriter, req *http.Request) {
	//db connnect
	db := createConn.ConnDb()
	defer db.Close()

	//login tocken check
	if !account.Login(res, req) {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "")
		return
	}

	var body _struct.Row
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
