package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"sixthHomework/conf"
)

func Login(context *gin.Context){
	str,_:=context.Get("user")
	str=str.(string)
	if str=="密码登录" {
		var PasswordSql string
		UseName := context.PostForm("username")
		PassWord := context.PostForm("password")
		err := conf.DB.Ping()
		if err != nil {
			panic(err)
		}
		sql := "select password from user where username=?"
		err = conf.DB.QueryRow(sql, UseName).Scan(&PasswordSql)
		if err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "用户不存在"})
		} else {
			if PassWord == PasswordSql {
				context.SetCookie("username",UseName,3600,"/","localhost",false,true)
				context.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录成功"})
			} else {
				context.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "密码错误,可输入密保答案进行密码修改及登录"})
			}
		}
	}else{
		var Answer string
		username:=context.PostForm("username")
		password := context.PostForm("password")
		answer:=context.PostForm("answer")
		sql:="select answer from user where username=?"
		err:=conf.DB.QueryRow(sql,username).Scan(&Answer)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"该用户不存在"})
			return
		}
		if Answer==answer{
			sql="update user set password=? where username=?"
			_,err=conf.DB.Exec(sql,password,username)
			if err!=nil{
				context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统错误"})
				return
			}else{
				context.SetCookie("username",username,3600,"/","localhost",true,false)
				context.JSON(http.StatusOK,gin.H{"code":500,"message":"登录成功且密码已修改"})
			}
		}else{
			context.JSON(http.StatusOK,gin.H{"code":200,"message":"密保错误"})
		}
	}
}
func CancelLogin(context *gin.Context){
	UseName:=context.Query("username")
	context.SetCookie("username",UseName,-1,"/","localhost",false,true)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":UseName+"退出登录成功"})
}
