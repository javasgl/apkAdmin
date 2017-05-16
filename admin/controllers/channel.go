package controllers

import "github.com/astaxie/beego"

type ChannelController struct {
	beego.Controller
}

func (this *ChannelController) Get() {
	beego.Debug("channel::get")
	this.Layout = "layout.tpl"
	this.TplName = "channel/channel.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHeader"] = "channel/header.tpl"
	this.LayoutSections["HtmlBody"] = "channel/channel.tpl"
	this.LayoutSections["HtmlScripts"] = "channel/scripts.tpl"
}
