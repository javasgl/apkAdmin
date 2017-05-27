package models

import (
	"errors"

	"github.com/javasgl/apkAdmin/admin/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginUser struct {
	UserId   int    `orm:"PK"`
	Username string `orm:"size(100)"`
	Password string `orm:"size(50)"`
}

func init() {
	orm.RegisterModel(new(LoginUser))
}

func GetUserId(user LoginUser) {

	beego.Debug(user)
	o := orm.NewOrm()
	o.Insert(&user)
}

func Login(user LoginUser) (res orm.Params, err error) {
	result := make(orm.Params)
	orm.NewOrm().Raw("SELECT user_id,username FROM login_user WHERE username=? AND password=?", user.Username, utils.String2md5(user.Password)).RowsToMap(&result, "username", "user_id")

	if _, ok := result[user.Username]; !ok {
		return nil, errors.New("user not exists")
	}
	return result, nil
}
