package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	webName,_ := beego.GetConfig("String","webname"," ")
	this.Data["webName"] = webName
	sUser := this.GetSession("session_user")
	if sUser == nil {
		this.Ctx.Redirect(302,"/")
		return
	}


}
