package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"tradeSystem/models"
)

type MyShopController struct {
	beego.Controller
}
type FunctionShopController struct {
	beego.Controller
}
type ShowMyShopController struct {
	beego.Controller
}

type CreateGoodsController struct {
	beego.Controller
}

type RemoveGoodsController struct {
	beego.Controller
}

type SearchAllShopController struct {
	beego.Controller
}

type SearchShopByIdController struct {
	beego.Controller
}

type SearchShopByNameController struct {
	beego.Controller
}

type ShowShopDetailByIdController struct {
	beego.Controller
}

func (ms *MyShopController) Get() {
	ms.TplName = "shop/myShop.html"
}

func (sc *ShowMyShopController) Get() {
	o := orm.NewOrm()
	userAccount := sc.GetSession("userAccount")
	var goodsList []*models.Goods
	user := models.User{}
	o.QueryTable("user").Filter("userAccount", userAccount).All(&user, "shopId")
	o.QueryTable("goods").Filter("shopId", user.ShopId).All(&goodsList)
	fmt.Println(goodsList)
	//查询不到记录说明这个商店还没有货物
	if goodsList == nil {
		sc.Data["states"] = "您的商店暂时还没有货物哟~~~"
		sc.TplName = "shop/showMyShop.html"
	} else {
		sc.Data["goodsList"] = goodsList
		sc.TplName = "shop/showMyShop.html"
	}
}

func (cg *CreateGoodsController) Get() {
	cg.TplName = "shop/createGoods.html"
}

func (cg *CreateGoodsController) Post() {
	goods := models.Goods{}
	o := orm.NewOrm()
	cg.ParseForm(&goods)
	ua := cg.GetSession("userAccount")
	var user = models.User{}
	o.QueryTable("user").Filter("userAccount", ua).All(&user, "shopId")
	goods.ShopId = user.ShopId
	o.Insert(&goods)
	cg.TplName = "shop/myShop.html"
}

func (rg *RemoveGoodsController) Get() {
	rg.TplName = "shop/removeGoods.html"
}

func (rg *RemoveGoodsController) Post() {
	var goodsList []*models.Goods
	user := models.User{}
	userAccount := rg.GetSession("userAccount")
	goodsId := rg.GetString("goodsId")
	id, _ := strconv.ParseInt(goodsId, 10, 64)
	var goods = models.Goods{GoodsId: id}
	o := orm.NewOrm()
	o.QueryTable("user").Filter("userAccount", userAccount).All(&user, "shopId")
	o.QueryTable("goods").Filter("shopId", user.ShopId).All(&goodsList)
	err := o.Read(&goods)
	if err == orm.ErrNoRows {
		rg.Ctx.WriteString("商品id输入错误，请重新输入")
	} else {
		goods.State = "OffSale"
		o.Update(&goods, "state")
		rg.Data["goodsList"] = goodsList
		rg.TplName = "shop/showMyShop.html"
	}
}

func (sa *SearchAllShopController) Get() {
	var maps []orm.Params
	o := orm.NewOrm()
	o.QueryTable("shop").Values(&maps)
	sa.Data["maps"] = maps
	sa.TplName = "shop/allShop.html"
}

func (si *SearchShopByIdController) Post() {
	shopId := si.GetString("shopId")
	var maps []orm.Params
	o := orm.NewOrm()
	o.QueryTable("shop").Filter("shopId", shopId).Values(&maps)
	if maps == nil {
		si.Ctx.WriteString("没有搜索的id的店铺，可能输入有误~~~~~")
	} else {
		si.Data["maps"] = maps
		si.TplName = "shop/searchShopById.html"
	}
}

func (sd *ShowShopDetailByIdController) Get() {
	shopId := sd.GetString("shopId")
	var maps []orm.Params
	o := orm.NewOrm()
	o.QueryTable("goods").Filter("shopId", shopId).Values(&maps)
	sd.Data["maps"] = maps
	fmt.Println(maps)
	sd.TplName = "shop/detailShop.html"
}

func (sn *SearchShopByNameController) Post() {
	goodsName := sn.GetString("goodsName")
	var maps []orm.Params
	o := orm.NewOrm()
	var sql string
	sql = "select * from goods where goodsName like '%" + goodsName + "%'"
	o.Raw(sql).Values(&maps)
	sn.Data["maps"] = maps
	sn.TplName = "shop/searchShopByName.html"
}

func (fs *FunctionShopController) Get() {
	fs.TplName = "functionShop.html"
}
