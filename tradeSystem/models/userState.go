package models

import "github.com/astaxie/beego/orm"

type UserState struct {
	UserAccount string `orm:"column(userAccount);pk"`
	State       string `orm:"column(state)"`
}

func init() {
	orm.RegisterModel(new(UserState))
}
