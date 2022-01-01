package Controller

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//学生签到

func SignIn(context *gin.Context)  {
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	Id:=context.PostForm("classCode")
	joinClass:=model.JoinClass{ClassCode: Id}
	model.DB.Find(&joinClass)
	if joinClass.StudentEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"你未加入该课堂"})
		return
	}
	_,err:=context.Cookie("SignTime")
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"签到时间已结束"})
		return
	}
	User:=model.User{Email: emailString}
	model.DB.Find(&User)
	name:=User.Name
	model.Init()
	nameString,err1:=redis.String(model.Redis.Do("GET","SignIn"+Id))
	if err1!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统Redis数据库存在问题"})
		return
	}
	nameString+=" "+name
	_,err1=model.Redis.Do("SET","SignIn"+Id,nameString)
	model.Redis.Close()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统Redis数据库存在问题"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"签到成功"})
}