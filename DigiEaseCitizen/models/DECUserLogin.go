package models

func DECUserLogin(objUserLoginRequest UserLoginRequest) UserLoginResponse {

	var objUserLoginResponse UserLoginResponse

	objUserLoginResponse.Status = "false"
	objUserLoginResponse.UserID = -1
	objUserLoginResponse.UserName = ""
	objUserLoginResponse.UserEmail = ""
	objUserLoginResponse.ExceptionMessage = "Undefined exception - models/DECUserLogin"

	db := GetDBConnection()
	defer db.Close()

	users, err := db.Query("select user_id, user_name, email from login where email=? and password=?",
		objUserLoginRequest.UserEmail, objUserLoginRequest.Password)

	HandleError(err)
	defer users.Close()

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not execute select query - models/DECUserLogin"

		return objUserLoginResponse

	}

	for users.Next() {

		err := users.Scan(&objUserLoginResponse.UserID, &objUserLoginResponse.UserName,
			&objUserLoginResponse.UserEmail)

		HandleError(err)

		if err != nil {

			objUserLoginResponse.Status = "false"
			objUserLoginResponse.UserID = -1
			objUserLoginResponse.UserName = ""
			objUserLoginResponse.UserEmail = ""
			objUserLoginResponse.ExceptionMessage = "Could not scan values - models/DECUserLogin"

			return objUserLoginResponse

		}

	}

	updateStatement, err := db.Prepare("update login set status=1 where email=? and password=?")
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not prepare update statement - models/DECUserLogin"

		return objUserLoginResponse

	}

	_, err = updateStatement.Exec(objUserLoginRequest.UserEmail, objUserLoginRequest.Password)
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not execute update statement - models/DECUserLogin"

		return objUserLoginResponse

	}

	objUserLoginResponse.Status = "true"
	objUserLoginResponse.ExceptionMessage = "-"

	return objUserLoginResponse

}
