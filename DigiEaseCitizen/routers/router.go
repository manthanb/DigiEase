package routers

import "DigiEaseCitizen/controllers"
import "github.com/astaxie/beego"

func init() {
	beego.Router("/DECUserLogin", &controllers.MainController{}, "post:DECUserLogin")
	beego.Router("/DECUserRegistration", &controllers.MainController{}, "post:DECUserRegistration")
	beego.Router("/DECTenantForm", &controllers.MainController{}, "post:DECTenantForm")
}
