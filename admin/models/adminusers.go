package models

import (
	"sync"

	"admin/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	UserMutex      sync.RWMutex
	UserSessionMap map[string]string
)

type LoginUser struct {
	UserId   int    `orm:"PK"`
	Username string `orm:size(100)`
	Password string `orm:size(50)`
}

func init() {
	UserSessionMap = make(map[string]string)
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

	UserMutex.Lock()
	UserSessionMap[username] = password
	UserMutex.Unlock()
	return true
}
func IsLogined(userId, session string) bool {
	if session == "" {
		return false
	}
	UserMutex.RLock()
	defer UserMutex.RUnlock()
	return UserSessionMap[userId] == session
}
