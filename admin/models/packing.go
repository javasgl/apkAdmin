package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PackingJobs struct {
	JobId           int      `json:"-" orm:"PK"`
	ApkVersion      string   `json:"apkVersion" orm:"size(20)"`
	CheckedChannels []string `json:"checkedChannels" orm:"-"`
	CreateTime      int
	CreatorId       int
	Status          int
	DownloadUrl     string `orm:"size(100)"`
}

func AddPackingJob(packJob PackingJobs) bool {
	//todo
	beego.Debug(packJob)

	orm.NewOrm().Insert(&packJob)

	return false
}
func init() {
	orm.RegisterModel(new(PackingJobs))
}
