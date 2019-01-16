package models

import "database/sql"
import "strconv"

func DEPUpdateStatus(objUpdateStatusRequest UpdateStatusRequest) UpdateStatusResponse {

	var objUpdateStatusResponse UpdateStatusResponse

	objUpdateStatusResponse.Status = "false"
	objUpdateStatusResponse.ExceptionMessage = "Undefined exception"

	db := GetDBConnection()
	defer db.Close()

	intStatus := objUpdateStatusRequest.Status

	intFormID, err := strconv.ParseInt(objUpdateStatusRequest.FormID, 10, 64)
	HandleError(err)

	if err != nil {
		objUpdateStatusResponse.Status = "false"
		objUpdateStatusResponse.ExceptionMessage = "form id not in proper format"
	}

	var updateStatement *sql.Stmt

	if objUpdateStatusRequest.Discard == "true" && objUpdateStatusRequest.AreaCodeStatus == "true" {

		intAreaCode, err := strconv.ParseInt(objUpdateStatusRequest.AreaCode, 10, 64)
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "area code not in proper format - discard, area code"
		}

		updateStatement, err = db.Prepare("update tenant_status set status=?, comment=?, area=? where form_id=?")
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "could not prepare update statement - discard, area code"
		}

		_, err = updateStatement.Exec(intStatus, objUpdateStatusRequest.Comment, intAreaCode, intFormID)
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "could not execute update statement - discard, area code"
		}

	} else if objUpdateStatusRequest.Discard == "true" && objUpdateStatusRequest.AreaCodeStatus == "false" {

		updateStatement, err = db.Prepare("update tenant_status set status=?, comment=? where form_id=?")
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "could not prepare update statement - discard, _"
		}

		_, err = updateStatement.Exec(intStatus, objUpdateStatusRequest.Comment, intFormID)
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "could not execute update statement - discard, _"
		}

	} else if objUpdateStatusRequest.Discard == "false" && objUpdateStatusRequest.AreaCodeStatus == "true" {

		intAreaCode, err := strconv.ParseInt(objUpdateStatusRequest.AreaCode, 10, 64)
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "area code not in proper format - _, area code"
		}

		updateStatement, err = db.Prepare("update tenant_status set status=?, area=? where form_id=?")
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "could not prepare update statement - _, area code"
		}

		_, err = updateStatement.Exec(intStatus, intAreaCode, intFormID)
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "could not execute update statement - non discard"
		}

	} else if objUpdateStatusRequest.Discard == "false" && objUpdateStatusRequest.AreaCodeStatus == "false" {

		updateStatement, err = db.Prepare("update tenant_status set status=? where form_id=?")
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "could not prepare update statement - _, _"
		}

		_, err = updateStatement.Exec(intStatus, intFormID)
		HandleError(err)

		if err != nil {
			objUpdateStatusResponse.Status = "false"
			objUpdateStatusResponse.ExceptionMessage = "could not execute update statement - _, _"
		}

	}

	objUpdateStatusResponse.Status = "true"
	objUpdateStatusResponse.ExceptionMessage = "-"

	return objUpdateStatusResponse

}
