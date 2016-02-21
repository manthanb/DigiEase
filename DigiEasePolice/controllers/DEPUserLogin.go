package controllers

import "encoding/json"
import "DigiEasePolice/models"

func (c *MainController) DEPUserLogin() {

	var objUserLoginRequest models.UserLoginRequest
	var objUserLoginResponse models.UserLoginResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objUserLoginRequest)
	models.HandleError(err)

	objUserLoginResponse = models.DEPUserLogin(objUserLoginRequest)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = objUserLoginResponse

	c.ServeJson()

}
