package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	SessionKey   string
	RegisterMail string
)

func Init() {
	beego.Debug("初始化.....")

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	initCustomConfig()

	initDBConn()

}

func initCustomConfig() {

	SessionKey = beego.AppConfig.String("apk::sessionKey")

	RegisterMail = beego.AppConfig.String("apk::registerMail")

}

func initDBConn() {

	beego.Debug("初始化-数据库")
	dbhost := beego.AppConfig.String("db::host")
	dbport := beego.AppConfig.String("db::port")
	dbuser := beego.AppConfig.String("db::user")
	dbpass := beego.AppConfig.String("db::password")
	dbname := beego.AppConfig.String("db::dbname")
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)
	orm.RegisterDataBase("default", "mysql", dsn)

}
