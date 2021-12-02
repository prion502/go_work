package controller

import (
	"fifthHomework/conf"
	"fifthHomework/tools"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
type Claims struct {
	UseName string
	jwt.StandardClaims
}
func Login(context *gin.Context){
	username, ok:=context.Get("user")
	if !ok || username==nil{
		var PasswordSql string
		UseName := context.PostForm("username")
		PassWord := context.PostForm("password")
		err := conf.DB.Ping()
		if err != nil {
			panic(err)
		}
		sql := "select password from user where username=?"
		err = conf.DB.QueryRow(sql, UseName).Scan(&PasswordSql)
		if err!= nil {
			context.JSON(http.StatusUnprocessableEntity,gin.H{"code":422, "message":"用户不存在"})
		} else {
			if PassWord == PasswordSql{
				//发放token
				token,err:=tools.ReleaseToken(UseName)
				if err!=nil{
					context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"token发放错误"})
					return
				}
				context.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"token":token},"message":"登录成功"})
			} else {
				context.JSON(http.StatusBadRequest,gin.H{"code":400, "message":"密码错误"})
			}
		}
	}else{
		context.JSON(http.StatusOK,gin.H{"code":200,"message":"登录成功"})
	}
}

