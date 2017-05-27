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

	// orm.NewOrm().Insert(&packJob)

	GetPackingJobs(1, 20)

	return false
}

func GetPackingJobs(page, size int) []PackingJobs {
	var jobs []PackingJobs
	orm.NewOrm().Raw("SELECT * FROM packing_jobs ORDER BY create_time DESC LIMIT ?,?", (page-1)*size, size).QueryRows(&jobs)
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
