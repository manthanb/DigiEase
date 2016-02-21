package models

import "strconv"

func DEPFormThumbnails(objFormThumbnailsRequest FormThumbnailsRequest) FormThumbnailsResponse {

	var objFormThumbnailsResponse FormThumbnailsResponse
	var objFormThumbnail FormThumbnail

	objFormThumbnailsResponse.ExceptionMessage = "Undefined exception"

	db := GetDBConnection()

	intAreaCode, err := strconv.ParseInt(objFormThumbnailsRequest.AreaCode, 10, 64)
	HandleError(err)

	query, err := db.Prepare(`select form_id, form_type, status, date_time, 
		user_id, first_name, middle_name, last_name 
		from  tenant_status 
		inner join users 
		where tenant_status.user_email = users.email and tenant_status.area = ?`)
	HandleError(err)

	if err != nil {
		objFormThumbnailsResponse.ExceptionMessage = "Could not prepare select query"
		return objFormThumbnailsResponse
	}

	rows, err := query.Query(intAreaCode)
	HandleError(err)

	if err != nil {
		objFormThumbnailsResponse.ExceptionMessage = "Could not execute select query"
		return objFormThumbnailsResponse
	}

	for rows.Next() {

		err = rows.Scan(&objFormThumbnail.FormID, &objFormThumbnail.FormType, &objFormThumbnail.Status,
			&objFormThumbnail.DateTime, &objFormThumbnail.UserID, &objFormThumbnail.FirstName,
			&objFormThumbnail.MiddleName, &objFormThumbnail.LastName)
		HandleError(err)

		if err != nil {
			objFormThumbnailsResponse.ExceptionMessage = "Propblem parsing data"
			return objFormThumbnailsResponse
		}

		objFormThumbnailsResponse.FormThumbnails = append(objFormThumbnailsResponse.FormThumbnails, objFormThumbnail)

	}

	objFormThumbnailsResponse.ExceptionMessage = "-"

	return objFormThumbnailsResponse

}
