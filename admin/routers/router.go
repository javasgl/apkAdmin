package routers

import (
	"github.com/javasgl/apkAdmin/admin/controllers"
	"github.com/javasgl/apkAdmin/admin/models"
	"github.com/javasgl/apkAdmin/admin/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	models.Init()
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.MainController{}, "post:DoLogin")
	beego.Router("/dashboard/packing", &controllers.PackingController{})
	beego.Router("/dashboard/dopacking", &controllers.PackingController{}, "post:DoPacking")

	var FilterUser = func(ctx *context.Context) {

		if utils.ValidateToken(ctx) {
			return
		}
		ctx.Redirect(302, "/")
	}
	beego.InsertFilter("/dashboard/*", beego.BeforeRouter, FilterUser)
}
