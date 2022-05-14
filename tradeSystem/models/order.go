package models

import "github.com/astaxie/beego/orm"

type Order struct {
	OrderId    int64  `orm:"column(orderId);pk;auto"`
	BuyerId    string `orm:"column(buyerId)"`
	SellerId   string `orm:"column(sellerId)"`
	CreateTime string `orm:"column(createTime)"`
	GoodsId    string `orm:"column(goodsId)"`
	GoodsName  string `orm:"column(goodsName)"`
	TotalPrice int32  `orm:"column(totalPrice)"`
	Amount     int32  `orm:"column(amount)"`
}

func init() {
	orm.RegisterModel(new(Order))
}
