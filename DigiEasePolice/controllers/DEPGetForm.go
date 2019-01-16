package controllers

import "DigiEasePolice/models"
import "encoding/json"
import _ "github.com/go-sql-driver/mysql"
import "strconv"

func (c *MainController) DEPGetForm() {

	var objGetFormRequest models.GetFormRequest

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objGetFormRequest)
	models.HandleError(err)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	intFormType, err := strconv.ParseInt(objGetFormRequest.FormType, 10, 64)
	models.HandleError(err)

	switch intFormType {

	case 1:
		var objTenantFormResponse models.TenantFormResponse
		objTenantFormResponse = models.DEPTenantForm(objGetFormRequest)
		c.Data["json"] = objTenantFormResponse
		c.ServeJson()

	}

}
