package main

import (
	"github.com/gin-gonic/gin"
	"sixthHomework/controller"
	"sixthHomework/middle"
)

func main(){
	r:=gin.Default()
	loginRouter:=r.Group("/login")
	{
		loginRouter.POST("/",middle.LoginMiddle(),controller.Login)//登录
		loginRouter.GET("/",controller.CancelLogin)//注销登录
		loginRouter.GET("/message",controller.GetMessage)//获取本账户留言
		loginRouter.POST("/message",controller.SetMessage)//设置留言
	}
	registerRouter:=r.Group("/register")
	{
		registerRouter.POST("/",controller.Register)//注册账户
		registerRouter.DELETE("/",controller.CancelRegister)//注销账户
	}
	r.Run()
}
