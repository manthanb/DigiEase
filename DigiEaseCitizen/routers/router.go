package routers

import (
	"DigiEaseCitizen/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/DECUserLogin", &controllers.MainController{}, "get,post:DECUserLogin")
	beego.Router("/DECUserRegistration", &controllers.MainController{}, "get,post:DECUserRegistration")
	beego.Router("/DECMyForms", &controllers.MainController{}, "get,post:DECMyForms")
	beego.Router("/EditAccount", &controllers.MainController{}, "get,post:EditAccount")
	beego.Router("/ViewAccount", &controllers.MainController{}, "get,post:ViewAccount")
	beego.Router("/ViewMyDocs", &controllers.MainController{}, "get,post:ViewMyDocs")
	beego.Router("/EditMyDocs", &controllers.MainController{}, "get,post:EditMyDocs")
	beego.Router("/TenantForm", &controllers.MainController{}, "get,post:TenantForm")
	beego.Router("/CheckDuplicateEmail", &controllers.MainController{}, "get,post:CheckDuplicateEmail")

}
