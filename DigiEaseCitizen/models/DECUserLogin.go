package models

func DECUserLogin(objUserLoginRequest UserLoginRequest) UserLoginResponse {

	var objUserLoginResponse UserLoginResponse

	objUserLoginResponse.UserID = -1
	objUserLoginResponse.UserEmail = ""
	objUserLoginResponse.Password = ""
	objUserLoginResponse.Status = "false"
	objUserLoginResponse.ExceptionMessage = "Undefined exception - models/DECUserLogin"

	db := GetDBConnection()

	defer db.Close()

	err := db.QueryRow("select user_id, user_name, email from login where email=? and password=?",
		objUserLoginRequest.UserEmail, objUserLoginRequest.Password).Scan(&objUserLoginResponse.UserID, &objUserLoginResponse.UserName, &objUserLoginResponse.UserEmail)

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "No user found - models/DECUserLogin"

		return objUserLoginResponse

	}

	updateStatement, err := db.Prepare("update login set status=1 where email=? and password=?")

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not prepare update statement - models/DECUserLogin"

		return objUserLoginResponse

	}

	_, err = updateStatement.Exec(objUserLoginRequest.UserEmail, objUserLoginRequest.Password)

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not execute update statement - models/DECUserLogin"

		return objUserLoginResponse

	}

	objUserLoginResponse.Status = "true"
	objUserLoginResponse.ExceptionMessage = "Everything went okay. - models/DECUserLogin"

	return objUserLoginResponse

}
