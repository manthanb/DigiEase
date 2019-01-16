package models

func DEPUserLogin(objUserLoginRequest UserLoginRequest) UserLoginResponse {

	var objUserLoginResponse UserLoginResponse

	objUserLoginResponse.Status = "false"
	objUserLoginResponse.AreaCode = -1
	objUserLoginResponse.ExceptionMessage = "Undefined Exception"

	db := GetDBConnection()
	defer db.Close()

	query, err := db.Prepare("select area_code from police_user where user_name=? and password=?")
	HandleError(err)

	if err != nil {
		objUserLoginResponse.ExceptionMessage = "could not prepare query"
		return objUserLoginResponse
	}

	err = query.QueryRow(objUserLoginRequest.UserName, objUserLoginRequest.Password).Scan(&objUserLoginResponse.AreaCode)
	HandleError(err)

	if err != nil {
		objUserLoginResponse.ExceptionMessage = "invalid credentials"
		return objUserLoginResponse
	}

	objUserLoginResponse.Status = "true"
	objUserLoginResponse.ExceptionMessage = "-"

	return objUserLoginResponse

}
