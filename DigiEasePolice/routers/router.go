package routers

import (
	"DigiEasePolice/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/DEPUserLogin", &controllers.MainController{}, "get,post:DEPUserLogin")
	beego.Router("/DEPFormThumbnails", &controllers.MainController{}, "get,post:DEPFormThumbnails")
}
