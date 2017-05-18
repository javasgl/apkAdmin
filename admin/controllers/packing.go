package controllers

import "github.com/astaxie/beego"

type PackingController struct {
	beego.Controller
}

func (this *PackingController) Get() {
	beego.Debug("channel::get")
	this.Layout = "layout.tpl"
	this.TplName = "packing/packing.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHeader"] = "packing/header.tpl"
	this.LayoutSections["HtmlBody"] = "packing/packing.tpl"
	this.LayoutSections["HtmlScripts"] = "packing/scripts.tpl"
}

func (this *PackingController) DoPacking() {

	beego.Debug(this.Ctx.Input.GetData("apkVersion"))
}
