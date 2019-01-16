package models

func DECMyForms(objUserEmailStruct UserEmailStruct) MyFormsResponse {

	var objMyFormsResponse MyFormsResponse
	var objMyForms MyForms

	objMyFormsResponse.Status = "false"
	objMyFormsResponse.ExceptionMessage = "Undefined exception - models/DECMyForms"

	db := GetDBConnection()

	defer db.Close()

	rows, err := db.Query("select form_id, form_type, status, date_time, area, comment, area_name from tenant_status, area where user_email = ? and tenant_status.area = area.area_id", objUserEmailStruct.UserEmail)

	HandleError(err)

	if err != nil {

		objMyFormsResponse.Status = "false"
		objMyFormsResponse.ExceptionMessage = "Could not execute query - models/DECMyForms"

		return objMyFormsResponse
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&objMyForms.FormID, &objMyForms.FormType, &objMyForms.Status, &objMyForms.DateAndTime, &objMyForms.AreaID, &objMyForms.Comment, &objMyForms.AreaName)

		HandleError(err)

		if err != nil {

			objMyFormsResponse.Status = "false"
			objMyFormsResponse.ExceptionMessage = "Could not scan rows - models/DECMyForms"

			return objMyFormsResponse
		}

		if objMyForms.FormType == 1 {
			objMyForms.FormName = "Tenant"
		}
		if objMyForms.FormType == 2 {
			objMyForms.FormName = "Police Clearance"
		}

		objMyFormsResponse.Status = "true"
		objMyFormsResponse.ExceptionMessage = "Everything went okay. - models/DECMyForms"

		objMyFormsResponse.Forms = append(objMyFormsResponse.Forms, objMyForms)

	}

	err = rows.Err()
	HandleError(err)

	if err != nil {

		objMyFormsResponse.Status = "false"
		objMyFormsResponse.ExceptionMessage = "Error after iterating through rows - models/DECMyForms"

		return objMyFormsResponse
	}

	return objMyFormsResponse

}
