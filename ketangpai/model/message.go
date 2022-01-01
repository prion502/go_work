package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	TitleID   int    //对应话题序号
	Sender string   //发送者
	SendTime int64  //发送时间
	Content string  //发送内容
}

func (Message)Init(){
	DB.AutoMigrate(&Message{})
}

