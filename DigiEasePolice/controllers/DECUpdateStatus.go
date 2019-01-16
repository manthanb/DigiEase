package controllers

import "DigiEasePolice/models"
import "encoding/json"
import _ "github.com/go-sql-driver/mysql"

func (c *MainController) DEPUpdateStatus() {

	var objUpdateStatusRequest models.UpdateStatusRequest
	var objUpdateStatusResponse models.UpdateStatusResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objUpdateStatusRequest)
	models.HandleError(err)

	objUpdateStatusResponse = models.DEPUpdateStatus(objUpdateStatusRequest)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = objUpdateStatusResponse
	c.ServeJson()

}
