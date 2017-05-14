package routers

import (
	"admin/controllers"
	"admin/models"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.MainController{}, "post:DoLogin")

	models.Init()
}
