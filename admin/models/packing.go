package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PackingJobs struct {
	JobId           int      `json:"-" orm:"PK"`
	ApkVersion      string   `json:"apkVersion" orm:"size(20)"`
	CheckedChannels []string `json:"checkedChannels" orm:"-"`
	ApkChannel      string   `json:"-"`
	CreateTime      int64
	CreatorId       int
	Status          int
	DownloadUrl     string `orm:"size(100)"`
}

func AddPackingJob(packJob PackingJobs) bool {
	//todo

	packJob.CreateTime = time.Now().Unix()
	packJob.ApkChannel = strings.Join(packJob.CheckedChannels, ",")

	beego.Debug(packJob)

	orm.NewOrm().Insert(&packJob)

	return false
}
func init() {
	orm.RegisterModel(new(PackingJobs))
}
