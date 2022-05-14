package models

import "github.com/astaxie/beego/orm"

type Shop struct {
	ShopId      string `orm:"column(shopId);pk"`
	UserAccount string `orm:"column(userAccount)"`
}

func init() {
	orm.RegisterModel(new(Shop))
}
