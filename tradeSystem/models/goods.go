package models

import "github.com/astaxie/beego/orm"

type Goods struct {
	ShopId    string `orm:"column(shopId)" form:"-"`
	GoodsId   int64  `orm:"column(goodsId);auto;pk" form:"-"`
	Amount    int32  `orm:"column(amount)" form:"Amount"`
	Price     int32  `orm:"column(price)" form:"Price"`
	State     string `orm:"column(state)" form:"state"`
	GoodsName string `orm:"column(goodsName)" form:"GoodsName"`
	ImageUrl  string `orm:"column(imageUrl)" form:"ImageUrl"`
}

func init() {
	orm.RegisterModel(new(Goods))
}
