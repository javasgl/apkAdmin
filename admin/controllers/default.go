package controllers

import (
	"encoding/json"

	"github.com/javasgl/apkAdmin/admin/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	if models.ValidateToken(this.Ctx) {
		this.Ctx.Redirect(302, "/dashboard/packing")
	}

	this.Layout = "layout.tpl"
	this.TplName = "login/login.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHeader"] = "login/header.tpl"
	this.LayoutSections["HtmlBody"] = "login/login.tpl"
	this.LayoutSections["HtmlScripts"] = "login/scripts.tpl"
}

func (this *MainController) DoLogin() {

	var loginUser models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &loginUser)

	loginRes, err := models.Login(loginUser)

	if err != nil {
		this.Data["json"] = 0
		this.DelSession(models.SessionKey)
	} else {
		this.Data["json"] = 1
		this.SetSession(models.SessionKey, loginRes)
	}
	this.ServeJSON()
}

func (this *MainController) Register() {
	this.Data["registerMail"] = "@" + models.RegisterMail
	this.Layout = "layout.tpl"
	this.TplName = "register/register.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHeader"] = "register/header.tpl"
	this.LayoutSections["HtmlBody"] = "register/register.tpl"
	this.LayoutSections["HtmlScripts"] = "register/scripts.tpl"
}

func (this *MainController) DoRegister() {
	var regiserUser models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &regiserUser)

	models.Register(regiserUser)

	beego.Debug(regiserUser)

	this.ServeJSON()
}
