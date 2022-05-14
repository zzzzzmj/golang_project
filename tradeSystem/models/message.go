package models

import "github.com/astaxie/beego/orm"

type Message struct {
	MessageId              int32  `orm:"column(messageId);pk"`
	PostMessage            string `orm:"column(postMessage)"`
	SendTime               string `orm:"column(sendTime)"`
	MessageType            string `orm:"column(messageType)"`
	MessageFromUserAccount string `orm:"column(messageFromUserAccount)"`
	MessageToUserAccount   string `orm:"column(messageToUserAccount)"`
}

func init() {
	orm.RegisterModel(new(Message))
}
