package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	UserAccount string `orm:"column(userAccount);pk" form:"userAccount" valid:"Required;MaxSize(8)"`
	NickName    string `orm:"column(nickName)" form:"nickName" valid:"Required;MaxSize(10)"`
	Pwd         string `orm:"column(pwd)" form:"password" valid:"Required;MaxSize(10)"`
	Sex         string `orm:"column(sex)" form:"sex"`
	Addr        string `orm:"column(addr)" form:"addr"`
	PhoneNumber string `orm:"column(phoneNumber)" form:"phoneNumber" valid:"Required;Mobile"`
	TrueName    string `orm:"column(trueName)" form:"trueName"`
	Email       string `orm:"column(email)" form:"email" valid:"Required"`
	ShopId      string `orm:"column(shopId)" form:"shopId" valid:"Required;MaxSize(8)"`
}

func init() {
	orm.RegisterModel(new(User))
}
