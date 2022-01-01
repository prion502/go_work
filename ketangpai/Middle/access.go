package Middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//权限判断函数

func AccessMiddle() gin.HandlerFunc{
	return func(context *gin.Context) {
		email,_:=context.Get("email")
		emailString,_:=email.(string)
		user:=model.User{Email: emailString}
		model.DB.Find(&user)
        url:=context.Request.URL.Path
        fmt.Println(url)
		access:=model.Access{}
		result:=model.DB.Where("identity=? and path=?",user.Identity,url).Find(&access)
        if result.RowsAffected==0{
        	context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"权限不足"})
        	context.Abort()
			return
		}
	}
}
