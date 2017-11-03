package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id uint
	Username string `orm:"size(60)"`
	Password string `orm:"size(60)"`
	Created_at int16
	Updated_at int16

}

func init() {
	orm.RegisterModel(new(User))
}