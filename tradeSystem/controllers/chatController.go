package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"tradeSystem/models"
)

type FunctionChatController struct {
	beego.Controller
}

type MyFriendsListsController struct {
	beego.Controller
}

type AddNewFriendsController struct {
	beego.Controller
}

type ShowRequestsController struct {
	beego.Controller
}

type DealRequestsController struct {
	beego.Controller
}

type CreateChatController struct {
	beego.Controller
}

type AddMessageController struct {
	beego.Controller
}

func (fc *FunctionChatController) Get() {
	fc.TplName = "chat/functionChat.html"
}

func (mf *MyFriendsListsController) Get() {
	var maps []orm.Params
	o := orm.NewOrm()
	ua := mf.GetSession("userAccount")
	o.QueryTable("friendship").Filter("userAccount", ua).Values(&maps)
	if maps == nil {
		mf.Data["empty"] = true
	} else {
		mf.Data["empty"] = false
	}
	mf.Data["fLists"] = maps
	mf.TplName = "chat/myFriendsLists.html"
}

func (af *AddNewFriendsController) Get() {
	af.TplName = "chat/addNewFriends.html"
}

func (af *AddNewFriendsController) Post() {
	friendId := af.GetString("searchId")
	notes := af.GetString("notes")
	ua := af.GetSession("userAccount")
	o := orm.NewOrm()
	var maps []orm.Params
	var maps2 []orm.Params
	o.QueryTable("user").Filter("userAccount", friendId).Values(&maps)
	o.QueryTable("friendship").Filter("userAccount", ua).Filter("friendId", friendId).Values(&maps2)
	//查询不到说明没有这个好友
	if maps2 == nil {
		//没有这个用户
		if maps == nil {
			//af.Data["noThisUser"] = true
			af.Data["jumpUrl"] = "http://localhost:8080/filter/addNewFriends"
			af.Data["jumpNotes"] = "没有这个用户~~~"
		} else {
			//af.Data["noThisUser"] = false
			fr := models.FriendsRequests{SenderId: ua.(string), ReceiverId: friendId,
				RequestsState: "等待对方验证", SendTime: time.Now().Format("2006-01-02 15:04:05"), Notes: notes}
			o.Insert(&fr)
			af.Data["jumpUrl"] = "http://localhost:8080/filter/chat"
			af.Data["jumpNotes"] = "请求发送成功"
			//af.Ctx.WriteString("请求发送成功！")
		}
	} else {
		af.Data["jumpUrl"] = "http://localhost:8080/filter/addNewFriends"
		af.Data["jumpNotes"] = "你们已经是好友啦，不能重复添加哟~~"
		//af.Ctx.WriteString("你们已经是好友啦~~~")
	}
	af.TplName = "jumpPage.html"
}

func (sr *ShowRequestsController) Get() {
	o := orm.NewOrm()
	ua := sr.GetSession("userAccount")
	var maps []orm.Params
	o.QueryTable("friendsrequests").Filter("receiverId", ua).Values(&maps)
	if maps == nil {
		//sr.Ctx.WriteString("您还没有好友通知哦~~~")
		sr.Data["jumpUrl"] = "http://localhost:8080/filter/myFriendsLists"
		sr.Data["jumpNotes"] = "您暂时还没有好友通知哟"
		sr.TplName = "jumpPage.html"
	} else {
		sr.Data["maps"] = maps
		sr.TplName = "chat/dealRequests.html"
	}
}

func (dr *DealRequestsController) Get() {
	res := dr.GetString("res")
	ua := dr.GetSession("userAccount")
	senderId := dr.GetString("senderId")
	fr := models.FriendsRequests{SenderId: senderId}
	o := orm.NewOrm()
	o.QueryTable("friendsrequests").Filter("senderId", senderId).One(&fr)
	var maps []orm.Params
	if res == "Yes" {
		fr.RequestsState = "已同意"

		if num, err := o.Update(&fr, "requestsState"); err == nil {
			fmt.Println("Update:", num)
		}
		fs1 := models.FriendShip{UserAccount: ua.(string), FriendId: senderId}
		fs2 := models.FriendShip{UserAccount: senderId, FriendId: ua.(string)}
		if num, err := o.Insert(&fs1); err == nil {
			fmt.Println(num)
		}
		o.Insert(&fs2)
	} else {
		fr.RequestsState = "已拒绝"
		if num, err := o.Update(&fr, "requestsState"); err == nil {
			fmt.Println("Update:", num)
		}
	}
	o.QueryTable("friendship").Filter("userAccount", ua).Values(&maps)
	fmt.Println(maps)
	dr.TplName = "chat/myFriendsLists.html"
}

func (cc *CreateChatController) Get() {
	friendId := cc.GetString("friendId")
	ua := cc.GetSession("userAccount")
	var maps []orm.Params
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("messageFromUserAccount", ua).Or("messageToUserAccount", ua)
	o.QueryTable("message").SetCond(cond1).Values(&maps)
	cc.Data["userAccount"] = ua
	cc.Data["friendId"] = friendId
	cc.Data["maps"] = maps
	cc.TplName = "chat/chatRoom.html"
}

func (am *AddMessageController) Get() {
	senderId := am.GetString("senderId")
	receiverId := am.GetString("receiverId")
	ua := am.GetString("userAccount")
	str := am.GetString("str")
	sendTime := time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	var maps []orm.Params
	message := models.Message{MessageType: "文本", PostMessage: str, MessageFromUserAccount: senderId, MessageToUserAccount: receiverId, SendTime: sendTime}
	o.Insert(&message)
	cond := orm.NewCondition()
	cond1 := cond.And("messageFromUserAccount", senderId).Or("messageToUserAccount", senderId)

	o.QueryTable("message").SetCond(cond1).Values(&maps)
	am.Data["maps"] = maps
	am.Data["userAccount"] = ua
	am.TplName = "chat/chatRoom.html"
}
