package Middle

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)
//登录状态判断

func LoginMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		session:=sessions.Default(context)
		email:=session.Get("Email")
		if email==nil{
			context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"用户未登录"})
			context.Abort()
			return
		}
		emailString,_:=email.(string)
		context.Set("email",emailString)
	}
}
