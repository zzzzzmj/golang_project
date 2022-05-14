package models

import "github.com/astaxie/beego/orm"

type Image struct {
	Id         int64  `orm:"column(id);pk;auto"`
	FileName   string `orm:"column(fileName)"`
	FilePath   string `orm:"column(filePath)"`
	Createtime string `orm:"column(createTime)"`
	GoodsId    string `orm:"column(goodsId)"`
}

func init() {
	orm.RegisterModel(new(Image))
}
