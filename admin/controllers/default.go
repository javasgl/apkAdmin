package controllers

import (
	"encoding/json"

	"github.com/javasgl/apkAdmin/admin/models"
	"github.com/javasgl/apkAdmin/admin/utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Layout = "layout.tpl"
	this.TplName = "login/login.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHeader"] = "login/header.tpl"
	this.LayoutSections["HtmlBody"] = "login/login.tpl"
	this.LayoutSections["HtmlScripts"] = "login/scripts.tpl"
}
func (this *MainController) DoLogin() {

	var loginUser models.LoginUser

	json.Unmarshal(this.Ctx.Input.RequestBody, &loginUser)

	loginSuccess := models.Login(loginUser.Username, loginUser.Password)
	if loginSuccess {
		this.Ctx.Output.Cookie("apksystoken", utils.GenerateToken(loginUser.Username), 86400*30)
		this.Ctx.Output.Cookie("username", loginUser.Username, 86400*30)
	}
	this.Data["json"] = loginSuccess
	this.ServeJSON()
}
