package controllers

import (
	"DigiEaseCitizen/models"
	"encoding/json"
)

func (c *MainController) CheckDuplicateEmail() {

	var isEmailValid models.DuplicateEmail
	var objUserEmailStruct models.UserEmailStruct

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objUserEmailStruct)

	models.HandleError(err)

	isEmailValid = models.CheckDuplicateEmail(objUserEmailStruct)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = isEmailValid
	c.ServeJson()
}
