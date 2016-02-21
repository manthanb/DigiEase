package models

import "strconv"

func DECUserRegistration(objUserRegistrationRequest UserRegistrationRequest) UserLoginResponse {

	var objUserLoginResponse UserLoginResponse
	var intUserID int64

	objUserLoginResponse.Status = "false"
	objUserLoginResponse.UserID = -1
	objUserLoginResponse.UserName = ""
	objUserLoginResponse.UserEmail = ""
	objUserLoginResponse.ExceptionMessage = "Undefined exception - models/DECUserRegistration"

	db := GetDBConnection()
	defer db.Close()

	insertUserStatement, err := db.Prepare(`insert into users(email, password, user_name, first_name, middle_name, 
		last_name, gender, dob_date, dob_month, dob_year, address1, address2, address3, city, state, pincode,
		mobile, landline, aadhar_card, pan_card, passport, driving_license, photo, address_proof) 
		values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)

	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not prepare insert user statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	intDate, err := strconv.ParseInt(objUserRegistrationRequest.Date, 10, 64)
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not convert date to int - models/DECUserRegistration"

		return objUserLoginResponse

	}

	intMonth, err := strconv.ParseInt(objUserRegistrationRequest.Month, 10, 64)
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not convert month to int - models/DECUserRegistration"

		return objUserLoginResponse

	}

	intYear, err := strconv.ParseInt(objUserRegistrationRequest.Year, 10, 64)
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not convert year to int - models/DECUserRegistration"

		return objUserLoginResponse

	}

	_, err = insertUserStatement.Exec(objUserRegistrationRequest.Email, objUserRegistrationRequest.Password,
		objUserRegistrationRequest.UserName, objUserRegistrationRequest.FirstName,
		objUserRegistrationRequest.MiddleName, objUserRegistrationRequest.LastName,
		objUserRegistrationRequest.Gender, intDate, intMonth, intYear, objUserRegistrationRequest.Address1,
		objUserRegistrationRequest.Address2, objUserRegistrationRequest.Address3,
		objUserRegistrationRequest.City, objUserRegistrationRequest.State, objUserRegistrationRequest.PinCode,
		objUserRegistrationRequest.Mobile, objUserRegistrationRequest.Landline,
		objUserRegistrationRequest.AadharCard, objUserRegistrationRequest.PanCard,
		objUserRegistrationRequest.Passport, objUserRegistrationRequest.DrivingLicense,
		objUserRegistrationRequest.Photo, objUserRegistrationRequest.AddressProof)

	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not execute insert user - models/DECUserRegistration"

		return objUserLoginResponse

	}

	getUserIDStatement, err := db.Prepare("select user_id from users where email = ?")
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not prepare select user_id statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	err = getUserIDStatement.QueryRow(objUserRegistrationRequest.Email).Scan(&intUserID)
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not execute select user_id statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	inputLoginStatement, err := db.Prepare("insert into login values(?,?,?,?,?)")
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not prepare insert into statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	_, err = inputLoginStatement.Exec(intUserID, objUserRegistrationRequest.UserName,
		objUserRegistrationRequest.Email, objUserRegistrationRequest.Password, 1)
	HandleError(err)

	if err != nil {

		objUserLoginResponse.Status = "false"
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserName = ""
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.ExceptionMessage = "Could not execute select user_id statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	objUserLoginResponse.Status = "true"
	objUserLoginResponse.UserID = intUserID
	objUserLoginResponse.UserName = objUserRegistrationRequest.UserName
	objUserLoginResponse.UserEmail = objUserRegistrationRequest.Email
	objUserLoginResponse.ExceptionMessage = "-"

	return objUserLoginResponse

}
