package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"tradeSystem/models"
)

type CreateOrderController struct {
	beego.Controller
}

func (co *CreateOrderController) Get() {
	goodsId := co.GetString("goodsId")
	shopId := co.GetStrings("shopId")
	var user models.User
	var goods models.Goods
	o := orm.NewOrm()
	o.QueryTable("user").Filter("shopId", shopId).One(&user)
	o.QueryTable("goods").Filter("goodsId", goodsId).One(&goods)
	paraMap := make(map[string]interface{})
	buyerId := co.GetSession("userAccount")
	paraMap["buyerId"] = buyerId.(string)
	paraMap["goodsId"] = goodsId
	paraMap["goodsName"] = goods.GoodsName
	paraMap["sellerId"] = user.UserAccount
	paraMap["totalPrice"] = goods.Price
	createTime := time.Now().Format("2006-01-02 15:04:05")
	paraMap["createTime"] = createTime
	order := models.Order{
		BuyerId:    buyerId.(string),
		SellerId:   user.UserAccount,
		CreateTime: createTime,
		GoodsId:    goodsId,
		GoodsName:  goods.GoodsName,
		TotalPrice: goods.Price,
		Amount:     goods.Amount,
	}
	o.Insert(&order)
	co.Data["map"] = paraMap
	co.TplName = "shop/createOrder.html"
}
