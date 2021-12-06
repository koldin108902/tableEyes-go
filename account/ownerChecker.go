package account

import "web/config"

type _response struct {
	Message string
	Code    int
	Boolean bool
}

func OwnerChecker(tocken string, storeId int) _response {
	db := config.Db
	var ownerId int
	var res _response
	res.Boolean = false

	//tocken to user id
	err := db.QueryRow("SELECT id FROM account WHERE tocken=?", tocken).Scan(&ownerId)

	//error handler during user id
	if err != nil {
		res.Code = 1
		res.Message = "error in get owner id"
		return res
	}

	//is store's owner
	err = db.QueryRow("SELECT id FROM store WHERE ownerId=?", ownerId).Err()

	//error handler during get user's store
	if err != nil {
		res.Code = 2
		res.Message = "error in get store info"
		return res
	}

	res.Code = 0
	res.Message = "accept"
	res.Boolean = true
	return res
}
