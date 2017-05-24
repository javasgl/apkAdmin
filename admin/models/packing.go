package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PackParams struct {
	JobId           int      `json:"-" orm:"PK"`
	ApkVersion      string   `json:"apkVersion" orm:"size(20)"`
	CheckedChannels []string `json:"checkedChannels" orm:"-"`
}

func AddPackingJob(packParams PackParams) bool {
	//todo
	beego.Debug(packParams)

	orm.NewOrm().Insert(&packParams)

	return false
}
func init() {
	orm.RegisterModel(new(PackParams))
}
