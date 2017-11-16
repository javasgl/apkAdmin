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

	this.Data["Logged"] = true
	this.Data["ActivedMenu"] = "/dashboard/packing"
	this.Data["loginUser"] = this.GetSession(models.SessionKey)
	this.Data["hostname"] = models.Hostname
	this.Layout = "layout.tpl"
	this.TplName = "packing/packing.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlScripts"] = "packing/scripts.tpl"
}

func (this *PackingController) AddJob() {

	var packJobs models.PackingJobs
	json.Unmarshal(this.Ctx.Input.RequestBody, &packJobs)

	session, ok := this.GetSession(models.SessionKey).(models.User)
	if !ok {
		beego.Error("session error:", this.GetSession(models.SessionKey))
		this.Data["json"] = "error"
		this.ServeJSON()
		return
	}
	for _, channelId := range packJobs.CheckedChannels {
		packJobs.ApkChannel = channelId
		packJobs.CreatorId = session.UserId
		models.AddPackingJob(packJobs)
	}
	this.Data["json"] = true
	this.ServeJSON()
}

func (this *PackingController) GetJobs() {

	res := make(map[string]interface{})

	res["jobs"] = models.GetJobsWithUser(1, 20)

	this.Data["json"] = res
	this.ServeJSON()
}
