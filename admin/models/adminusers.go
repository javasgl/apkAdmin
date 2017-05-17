package models

import (
	"github.com/javasgl/apkAdmin/admin/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginUser struct {
	UserId   int    `orm:"PK"`
	Username string `orm:size(100)`
	Password string `orm:size(50)`
}

func init() {
}

func GetUserId(user LoginUser) {

	beego.Debug(user)
	o := orm.NewOrm()
	o.Insert(&user)
}
func Login(username, password string) bool {
	result := make(orm.Params)
	orm.NewOrm().Raw("SELECT username,password FROM login_user WHERE username=? AND password=?", username, utils.String2md5(password)).RowsToMap(&result, "username", "password")

	if _, ok := result[username]; !ok {
		return false
	}
	return true
}
