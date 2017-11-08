package routers

import (
	"sale/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.AccountController{})
    beego.Router("/admin", &controllers.AdminController{})
    beego.Router("/admin/user", &controllers.UserController{})
	beego.Router("/admin/user/get_user", &controllers.UserController{}, "post:GetUser")

}
