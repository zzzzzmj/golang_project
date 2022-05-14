package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"tradeSystem/controllers"
	_ "tradeSystem/routers"
)

func main() {
	beego.InsertFilter("/filter/*", beego.BeforeRouter, controllers.MyFilter)
	//beego.SetStaticPath("/filter/video/*", "/static/video")
	//beego.SetStaticPath("/filter/shop/*", "/static/img")
	beego.Run()
}

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:0330zmj@tcp(127.0.0.1:3306)/gotradesystem?charset=utf8")
}
