package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	TABLE_USER = "apk_users"
)

type User struct {
	UserId        int    `orm:"PK"`
	Username      string `orm:"size(100)"`
	Password      string `orm:"size(50)"`
	Email         string `orm:"size(50)"`
	Validated     byte
	ValidateToken string
	RegisterTime  int64
}

func (c *User) TableName() string {
	return TABLE_USER
}

func init() {
	orm.RegisterModel(new(User))
}

func Register(user User) bool {
	var count int
	orm.NewOrm().Raw(fmt.Sprintf("SELECT COUNT(1) AS count FROM %s WHERE username=?", TABLE_USER), user.Username).QueryRow(&count)
	if count > 0 {
		return false
	}
	user.Email = user.Username + "@" + RegisterMail
	user.Password = String2md5(user.Password)
	user.Validated = 0
	user.ValidateToken = ""
	user.RegisterTime = time.Now().Unix()

	beego.Debug(user)
	_, err := orm.NewOrm().Insert(&user)

	if err != nil {
		return false
	}
	return true
}

func GetUserId(user User) {

	beego.Debug(user)
	o := orm.NewOrm()
	o.Insert(&user)
}

func Login(user User) (res User, err error) {
	err = orm.NewOrm().Raw(fmt.Sprintf("SELECT user_id,username,email,validated FROM %s WHERE username=? AND password=?", TABLE_USER), user.Username, String2md5(user.Password)).QueryRow(&res)
	return res, err
}
