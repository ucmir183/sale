package routers

import (
	"sale/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.AccountController{})

}
