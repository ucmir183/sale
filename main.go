package main

import (
	_ "sale/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"

	_"sale/models"
	"github.com/astaxie/beego/session"
)

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}

	globalSessions, _ := session.NewManager("memory",sessionConfig)
	go globalSessions.GC()

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:as112233@/orm_test?charset=utf8")
}

func main() {
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400
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

