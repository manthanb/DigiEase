package controllers

import (
	"DigiEaseCitizen/models"
	"encoding/json"
)

func (c *MainController) EditAccount() {
	var objEditAccountRequest models.EditAccountRequest
	var objEditAccountResponse models.EditAccountResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objEditAccountRequest)

	models.HandleError(err)

	objEditAccountResponse = models.EditAccount(objEditAccountRequest)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = objEditAccountResponse
	c.ServeJson()
}
