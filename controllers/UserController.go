package controllers

import (
	"github.com/astaxie/beego/orm"
	"sale/models"
	"fmt"
	"encoding/json"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	this.TplName = "admin/user/index.html"
	this.Data["is_user"] = true;
}


func (this * UserController) GetUser() {
	o := orm.NewOrm()
	var users []models.User
	//user := new(models.User)
	qs := o.QueryTable("user")
	res,_ := qs.All(&users);

	fmt.Println(json.Marshal(users),res);


}