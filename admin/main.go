package main

import (
	"github.com/astaxie/beego"
	"github.com/javasgl/apkAdmin/admin/models"
	_ "github.com/javasgl/apkAdmin/admin/routers"
)

func init() {
	models.Init()
}

func main() {
	beego.Run()
}
