package controllers

import (
	"DigiEaseCitizen/models"
	"encoding/json"
)

func (c *MainController) EditMyDocs() {
	var objMyDocsRequest models.MyDocsRequest
	var objMyDocsResponse models.MyDocsResponse
	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objMyDocsRequest)

	models.HandleError(err)

	objMyDocsResponse = models.EditMyDocs(objMyDocsRequest)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = objMyDocsResponse
	c.ServeJson()
}
