package models

func CheckDuplicateEmail(objUserEmailStruct UserEmailStruct) DuplicateEmail {

	var objDuplicateEmail DuplicateEmail
	var numberOfCopies int64

	objDuplicateEmail.IsEmailValid = "false"

	db := GetDBConnection()

	defer db.Close()

	rows, err := db.Query("select * from users where email = ?", objUserEmailStruct.UserEmail)

	HandleError(err)

	if err != nil {

		return objDuplicateEmail
	}

	defer rows.Close()

	for rows.Next() {

		numberOfCopies++

	}

	if numberOfCopies == 1 {
		objDuplicateEmail.IsEmailValid = "true"
	} else {
		objDuplicateEmail.IsEmailValid = "false"
	}

	return objDuplicateEmail

}
