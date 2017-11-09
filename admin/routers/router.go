package routers

import (
	"github.com/javasgl/apkAdmin/admin/controllers"
	"github.com/javasgl/apkAdmin/admin/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{}, "get:Register")
	beego.Router("/api/login", &controllers.MainController{}, "post:DoLogin")
	beego.Router("/api/register", &controllers.MainController{}, "post:DoRegister")

	beego.Router("/api/upload", &controllers.UploadController{}, "post:DoUpload")

	beego.Router("/dashboard/loginout", &controllers.MainController{}, "post:Loginout")

	beego.Router("/dashboard/packing", &controllers.PackingController{})
	beego.Router("/dashboard/addJob", &controllers.PackingController{}, "post:AddJob")
	beego.Router("/dashboard/getJobs", &controllers.PackingController{}, "get:GetJobs")

	beego.Router("/dashboard/channel", &controllers.ChannelController{})
	beego.Router("/dashboard/addChannel", &controllers.ChannelController{}, "post:AddChannel")
	beego.Router("/dashboard/getChannels", &controllers.ChannelController{}, "get:GetChannels")
	beego.Router("/dashboard/delChannel", &controllers.ChannelController{}, "post:DelChannel")

	var FilterUser = func(ctx *context.Context) {

		if models.ValidateToken(ctx) {
			return
		}
		ctx.Redirect(302, "/")
	}
	beego.InsertFilter("/dashboard/*", beego.BeforeRouter, FilterUser)
}
