package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}

func (I *IndexController) Get() {
	I.TplName = "index.html"
}

func (L *LoginController) Get() {
	L.TplName = "login.html"
}

func (R *RegisterController) Get() {
	R.TplName = "register.html"
}
