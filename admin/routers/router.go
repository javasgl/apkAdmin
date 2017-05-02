package routers

import (
	"admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.MainController{}, "post:DoLogin")
}
