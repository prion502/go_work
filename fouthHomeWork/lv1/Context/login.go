package Context

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginStruct struct {
}

func (con LoginStruct) Login(context *gin.Context){
	context.HTML(http.StatusOK,"login.html",gin.H{})
}
func (con LoginStruct)DoLogin(context *gin.Context)  {
	UseName:=context.PostForm("username")
	PassWord:=context.PostForm("password")
	cookies,err:=context.Request.Cookie(UseName)
	if err!=nil{
		cookies1:=&http.Cookie{
			Name: UseName,
			Value: PassWord,
			MaxAge: 3600,

		}
		fmt.Println(UseName,PassWord,err)
		http.SetCookie(context.Writer,cookies1)
		context.HTML(http.StatusOK,"error.html",gin.H{
			"message":"查不到该用户,将注册该账户",
		})
	}else{
		if cookies.Value==PassWord{
			context.HTML(http.StatusOK,"success.html",gin.H{
				"message":"用户账号密码正确,用户登录成功",
			})
		}else{
			context.HTML(http.StatusOK,"error.html",gin.H{
				"message":"用户登录失败，账户密码错误",
			})
		}
	}
}
