package models

import "strconv"
import "database/sql"

func DEPFormThumbnails(objFormThumbnailsRequest FormThumbnailsRequest) FormThumbnailsResponse {

	var objFormThumbnailsResponse FormThumbnailsResponse
	var objFormThumbnail FormThumbnail

	objFormThumbnailsResponse.ExceptionMessage = "Undefined exception"

	db := GetDBConnection()
	defer db.Close()

	intAreaCode, err := strconv.ParseInt(objFormThumbnailsRequest.AreaCode, 10, 64)
	HandleError(err)

	if err != nil {
		objFormThumbnailsResponse.ExceptionMessage = "Area Code was not in proper format"
		return objFormThumbnailsResponse
	}

	var query *sql.Stmt
	var thumbnails *sql.Rows

	if intAreaCode != 1000 {

		query, err = db.Prepare(`select form_id, form_type, status, date_time, area,
			user_id, first_name, middle_name, last_name
			from tenant_status
			inner join users
			where tenant_status.user_email = users.email and tenant_status.area = ?`)
		HandleError(err)

		if err != nil {
			objFormThumbnailsResponse.ExceptionMessage = "Could not prepare select query - Local"
			return objFormThumbnailsResponse
		}

		thumbnails, err = query.Query(&intAreaCode)
		HandleError(err)

		if err != nil {
			objFormThumbnailsResponse.ExceptionMessage = "Could not execute select query - Local"
			return objFormThumbnailsResponse
		}

	} else {

		query, err = db.Prepare(`select form_id, form_type, status, date_time, area,
			user_id, first_name, middle_name, last_name
			from tenant_status
			inner join users
			where tenant_status.user_email = users.email`)
		HandleError(err)

		if err != nil {
			objFormThumbnailsResponse.ExceptionMessage = "Could not prepare select query - Commissioner"
			return objFormThumbnailsResponse
		}

		thumbnails, err = query.Query()
		HandleError(err)

		if err != nil {
			objFormThumbnailsResponse.ExceptionMessage = "Could not execute select query - Commissioner"
			return objFormThumbnailsResponse
		}

	}

	for thumbnails.Next() {

		err = thumbnails.Scan(&objFormThumbnail.FormID, &objFormThumbnail.FormType, &objFormThumbnail.Status,
			&objFormThumbnail.DateTime, &objFormThumbnail.Area, &objFormThumbnail.UserID, &objFormThumbnail.FirstName,
			&objFormThumbnail.MiddleName, &objFormThumbnail.LastName)
		HandleError(err)

		if err != nil {
			objFormThumbnailsResponse.ExceptionMessage = "Error reading data"
			return objFormThumbnailsResponse
		}

		objFormThumbnail.ApplicationID = GetTimeStampString(objFormThumbnail.DateTime)

		objFormThumbnailsResponse.FormThumbnails = append(objFormThumbnailsResponse.FormThumbnails, objFormThumbnail)

	}

	objFormThumbnailsResponse.ExceptionMessage = "-"

	return objFormThumbnailsResponse

}
