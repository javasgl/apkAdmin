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
	beego.Router("/dashboard/addJob", &controllers.PackingController{}, "post:AddJob")
	beego.Router("/dashboard/getJobs", &controllers.PackingController{}, "get:GetJobs")

	beego.Router("/dashboard/channel", &controllers.ChannelController{})
	beego.Router("/dashboard/addChannel", &controllers.ChannelController{}, "post:AddChannel")
	beego.Router("/dashboard/getChannels", &controllers.ChannelController{}, "get:GetChannels")

	var FilterUser = func(ctx *context.Context) {

		if utils.ValidateToken(ctx) {
			return
		}
		ctx.Redirect(302, "/")
	}
	beego.InsertFilter("/dashboard/*", beego.BeforeRouter, FilterUser)
}
