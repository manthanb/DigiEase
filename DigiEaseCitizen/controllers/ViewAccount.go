package controllers

import (
	"DigiEaseCitizen/models"
	"encoding/json"
)

func (c *MainController) ViewAccount() {
	var accountDetails models.EditAccountRequest
	var objUserEmailStruct models.UserEmailStruct

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objUserEmailStruct)

	models.HandleError(err)

	accountDetails = models.ViewAccount(objUserEmailStruct)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = accountDetails
	c.ServeJson()
}
