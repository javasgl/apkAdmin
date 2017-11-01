package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/javasgl/apkAdmin/admin/models"
)

type ChannelController struct {
	beego.Controller
}

func (this *ChannelController) Get() {

	this.Data["loginUser"] = this.GetSession(models.SessionKey)
	this.Layout = "layout.tpl"
	this.TplName = "channel/channel.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHeader"] = "channel/header.tpl"
	this.LayoutSections["HtmlBody"] = "channel/channel.tpl"
	this.LayoutSections["HtmlScripts"] = "channel/scripts.tpl"
}

func (this *ChannelController) AddChannel() {
	var channel models.Channel

	json.Unmarshal(this.Ctx.Input.RequestBody, &channel)

	models.AddChannel(channel)

	this.Data["json"] = models.GetPackingJobs(1, 30)
	this.ServeJSON()
}
func (this *ChannelController) GetChannels() {

	res := make(map[string]interface{})

	res["channels"] = models.GetChannels()

	this.Data["json"] = res
	this.ServeJSON()
}

func (this *ChannelController) DelChannel() {
	var channel models.Channel
	json.Unmarshal(this.Ctx.Input.RequestBody, &channel)

	models.DelChannel(channel)

	this.Data["json"] = channel
	this.ServeJSON()
}
