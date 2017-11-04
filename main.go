package main

import (
	_ "sale/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"

	_"sale/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:as112233@/orm_test?charset=utf8")
}

func main() {
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name,force,verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.SetStaticPath("/plugin", "static/js/plugin")//注册插件静态文件目录

	beego.Run()
}

