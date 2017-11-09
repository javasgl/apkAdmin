package controllers

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/javasgl/apkAdmin/admin/models"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) DoUpload() {

	file, header, err := this.GetFile("splashImage")
	if err != nil {
		beego.Error(err)
	}
	defer file.Close()

	contentType := header.Header.Get("Content-Type")

	models.GetImageType(contentType)

	filename := "static/upload/splash/" + models.String2md5(header.Filename+strconv.Itoa(time.Now().Nanosecond())) + "." + models.GetImageType(contentType)

	this.SaveToFile("splashImage", filename)

	this.Data["json"] = map[string]string{
		"image": filename,
	}
	this.ServeJSON()
}
