package controllers

import (
	"DigiEaseCitizen/models"
	"encoding/json"
)

func (c *MainController) DECUserRegistration() {
	var objUserRegistrationRequest models.UserRegistrationRequest
	var objUserLoginResponse models.UserLoginResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objUserRegistrationRequest)

	models.HandleError(err)

	objUserLoginResponse = models.DECUserRegistration(objUserRegistrationRequest)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = objUserLoginResponse
	c.ServeJson()
}
