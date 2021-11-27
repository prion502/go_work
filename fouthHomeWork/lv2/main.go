package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.LoadHTMLGlob("templates/*")
	//登录界面
	r.GET("/main", func(context *gin.Context) {
		context.HTML(http.StatusOK,"login.html",gin.H{})
	})
	r.POST("/main", func(context *gin.Context) {
		username:=context.PostForm("username")
		_,err:=context.Cookie(username)
		if err!=nil{
			context.SetCookie(username,"1",3600,"/","localhost",true,false)
			context.String(http.StatusOK,"游客你好")
		}else{
			context.String(http.StatusOK,"%v你好",username)
		}
	})

	r.Run()
}
