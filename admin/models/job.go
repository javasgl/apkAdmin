package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	TABLE_JOB = "apk_jobs"
)

type PackingJobs struct {
	JobId           int    `json:"-" orm:"PK"`
	AppId           int    `json:"appId"`
	ReleaseType     int    `json:"releaseType"`
	AppName         string `json:"appName"`
	ApkVersion      string `json:"apkVersion" orm:"size(20)"`
	CheckedChannels []int  `json:"checkedChannels" orm:"-"`
	ApkChannel      int    `json:"-"`
	CreateTime      int64  `json:"createTime"`
	CreatorId       int    `json:"creatorId"`
	Status          int    `json:"status"`
	DownloadUrl     string `orm:"size(100)" json:"downloadUrl`
	SplashImage     string `json:"splashImage"`
	Username        string `json:"username"`
}

func (c *PackingJobs) TableName() string {
	return TABLE_JOB
}

func AddPackingJob(packJob PackingJobs) bool {

	packJob.CreateTime = time.Now().Unix()

	orm.NewOrm().Insert(&packJob)

	GetPackingJobs(1, 20)

	return false
}

func GetPackingJobs(page, size int) []PackingJobs {
	var jobs []PackingJobs
	orm.NewOrm().Raw(fmt.Sprintf("SELECT * FROM `%s` ORDER BY `create_time` DESC LIMIT ?,?", TABLE_JOB), (page-1)*size, size).QueryRows(&jobs)
	return jobs
}

func GetJobsWithUser(page, size int) []PackingJobs {

	sql := "SELECT jobs.*,users.username FROM `%s` jobs LEFT JOIN `%s` users ON jobs.creator_id=users.user_id ORDER BY `create_time` DESC LIMIT ?,?"

	sql = fmt.Sprintf(sql, TABLE_JOB, TABLE_USER)

	var jobs []PackingJobs

	orm.NewOrm().Raw(sql, (page-1)*size, size).QueryRows(&jobs)

	return jobs
}

func init() {
	orm.RegisterModel(new(PackingJobs))
}
