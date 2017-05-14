package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginUser struct {
	UserId   int    `orm:"PK"`
	Username string `orm:size(100)`
	Password string `orm:size(50)`
}

func GetUserId(user LoginUser) {

	beego.Debug(user)
	o := orm.NewOrm()
	o.Insert(&user)
}
