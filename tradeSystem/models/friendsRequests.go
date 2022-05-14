package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type FriendsRequests struct {
	Id            int32  `orm:"column(id);pk"`
	SenderId      string `orm:"column(senderId)"`
	ReceiverId    string `orm:"column(receiverId)"`
	RequestsState string `orm:"column(requestsState)"`
	Notes         string `orm:"column(notes)"`
	SendTime      string `orm:"column(sendTime)"`
}

func init() {
	orm.RegisterModel(new(FriendsRequests))
}

func (fr *FriendsRequests) TableName() string {
	return fmt.Sprintf("%s", "friendsrequests")
}
