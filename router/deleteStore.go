package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/account"
	"web/config"
	structure "web/struct"
)

func DeleteStore(res http.ResponseWriter, req *http.Request) {
	db := config.Db

	//body값 받아오기
	post := structure.UserInfo{}
	json.NewDecoder(req.Body).Decode(&post)

	//로그인 체크
	userTocken, isLogin := account.Login(req)

	if !isLogin {
		res.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(res, "need login")
		panic("err")
	}

	//오너 체크
	checker := account.OwnerChecker(userTocken, post.StoreId)

	//ownerChecker error handler
	if !checker.Boolean {
		switch checker.Code {
		case 1:
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(res, "error in get user's information")
		case 2:
			res.WriteHeader(500)
			fmt.Fprint(res, "error in get user's store")
		}
		panic("error during checker error handler")
	}

	//delete query
	_, err := db.Exec("DELETE FROM store WHERE id=?", post.StoreId)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "error during delete row")
		panic(err.Error())
	}

	res.WriteHeader(http.StatusAccepted)
	fmt.Fprint(res, "deleted!")
}
