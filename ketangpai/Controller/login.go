package Controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//登录，使用sessions

func Login(context *gin.Context){
	session:=sessions.Default(context)
	email:=session.Get("Email")
	email1:=context.PostForm("email")
	if email!=nil{
		emailString,_:=email.(string)
		if email1==emailString{
			context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"用户已登录"})
			return
		}
	}
	password:=context.PostForm("password")
	user:=model.User{Email: email1,Password:password}
	result=model.DB.Find(&user)
	if result.RowsAffected==0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"该用户不存在,无法登录"})
		return
	}
	sessions:=sessions.Default(context)
	sessions.Set("Email",user.Email)
	sessions.Save()
	if user.Identity==1{
		context.Request.Method="GET"
		context.Redirect(http.StatusMovedPermanently,"http://localhost:8080/teacher")
	} else {
		context.Request.Method="GET"
		context.Redirect(http.StatusMovedPermanently,"http://localhost:8080/student")
	}

}
//取消登录

func CancelLogin(context *gin.Context){
	session:=sessions.Default(context)
	Email:=session.Get("Email")
	_,ok:=Email.(string)
	if !ok{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统错误,未取消登录"})
		return
	}
	session.Delete("Email")
	session.Save()
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"退出登录成功"})
}
