package controllers

import "DigiEaseCitizen/models"
import "encoding/json"

func (c *MainController) DECUserLogin() {

	var objUserLoginRequest models.UserLoginRequest
	var objUserLoginResponse models.UserLoginResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objUserLoginRequest)
	models.HandleError(err)

	objUserLoginResponse = models.DECUserLogin(objUserLoginRequest)

	c.Data["json"] = objUserLoginResponse
	c.ServeJson()

}
