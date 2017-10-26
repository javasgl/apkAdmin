package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	TABLE_CHANNEL = "apk_channels"
)

type Channel struct {
	ChannelId     int    `orm:"PK;auto" json:"channelId"`
	Channel       string `orm:"size(50)" json:"channel"`
	Remark        string `orm:"size(100)" json:"remark"`
	CreateTime    int64  `json:"-"`
	CreateTimeStr string `json:"createTime" orm:"-"`
}

func (c *Channel) TableName() string {
	return TABLE_CHANNEL
}

func init() {
	orm.RegisterModel(new(Channel))
}

func AddChannel(channel Channel) bool {

	var count int
	orm.NewOrm().Raw(fmt.Sprintf("SELECT COUNT(1) AS count FROM %s WHERE channel=?", TABLE_CHANNEL), channel.Channel).QueryRow(&count)

	if count > 0 {
		return false
	}

	channel.CreateTime = time.Now().Unix()

	_, err := orm.NewOrm().Insert(&channel)

	if err != nil {
		return false
	}
	return true
}

func GetChannels() []Channel {
	var channels []Channel
	orm.NewOrm().Raw(fmt.Sprintf("SELECT * FROM %s ORDER BY create_time DESC", TABLE_CHANNEL)).QueryRows(&channels)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for index, channel := range channels {
		channels[index].CreateTimeStr = time.Unix(channel.CreateTime, 0).In(loc).Format("2006-01-02 15:04:05")
	}
	return channels
}

func DelChannel(channel Channel) bool {
	_, err := orm.NewOrm().Delete(&channel)
	if err != nil {
		return false
	}
	return true
}
