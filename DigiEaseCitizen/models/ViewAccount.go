package models

func ViewAccount(objUserEmailStruct UserEmailStruct) EditAccountRequest {
	var accountDetails EditAccountRequest

	db := GetDBConnection()

	defer db.Close()

	err := db.QueryRow("select email, password, address1, address2, address3, city, state, pincode, mobile, landline from users where email = ?", objUserEmailStruct.UserEmail).Scan(&accountDetails.UserEmail, &accountDetails.Password, &accountDetails.Address1, &accountDetails.Address2, &accountDetails.Address3, &accountDetails.City, &accountDetails.State, &accountDetails.PinCode, &accountDetails.Mobile, &accountDetails.Landline)

	HandleError(err)

	if err != nil {

		return accountDetails
	}

	return accountDetails

}
