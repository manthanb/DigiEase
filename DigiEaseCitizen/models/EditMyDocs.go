package models

func EditMyDocs(objMyDocsRequest MyDocsRequest) MyDocsResponse {
	var objMyDocsResponse MyDocsResponse
	var currentData MyDocsRequest

	objMyDocsResponse.ExceptionMessage = "Undefined exception - models/EditMyDocs"
	objMyDocsResponse.Status = "false"

	db := GetDBConnection()

	defer db.Close()

	updateStatement, err := db.Prepare("update users set aadhar_card = ?, pan_card = ?, passport = ?, driving_license = ?, photo = ?, address_proof = ? where email = ?")

	HandleError(err)

	if err != nil {
		objMyDocsResponse.ExceptionMessage = "Could not prepare update statement - models/EditMyDocs"

		return objMyDocsResponse
	}

	err = db.QueryRow("select aadhar_card, pan_card, passport, driving_license, photo, address_proof from users where email = ?",
		objMyDocsRequest.UserEmail).Scan(&currentData.AadharCard, &currentData.PanCard, &currentData.Passport, &currentData.DrivingLicense, &currentData.Photo, &currentData.AddressProof)

	HandleError(err)

	if err != nil {
		objMyDocsResponse.ExceptionMessage = "Could not get existing data - models/EditMyDocs"

		return objMyDocsResponse
	}

	if objMyDocsRequest.AadharCard == "" {
		objMyDocsRequest.AadharCard = currentData.AadharCard
	}
	if objMyDocsRequest.PanCard == "" {
		objMyDocsRequest.PanCard = currentData.PanCard
	}
	if objMyDocsRequest.Passport == "" {
		objMyDocsRequest.Passport = currentData.Passport
	}
	if objMyDocsRequest.DrivingLicense == "" {
		objMyDocsRequest.DrivingLicense = currentData.DrivingLicense
	}
	if objMyDocsRequest.Photo == "" {
		objMyDocsRequest.Photo = currentData.Photo
	}
	if objMyDocsRequest.AddressProof == "" {
		objMyDocsRequest.AddressProof = currentData.AddressProof
	}

	_, err = updateStatement.Exec(objMyDocsRequest.AadharCard, objMyDocsRequest.PanCard, objMyDocsRequest.Passport, objMyDocsRequest.DrivingLicense, objMyDocsRequest.Photo, objMyDocsRequest.AddressProof, objMyDocsRequest.UserEmail)

	HandleError(err)

	if err != nil {
		objMyDocsResponse.ExceptionMessage = "Could not execute update statement - models/EditMyDocs"

		return objMyDocsResponse
	}

	objMyDocsResponse.Status = "true"
	objMyDocsResponse.ExceptionMessage = "Everything went okay. - models/EditMyDocs"

	return objMyDocsResponse

}
