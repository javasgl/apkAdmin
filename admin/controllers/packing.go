package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/javasgl/apkAdmin/admin/models"
)

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

	beego.Debug(this.GetSession("apkuser"))
	var packJob models.PackingJobs
	json.Unmarshal(this.Ctx.Input.RequestBody, &packJob)
	models.AddPackingJob(packJob)
	this.Data["json"] = models.GetPackingJobs(1, 30)
	this.ServeJSON()
}
