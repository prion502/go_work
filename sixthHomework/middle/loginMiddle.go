package middle

import "github.com/gin-gonic/gin"

func LoginMiddle()gin.HandlerFunc{
	return func(context *gin.Context) {
		answer:=context.PostForm("answer")
		if answer==""{
			context.Set("user","密码登录")
		}else {
			context.Set("user","密保登录")
		}
	}
}
