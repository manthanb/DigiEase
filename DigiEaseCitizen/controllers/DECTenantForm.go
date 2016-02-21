package controllers

func (c *MainController) DECTenantForm() {

	var objTenantFormRequest models.TenantFormRequest
	var objTenantFormResponse models.TenantFormResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal(&objTenantFormRequest)
	HandleError(err)

	objTenantFormResponse = models.DECTenantForm(objTenantFormRequest)

	c.Data["json"] = objTenantFormResponse

	c.ServeJson()

}
