package models

func EditAccount(objEditAccountRequest EditAccountRequest) EditAccountResponse {
	var objEditAccountResponse EditAccountResponse
	var currentData EditAccountRequest

	objEditAccountResponse.ExceptionMessage = "Undefined exception - models/EditAccount"
	objEditAccountResponse.Status = "false"

	db := GetDBConnection()

	defer db.Close()

	updateStatement, err := db.Prepare("update users set email = ?, password = ?, address1 = ?, address2 = ?, address3 = ?, city = ?, state = ?, pincode = ?, mobile = ?, landline = ? where email = ?")

	HandleError(err)

	if err != nil {
		objEditAccountResponse.ExceptionMessage = "Could not prepare update statement - models/EditAccount"

		return objEditAccountResponse
	}

	err = db.QueryRow("select email, password, address1, address2, address3, city, state, pincode, mobile, landline from users where email = ?",
		objEditAccountRequest.UserEmail).Scan(&currentData.UserEmail, &currentData.Password, &currentData.Address1, &currentData.Address2, &currentData.Address3, &currentData.City, &currentData.State, &currentData.PinCode, &currentData.Mobile, &currentData.Landline)

	HandleError(err)

	if err != nil {
		objEditAccountResponse.ExceptionMessage = "Could not get existing data - models/EditAccount"

		return objEditAccountResponse
	}

	if objEditAccountRequest.UserEmail == "" {
		objEditAccountRequest.UserEmail = currentData.UserEmail
	}
	if objEditAccountRequest.Password == "" {
		objEditAccountRequest.Password = currentData.Password
	}
	if objEditAccountRequest.Address1 == "" {
		objEditAccountRequest.Address1 = currentData.Address1
	}
	if objEditAccountRequest.Address2 == "" {
		objEditAccountRequest.Address2 = currentData.Address2
	}
	if objEditAccountRequest.Address3 == "" {
		objEditAccountRequest.Address3 = currentData.Address3
	}

	//halfway

	if objEditAccountRequest.City == "" {
		objEditAccountRequest.City = currentData.City
	}
	if objEditAccountRequest.State == "" {
		objEditAccountRequest.State = currentData.State
	}
	if objEditAccountRequest.PinCode == "" {
		objEditAccountRequest.PinCode = currentData.PinCode
	}
	if objEditAccountRequest.Mobile == "" {
		objEditAccountRequest.Mobile = currentData.Mobile
	}
	if objEditAccountRequest.Landline == "" {
		objEditAccountRequest.Landline = currentData.Landline
	}

	_, err = updateStatement.Exec(objEditAccountRequest.UserEmail, objEditAccountRequest.Password, objEditAccountRequest.Address1, objEditAccountRequest.Address2, objEditAccountRequest.Address3, objEditAccountRequest.City, objEditAccountRequest.State, objEditAccountRequest.PinCode, objEditAccountRequest.Mobile, objEditAccountRequest.Landline, objEditAccountRequest.UserEmail)

	HandleError(err)

	if err != nil {
		objEditAccountResponse.ExceptionMessage = "Could not execute update statement - models/EditAccount"

		return objEditAccountResponse
	}

	updateStatement, err = db.Prepare("update login set password = ? where email = ?")

	HandleError(err)

	if err != nil {

		objEditAccountResponse.ExceptionMessage = "Could not prepare to update login table - models/EditAccount"

		return objEditAccountResponse

	}

	_, err = updateStatement.Exec(objEditAccountRequest.Password, objEditAccountRequest.UserEmail)

	HandleError(err)

	if err != nil {

		objEditAccountResponse.ExceptionMessage = "Could not update login table - models/EditAccount"

		return objEditAccountResponse
	}

	objEditAccountResponse.Status = "true"
	objEditAccountResponse.ExceptionMessage = "Everything went okay. - models/EditAccount"

	return objEditAccountResponse

}
