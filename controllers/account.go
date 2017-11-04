package controllers

import (
	"fmt"
	"sale/models"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
)

type AccountController struct {
	BaseController
}

func (this *AccountController) Get(){
	this.TplName = "admin/account/login.html"
}

func (this *AccountController) Post() {
	o := orm.NewOrm()
	user := models.User{Id: 1}

	err := o.Read(&user)
	h := md5.New()
	h.Write([]byte("123456"))
	str := h.Sum(nil)

	fmt.Println(user.Username,user.Id,err,hex.EncodeToString(str))

	return;
}