package controller

import (
	"fifthHomework/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)
type User struct {
	Name  string
	Password string
}
func Register(context *gin.Context)  {
	var user User
	username:=context.PostForm("username")
	password:=context.PostForm("password")
	sql:="select * from user where username=?"
	conf.DB.Ping()
	rows,err:=conf.DB.Query(sql,username)
	for rows.Next(){
		rows.Scan(&user.Name,&user.Password)
		if user.Name==username{
			context.JSON(http.StatusBadRequest,gin.H{"code":400, "message":"该用户已存在"})
			return
		}
	}
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"code":400, "message":"该用户已存在"})
		return
	}
	rows.Close()
	sql="insert into user(username,password)values (?,?)"
	_,err=conf.DB.Exec(sql,username,password)
	if err!=nil{
		context.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"message":"系统错误"})
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"注册成功"})
}
func CancelRegister(context *gin.Context){
	username:=context.PostForm("username")
	sql:="delete from user where username=?"
	conf.DB.Ping()
	ret,err:=conf.DB.Exec(sql,username)
	if err!=nil{
		context.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"message":"系统错误"})
		return
	}
	n,err1:=ret.RowsAffected()
	if err1!=nil{
		context.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"message":"系统错误"})
		return
	}
	if n<1{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"该用户未注册过"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"注销用户成功"})
}
