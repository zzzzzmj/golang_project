package routers

import (
	"github.com/astaxie/beego"
	"tradeSystem/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//首页，不需要过滤
	beego.Router("/index", &controllers.IndexController{})
	//登陆页面不需要过滤
	beego.Router("/login", &controllers.LoginController{})

	//登出页面，不需要过滤
	beego.Router("/logout", &controllers.LogoutController{})

	//注册页面不需要过滤
	beego.Router("/register", &controllers.RegisterController{})
	//验证登陆页面，不需要过滤
	beego.Router("/verify/login", &controllers.LoginVerifyController{})
	//验证注册页面不过滤
	beego.Router("/verify/register", &controllers.RegisterVerifyController{})
	//数据验证页面，不过滤
	beego.Router("/valid", &controllers.FormValidationController{})
	//需要过滤

	//商圈功能首页
	beego.Router("/filter/shop", &controllers.FunctionShopController{})
	//商店首页
	beego.Router("/filter/shop/index", &controllers.MyShopController{})
	//展示商店
	beego.Router("/filter/shop/showMyShop", &controllers.ShowMyShopController{})
	//添加物品页
	beego.Router("/filter/shop/createGoodsPage", &controllers.CreateGoodsController{})
	//删除物品
	beego.Router("/filter/shop/removeGoodsPage", &controllers.RemoveGoodsController{})
	//添加物品详情页
	beego.Router("/filter/shop/createGoods", &controllers.CreateGoodsController{})
	//删除物品详情页
	beego.Router("/filter/shop/removeGoods", &controllers.RemoveGoodsController{})
	//搜索所有注册的小店
	beego.Router("/filter/shop/searchAllShop", &controllers.SearchAllShopController{})
	//搜索指定id的小店
	beego.Router("/filter/shop/searchShopById", &controllers.SearchShopByIdController{})
	//搜索物品关键字展示小店
	beego.Router("/filter/shop/searchShopByName", &controllers.SearchShopByNameController{})
	//根据商店id展示商店
	beego.Router("/filter/shop/shopdetails", &controllers.ShowShopDetailByIdController{})
	beego.Router("/filter/buy", &controllers.CreateOrderController{})
	beego.Router("/filter/chat", &controllers.FunctionChatController{})
	beego.Router("/filter/myFriendsLists", &controllers.MyFriendsListsController{})
	beego.Router("/filter/addNewFriends", &controllers.AddNewFriendsController{})
	beego.Router("/filter/dealRequests", &controllers.DealRequestsController{})
	beego.Router("/filter/showFriendsRequests", &controllers.ShowRequestsController{})
	beego.Router("/filter/chat/createChat", &controllers.CreateChatController{})
	beego.Router("/filter/chat/addMessage", &controllers.AddMessageController{})
	beego.Router("/filter/shop/addImage", &controllers.AddImageController{})
	beego.Router("/filter/shop/upLoadImage", &controllers.ImageUploadController{})
	beego.Router("/filter/video/submitMyVideo", &controllers.SubmitMyVideo{})
	beego.Router("/filter/video", &controllers.ToVideoIndex{})
	beego.Router("/filter/video/watchAllVideo", &controllers.WatchSubmitedVideo{})

}
