package controllers

import (
	"sale/models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/lib"
	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

func (this *AccountController) Get(){
	this.TplName = "admin/account/login.html"
}

func (this *AccountController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	o := orm.NewOrm()
	user := models.User{Username:username}
	err := o.Read(&user, "username")
	if err != nil {
		fmt.Println(err)
	}

	if  user.Password == lib.Md5_encode(password)  {
		user.Password = ""
		this.SetSession("session_user",user)

		this.Ctx.Redirect(302, "/admin")
	}



	return;
}