package models

import (
	"errors"
	"fmt"

	"github.com/javasgl/apkAdmin/admin/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	TABLE_USER = "apk_users"
)

type User struct {
	UserId   int    `orm:"PK"`
	Username string `orm:"size(100)"`
	Password string `orm:"size(50)"`
}

func (c *User) TableName() string {
	return TABLE_USER
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserId(user User) {

	beego.Debug(user)
	o := orm.NewOrm()
	o.Insert(&user)
}

func Login(user User) (res orm.Params, err error) {
	result := make(orm.Params)
	orm.NewOrm().Raw(fmt.Sprintf("SELECT user_id,username FROM %s WHERE username=? AND password=?", TABLE_USER), user.Username, utils.String2md5(user.Password)).RowsToMap(&result, "username", "user_id")

	if _, ok := result[user.Username]; !ok {
		return nil, errors.New("user not exists")
	}
	return result, nil
}
