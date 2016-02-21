package controllers

import "DigiEasePolice/models"
import "encoding/json"

func (c *MainController) DEPFormThumbnails() {

	var objFormThumbnailsRequest models.FormThumbnailsRequest
	var objFormThumbnailsResponse models.FormThumbnailsResponse

	strRequestJson := c.Ctx.Input.Query("json")

	err := json.Unmarshal([]byte(strRequestJson), &objFormThumbnailsRequest)
	models.HandleError(err)

	objFormThumbnailsResponse = models.DEPFormThumbnails(objFormThumbnailsRequest)

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	c.Data["json"] = objFormThumbnailsResponse

	c.ServeJson()

}
