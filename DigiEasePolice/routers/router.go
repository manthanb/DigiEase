package routers

import (
	"DigiEasePolice/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/DEPUserLogin", &controllers.MainController{}, "get,post:DEPUserLogin")
	beego.Router("/DEPFormThumbnails", &controllers.MainController{}, "get,post:DEPFormThumbnails")
	beego.Router("/DEPGetForm", &controllers.MainController{}, "get,post:DEPGetForm")
	beego.Router("/DEPUpdateStatus", &controllers.MainController{}, "get,post:DEPUpdateStatus")
}
