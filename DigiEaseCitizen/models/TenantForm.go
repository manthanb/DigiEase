package models

import "strconv"
import "fmt"

func TenantForm(objTenantFormRequest TenantFormRequest) TenantFormResponse {

	var objTenantFormResponse TenantFormResponse

	fmt.Println(objTenantFormRequest)

	objTenantFormResponse.Status = "false"
	objTenantFormResponse.ExceptionMessage = "Undefined exception - models/TenantForm"

	db := GetDBConnection()

	defer db.Close()

	insertTenant, err := db.Prepare(`insert into tenant_form(owner_first_name, owner_middle_name, owner_last_name, owner_age,
		owner_address1, owner_address2, owner_address3, owner_city, owner_state, owner_pincode, owner_mobile, 
		owner_landline, 
		property_address1, property_address2, property_address3, property_city, property_state, 
		property_pincode, property_use, 
		tenant_first_name, tenant_middle_name, tenant_last_name, tenant_age,
		tenant_address1, tenant_address2, tenant_address3, tenant_city, tenant_state, tenant_pincode, tenant_mobile,
		tenant_landline, tenant_permanent_address, tenant_office_address, 
		agent_first_name, agent_middle_name,
		agent_last_name, agent_age, agent_address1, agent_address2, agent_address3, agent_city, 
		agent_state, agent_pincode, agent_mobile, agent_landline,
		owner_photo, agent_photo, tenant_photo, owner_photo_id, tenant_photo_id) 
		values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)

	HandleError(err)

	if err != nil {
		objTenantFormResponse.Status = "false"
		objTenantFormResponse.ExceptionMessage = "Could not prepare insert tenant statement - models/TenantForm"

		return objTenantFormResponse
	}

	intOwnerAge, err := strconv.ParseInt(objTenantFormRequest.Owner.Age, 10, 64)

	HandleError(err)

	intTenantAge, err := strconv.ParseInt(objTenantFormRequest.Tenant.Age, 10, 64)

	HandleError(err)

	intAgentAge, err := strconv.ParseInt(objTenantFormRequest.Agent.Age, 10, 64)

	HandleError(err)

	res, err := insertTenant.Exec(objTenantFormRequest.Owner.FirstName, objTenantFormRequest.Owner.MiddleName,
		objTenantFormRequest.Owner.LastName, intOwnerAge, objTenantFormRequest.Owner.Address1,
		objTenantFormRequest.Owner.Address2, objTenantFormRequest.Owner.Address3, objTenantFormRequest.Owner.City,
		objTenantFormRequest.Owner.State, objTenantFormRequest.Owner.PinCode, objTenantFormRequest.Owner.Mobile,
		objTenantFormRequest.Owner.Landline,

		objTenantFormRequest.Property.Address1, objTenantFormRequest.Property.Address2,
		objTenantFormRequest.Property.Address3, objTenantFormRequest.Property.City, objTenantFormRequest.Property.State,
		objTenantFormRequest.Property.PinCode, objTenantFormRequest.Property.Use,

		objTenantFormRequest.Tenant.FirstName, objTenantFormRequest.Tenant.MiddleName, objTenantFormRequest.Tenant.LastName,
		intTenantAge, objTenantFormRequest.Tenant.Address1, objTenantFormRequest.Tenant.Address2,
		objTenantFormRequest.Tenant.Address3, objTenantFormRequest.Tenant.City, objTenantFormRequest.Tenant.State,
		objTenantFormRequest.Tenant.PinCode, objTenantFormRequest.Tenant.Mobile, objTenantFormRequest.Tenant.Landline,
		objTenantFormRequest.Tenant.PermanentAddress, objTenantFormRequest.Tenant.OfficeAddress,

		objTenantFormRequest.Agent.FirstName, objTenantFormRequest.Agent.MiddleName,
		objTenantFormRequest.Agent.LastName, intAgentAge, objTenantFormRequest.Agent.Address1,
		objTenantFormRequest.Agent.Address2, objTenantFormRequest.Agent.Address3, objTenantFormRequest.Agent.City,
		objTenantFormRequest.Agent.State, objTenantFormRequest.Agent.PinCode, objTenantFormRequest.Agent.Mobile,
		objTenantFormRequest.Agent.Landline,

		objTenantFormRequest.OwnerPhoto, objTenantFormRequest.AgentPhoto, objTenantFormRequest.TenantPhoto,
		objTenantFormRequest.OwnerPhotoID, objTenantFormRequest.TenantPhotoID)

	HandleError(err)

	if err != nil {
		objTenantFormResponse.Status = "false"
		objTenantFormResponse.ExceptionMessage = "Could not execute insert tenant statement - models/TenantForm"

		return objTenantFormResponse
	}

	intFormID, err := res.LastInsertId()
	HandleError(err)

	insertTenantStatus, err := db.Prepare(`insert into tenant_status(user_email,form_id,form_type,status,date_time) 
		values(?,?,?,?,?)`)

	HandleError(err)

	if err != nil {
		objTenantFormResponse.Status = "false"
		objTenantFormResponse.ExceptionMessage = "Could not prepare insert tenant status statement - models/TenantForm"

		return objTenantFormResponse
	}

	_, err = insertTenantStatus.Exec(objTenantFormRequest.UserEmail, intFormID, 1, 0, GetCurrentTimeStamp())

	HandleError(err)

	if err != nil {
		objTenantFormResponse.Status = "false"
		objTenantFormResponse.ExceptionMessage = "Could not execute insert tenant status statement - models/TenantForm"

		return objTenantFormResponse
	}

	objTenantFormResponse.ExceptionMessage = "Everything went okay. - models/TenantForm"
	objTenantFormResponse.Status = "true"

	return objTenantFormResponse

}
