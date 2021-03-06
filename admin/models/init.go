package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Hostname     string
	SessionKey   string
	RegisterMail string
	PasswordSalt string
)

func Init() {
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	initCustomConfig()

	initDBConn()

}

func initCustomConfig() {

	Hostname = beego.AppConfig.String("apk::hostname")

	SessionKey = beego.AppConfig.String("apk::sessionKey")

	RegisterMail = beego.AppConfig.String("apk::registerMail")

	PasswordSalt = beego.AppConfig.String("apk::passwordSalt")
}

func initDBConn() {

	dbhost := beego.AppConfig.String("db::host")
	dbport := beego.AppConfig.String("db::port")
	dbuser := beego.AppConfig.String("db::user")
	dbpass := beego.AppConfig.String("db::password")
	dbname := beego.AppConfig.String("db::dbname")
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)
	orm.RegisterDataBase("default", "mysql", dsn)

}
