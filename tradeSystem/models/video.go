package models

import "github.com/astaxie/beego/orm"

type Video struct {
	VideoId    int64  `orm:"column(videoId);pk;auto"`
	VideoUrl   string `orm:"column(videoUrl)"`
	VideoTitle string `orm:"column(videoTitle)"`
	SubmitUser string `orm:"column(submitUser)"`
	SubmitTime string `orm:"column(submitTime)"`
}

func init() {
	orm.RegisterModel(new(Video))
}

