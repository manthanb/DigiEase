package controllers

import "DigiEaseCitizen/models"
import "encoding/json"

func (c *MainController) DECUserRegistration() {

	var objUserRegistrationRequest models.UserRegistrationRequest
	var objUserLoginResponse models.UserLoginResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objUserRegistrationRequest)
	models.HandleError(err)

	objUserLoginResponse = models.DECUserRegistration(objUserRegistrationRequest)

	c.Data["json"] = objUserLoginResponse

	c.ServeJson()

}
