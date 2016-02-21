package models

func DECTenantForm(objTenantFormRequest TenantFormRequest) TenantFormResponse {

	var objTenantFormResponse TenantFormResponse

	objTenantFormResponse.Status = "false"
	objTenantFormResponse.ExceptionMessage = "Default exception - models/DECTenantForm"

	db := GetDBConnection()
	defer db.Close()

	insertFormStatement, err := db.Prepare(`insert into 
		tenant_form(owner_first_name, owner_middle_name, 
		owner_last_name, owner_age, owner_address1, owner_address2, owner_address3, owner_city, owner_state,
		owner_pincode, owner_mobile, owner_landline, property_address1, property_address2, property_address3,
		property_city, property_state, property_pincode, property_use, tenant_first_name, tenant_middle_name,
		tenant_last_name, tenant_age, tenant_address1, tenant_address2, tenant_address3, tenant_city, 
		tenant_state, tenant_pincode, tenant_mobile, tenant_landline, tenant_permanent_address, 
		tenant_office_address, agent_first_name, agent_middle_name, agent_last_name, agent_age, 
		agent_address1, agent_address2, agent_address3, agent_city, agent_state, agent_pincode, agent_mobile, 
		agent_landline 
		values(objTenantFormRequest.Owner.FirstName, objTenantFormRequest.Owner.MiddleName, 
		objTenantFormRequest.Owner.`)

}
