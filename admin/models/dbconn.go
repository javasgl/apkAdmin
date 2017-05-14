package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("db::host")
	dbport := beego.AppConfig.String("db::port")
	dbuser := beego.AppConfig.String("db::user")
	dbpass := beego.AppConfig.String("db::password")
	dbname := beego.AppConfig.String("db::dbname")
	beego.Debug("dbname" + dbhost)
	beego.Debug("dbuser" + dbuser)
	beego.Debug("dbpass" + dbpass)
	beego.Debug("dbname" + dbname)
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)
	orm.RegisterDataBase("default", "mysql", dsn)

	beego.Debug(dsn)
	orm.RegisterModel(new(LoginUser))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
