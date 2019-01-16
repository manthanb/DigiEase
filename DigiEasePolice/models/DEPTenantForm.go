package models

import "strconv"

func DEPTenantForm(objGetFormRequest GetFormRequest) TenantFormResponse {

	var objTenantFormResponse TenantFormResponse

	objTenantFormResponse.ErrorMessage = "Undefined exception"

	db := GetDBConnection()
	defer db.Close()

	intFormID, err := strconv.ParseInt(objGetFormRequest.FormID, 10, 64)
	HandleError(err)

	query, err := db.Prepare(`select owner_first_name, owner_middle_name, owner_last_name, owner_age,
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
		owner_photo, agent_photo, tenant_photo, owner_photo_id, tenant_photo_id from tenant_form where form_id=?`)
	HandleError(err)

	err = query.QueryRow(intFormID).Scan(&objTenantFormResponse.Owner.FirstName, &objTenantFormResponse.Owner.MiddleName,
		&objTenantFormResponse.Owner.LastName, &objTenantFormResponse.Owner.Age, &objTenantFormResponse.Owner.Address1,
		&objTenantFormResponse.Owner.Address2, &objTenantFormResponse.Owner.Address3, &objTenantFormResponse.Owner.City,
		&objTenantFormResponse.Owner.State, &objTenantFormResponse.Owner.PinCode, &objTenantFormResponse.Owner.Mobile,
		&objTenantFormResponse.Owner.Landline,

		&objTenantFormResponse.Property.Address1, &objTenantFormResponse.Property.Address2,
		&objTenantFormResponse.Property.Address3, &objTenantFormResponse.Property.City, &objTenantFormResponse.Property.State,
		&objTenantFormResponse.Property.PinCode, &objTenantFormResponse.Property.Use,

		&objTenantFormResponse.Tenant.FirstName, &objTenantFormResponse.Tenant.MiddleName, &objTenantFormResponse.Tenant.LastName,
		&objTenantFormResponse.Tenant.Age, &objTenantFormResponse.Tenant.Address1, &objTenantFormResponse.Tenant.Address2,
		&objTenantFormResponse.Tenant.Address3, &objTenantFormResponse.Tenant.City, &objTenantFormResponse.Tenant.State,
		&objTenantFormResponse.Tenant.PinCode, &objTenantFormResponse.Tenant.Mobile, &objTenantFormResponse.Tenant.Landline,
		&objTenantFormResponse.Tenant.PermanentAddress, &objTenantFormResponse.Tenant.OfficeAddress,

		&objTenantFormResponse.Agent.FirstName, &objTenantFormResponse.Agent.MiddleName,
		&objTenantFormResponse.Agent.LastName, &objTenantFormResponse.Agent.Age, &objTenantFormResponse.Agent.Address1,
		&objTenantFormResponse.Agent.Address2, &objTenantFormResponse.Agent.Address3, &objTenantFormResponse.Agent.City,
		&objTenantFormResponse.Agent.State, &objTenantFormResponse.Agent.PinCode, &objTenantFormResponse.Agent.Mobile,
		&objTenantFormResponse.Agent.Landline,

		&objTenantFormResponse.OwnerPhoto, &objTenantFormResponse.AgentPhoto, &objTenantFormResponse.TenantPhoto,
		&objTenantFormResponse.OwnerPhotoID, &objTenantFormResponse.TenantPhotoID)

	HandleError(err)

	statusQuery, err := db.Prepare("select status from tenant_status where form_id=?")
	HandleError(err)

	err = statusQuery.QueryRow(intFormID).Scan(&objTenantFormResponse.Status)
	HandleError(err)

	objTenantFormResponse.ErrorMessage = "-"

	return objTenantFormResponse

}
