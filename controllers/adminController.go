package controllers

import "fmt"

type AdminController struct {
	BaseController
}

func (this *AdminController) Get(){
	fmt.Println(this.GetSession("session_user"))
	this.Layout = "admin/base/base.html"
	this.TplName = "admin/index/index.html"
}