package models

func ViewMyDocs(objUserEmailStruct UserEmailStruct) MyDocsRequest {

	var objMyDocsRequest MyDocsRequest

	db := GetDBConnection()

	defer db.Close()

	err := db.QueryRow("select aadhar_card, pan_card, passport, driving_license, photo, address_proof from users where email = ?", objUserEmailStruct.UserEmail).Scan(&objMyDocsRequest.AadharCard, &objMyDocsRequest.PanCard, &objMyDocsRequest.Passport, &objMyDocsRequest.DrivingLicense, &objMyDocsRequest.Photo, &objMyDocsRequest.AddressProof)

	HandleError(err)

	if err != nil {

		return objMyDocsRequest
	}

	return objMyDocsRequest

}
