package models

func DEPUserLogin(objUserLoginRequest UserLoginRequest) UserLoginResponse {

	var objUserLoginResponse UserLoginResponse
	var intAreaCode int64

	objUserLoginResponse.Status = "false"
	objUserLoginResponse.AreaCode = -1
	objUserLoginResponse.ExceptionMessage = "Undefined exception"

	db := GetDBConnection()

	query, err := db.Prepare("select area_code from police_user where user_name=? and password=?")
	HandleError(err)

	if err != nil {
		objUserLoginResponse.ExceptionMessage = "could not prepare select query"
		return objUserLoginResponse
	}

	err = query.QueryRow(objUserLoginRequest.UserName, objUserLoginRequest.Password).Scan(&intAreaCode)
	HandleError(err)

	if err != nil {
		objUserLoginResponse.ExceptionMessage = "invalid credentials"
		return objUserLoginResponse
	}

	objUserLoginResponse.Status = "true"
	objUserLoginResponse.AreaCode = intAreaCode
	objUserLoginResponse.ExceptionMessage = "-"

	return objUserLoginResponse

}
