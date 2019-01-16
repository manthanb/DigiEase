package controllers

import "DigiEaseCitizen/models"
import "encoding/json"

func (c *MainController) TenantForm() {

	var objTenantFormRequest models.TenantFormRequest
	var objTenantFormResponse models.TenantFormResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objTenantFormRequest)

	models.HandleError(err)

	objTenantFormResponse = models.TenantForm(objTenantFormRequest)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = objTenantFormResponse
	c.ServeJson()

}
