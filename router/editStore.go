package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/account"
	"web/config"
	structure "web/struct"
)

func EditStore(res http.ResponseWriter, req *http.Request) {
	db := config.Db
	defer db.Close()

	//get body value
	var body structure.Row
	json.NewDecoder(req.Body).Decode(&body)

	//check login and get tocken
	tocken, isLogin := account.Login(req)

	if !isLogin {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprint(res, "need login")
		return
	}

	//매장의 오너가 맞는지 체크
	cheker := account.OwnerChecker(tocken, body.Id)

	//checker error handler
	if !cheker.Boolean {
		switch cheker.Code {
		case 1:
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(res, "error in get user's information")
		case 2:
			res.WriteHeader(500)
			fmt.Fprint(res, "error in get user's store")
		}
	}

	//update
	_, err := db.Query("UPDATE store SET (name, description, category, waitingState, detail) VALUE (?, ?, ?, ?, ?) WHERE id=?", body.Name, body.Description, body.Category, body.WaitingState, body.Detail, body.Id)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "error during update")
	}

	res.WriteHeader(http.StatusUpgradeRequired)
	fmt.Fprint(res, "update accepted")
}
