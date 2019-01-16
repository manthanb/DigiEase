package models

import (
	"strconv"
)

func DECUserRegistration(objUserRegistrationRequest UserRegistrationRequest) UserLoginResponse {

	var objUserLoginResponse UserLoginResponse
	var intUserID int64

	objUserLoginResponse.UserID = -1
	objUserLoginResponse.UserEmail = ""
	objUserLoginResponse.Password = ""
	objUserLoginResponse.Status = "false"
	objUserLoginResponse.ExceptionMessage = "Undefined exception - models/DECUserRegistration"

	db := GetDBConnection()

	defer db.Close()

	insertUserStatement, err := db.Prepare("insert into users(email, password, user_name, first_name, middle_name, last_name, gender, dob_date, dob_month, dob_year, address1, address2, address3, city, state, pincode, mobile, landline, aadhar_card, pan_card, passport, driving_license, photo, address_proof) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not prepare insert-into-users statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	intDate, err := strconv.ParseInt(objUserRegistrationRequest.Date, 10, 64)

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not convert date to int - models/DECUserRegistration"

		return objUserLoginResponse

	}

	intMonth, err := strconv.ParseInt(objUserRegistrationRequest.Month, 10, 64)

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not convert month to int - models/DECUserRegistration"

		return objUserLoginResponse

	}

	intYear, err := strconv.ParseInt(objUserRegistrationRequest.Year, 10, 64)

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not convert year to int - models/DECUserRegistration"

		return objUserLoginResponse

	}

	_, err = insertUserStatement.Exec(objUserRegistrationRequest.UserEmail, objUserRegistrationRequest.Password, objUserRegistrationRequest.UserName, objUserRegistrationRequest.FirstName, objUserRegistrationRequest.MiddleName, objUserRegistrationRequest.LastName, objUserRegistrationRequest.Gender, intDate, intMonth, intYear, objUserRegistrationRequest.Address1, objUserRegistrationRequest.Address2, objUserRegistrationRequest.Address3, objUserRegistrationRequest.City, objUserRegistrationRequest.State, objUserRegistrationRequest.PinCode, objUserRegistrationRequest.Mobile, objUserRegistrationRequest.Landline, objUserRegistrationRequest.AadharCard, objUserRegistrationRequest.PanCard, objUserRegistrationRequest.Passport, objUserRegistrationRequest.DrivingLicense, objUserRegistrationRequest.Photo, objUserRegistrationRequest.AddressProof)

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not execute insert-into-users statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	err = db.QueryRow("select user_id from users where email=?",
		objUserRegistrationRequest.UserEmail).Scan(&intUserID)

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not get userID - models/DECUserRegistration"

		return objUserLoginResponse

	}

	insertIntoLoginStatement, err := db.Prepare("insert into login values(?, ?, ?, ?, ?)")

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not prepare insert-into-login statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	_, err = insertIntoLoginStatement.Exec(intUserID, objUserRegistrationRequest.UserName, objUserRegistrationRequest.UserEmail, objUserRegistrationRequest.Password, 1)

	HandleError(err)

	if err != nil {
		objUserLoginResponse.UserID = -1
		objUserLoginResponse.UserEmail = ""
		objUserLoginResponse.Password = ""
		objUserLoginResponse.Status = "false"
		objUserLoginResponse.ExceptionMessage = "Could not execute insert-into-login statement - models/DECUserRegistration"

		return objUserLoginResponse

	}

	objUserLoginResponse.UserID = intUserID
	objUserLoginResponse.Status = "true"
	objUserLoginResponse.UserEmail = objUserRegistrationRequest.UserEmail
	objUserLoginResponse.UserName = objUserRegistrationRequest.UserName
	objUserLoginResponse.ExceptionMessage = "Everything went okay. - models/DECUserRegistration"

	return objUserLoginResponse

}
