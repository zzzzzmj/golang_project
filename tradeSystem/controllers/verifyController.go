package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"tradeSystem/models"
)

type LoginVerifyController struct {
	beego.Controller
}

type RegisterVerifyController struct {
	beego.Controller
}

type FormValidationController struct {
	beego.Controller
}
type LogoutController struct {
	beego.Controller
}

//判断是否满足登录请求
func (L *LoginVerifyController) Post() {
	ua := L.GetString("userAccount")
	pwd := L.GetString("password")
	user := models.User{UserAccount: ua, Pwd: pwd}
	o := orm.NewOrm()
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		L.Data["jumpNotes"] = "用户名或密码错误，登陆失败，请重新登录"
		L.Data["jumpUrl"] = "http://localhost:8080/login"
		L.TplName = "jumpPage.html"
	} else {
		L.SetSession("isLogin", true)
		L.SetSession("userAccount", user.UserAccount)
		L.Data["stateString"] = "登陆成功！"
		L.Data["ua"] = ua
		L.Data["isLogin"] = true
		L.TplName = "index.html"
	}
}

//接收表单参数验证注册信息
func (R *RegisterVerifyController) Post() {
	var user = models.User{}
	R.ParseForm(&user)
	o := orm.NewOrm()
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		o.Insert(&user)
		R.Data["jumpNotes"] = "注册成功"
		R.Data["jumpUrl"] = "http://localhost:8080/index"

	} else {
		R.Data["jumpNotes"] = "注册失败"
		R.Data["jumpUrl"] = "http://localhost:8080/register"
	}
	R.TplName = "jumpPage.html"
}

//数据验证并且注册
func (v *FormValidationController) Post() {
	user := models.User{}
	shop := models.Shop{}
	valid := validation.Validation{}
	v.ParseForm(&user)
	shop.UserAccount = user.UserAccount
	shop.ShopId = user.ShopId
	var MessageTmpls = map[string]string{
		"Required": "不能为空",
		"MaxSize":  "最长长度为 %d",
		"Length":   "长度必须为 %d",
		"Mobile":   "必须是有效的手机号码",
	}
	validation.SetDefaultMessage(MessageTmpls)
	b, err := valid.Valid(&user)
	if err != nil {
		// 验证StructTag 是否正确
		v.Data["ParseFormErr"] = err
	}
	if !b {
		// 验证没通过 输出错误信息
		//valid.Errors 的信息存储在预定义的MessageTmpls中

		for _, err := range valid.Errors {
			//log.Println(err.Field, err.Message)
			data := "Verify" + err.Field
			v.Data[data] = err.Field + err.Message
		}
		v.Data["jumpUrl"] = "http://localhost:8080/register"
		v.TplName = "validError.html"
	} else {
		o := orm.NewOrm()
		err := o.Read(&user)
		if err == orm.ErrNoRows {
			o.Insert(&user)
			o.Insert(&shop)
			fmt.Println(shop.ShopId, shop.UserAccount)
			v.Data["jumpUrl"] = "http://localhost:8080/index"
			v.Data["jumpNotes"] = "注册成功"
		} else {
			v.Data["jumpUrl"] = "http://localhost:8080/register"
			v.Data["jumpNotes"] = "注册失败！"
		}
		v.TplName = "jumpPage.html"
	}

}

func (logout *LogoutController) Get() {
	ua := logout.GetSession("userAccount")
	if ua == nil {
		logout.Data["jumpUrl"] = "http://localhost:8080/login"
		logout.Data["jumpNotes"] = "退出失败，您还未登录！"
	} else {
		logout.DelSession("userAccount")
		logout.DelSession("isLogin")
		logout.Data["jumpUrl"] = "http://localhost:8080/index"
		logout.Data["jumpNotes"] = "退出成功！"
	}
	logout.TplName = "jumpPage.html"

}

/*
跳转页面的路由已修改并测试
*/
