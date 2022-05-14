package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type FriendShip struct {
	Id          int32  `orm:"column(id);pk"`
	FriendId    string `orm:"column(friendId)"`
	UserAccount string `orm:"column(userAccount)"`
	//FriendName  string `orm:"column(friendName)"`
}

func init() {
	orm.RegisterModel(new(FriendShip))
}

func (friendship *FriendShip) TableName() string {
	// utils.GetConfig为自定义方法可以忽略只是在读取配置文件中的表前缀
	return fmt.Sprintf("%s", "friendship")
}
