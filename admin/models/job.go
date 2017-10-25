package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	TABLE_JOB = "apk_jobs"
)

type PackingJobs struct {
	JobId           int      `json:"-" orm:"PK"`
	AppId           int      `json:"appId"`
	ReleaseType     int      `json:"releaseType"`
	AppName         string   `json:"appName"`
	ApkVersion      string   `json:"apkVersion" orm:"size(20)"`
	CheckedChannels []string `json:"checkedChannels" orm:"-"`
	ApkChannel      string   `json:"-"`
	CreateTime      int64    `json:"createTime"`
	CreatorId       int      `json:"createId"`
	Status          int      `json:"status"`
	DownloadUrl     string   `orm:"size(100)" json:"downloadUrl`
	SplashImage     string   `json:"splashImage"`
}

func (c *PackingJobs) TableName() string {
	return TABLE_JOB
}

func AddPackingJob(packJob PackingJobs) bool {
	//todo

	packJob.CreateTime = time.Now().Unix()
	packJob.ApkChannel = strings.Join(packJob.CheckedChannels, ",")

	beego.Debug(packJob)

	orm.NewOrm().Insert(&packJob)

	GetPackingJobs(1, 20)

	return false
}

func GetPackingJobs(page, size int) []PackingJobs {
	var jobs []PackingJobs
	orm.NewOrm().Raw(fmt.Sprintf("SELECT * FROM %s ORDER BY create_time DESC LIMIT ?,?", TABLE_JOB), (page-1)*size, size).QueryRows(&jobs)
	for index, job := range jobs {
		if job.ApkChannel != "" {
			jobs[index].CheckedChannels = strings.Split(job.ApkChannel, ",")
		} else {
			jobs[index].CheckedChannels = []string{}
		}
	}
	return jobs
}
func init() {
	orm.RegisterModel(new(PackingJobs))
}
