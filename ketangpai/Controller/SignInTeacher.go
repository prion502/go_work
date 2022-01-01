package Controller

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
	"strconv"
	"strings"
)
//发布签到、使用redis数据库，存储签到信息(应签人数,实签人数及姓名),
//设置了redis数据库过期时间为一天，同时使用cookie控制签到时间为1分钟

func CreateSignIn(context *gin.Context)  {
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂!"})
		return
	}
	studentNum:=class.StudentNum
	model.Init()
	var name=""
	_,err:=model.Redis.Do("SET","SignIn"+Id,name)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统Redis数据库存在问题"})
		return
	}
	_,err=model.Redis.Do("SET","StudentSum"+Id,strconv.Itoa(studentNum))
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统Redis数据库存在问题"})
		return
	}
	_,err=model.Redis.Do("expire","StudentSum"+Id,3600)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统Redis数据库存在问题"})
		return
	}
	_,err=model.Redis.Do("expire","SignIn"+Id,3600)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统Redis数据库存在问题"})
		return
	}
	context.SetCookie("SignTime","0",120,"/", "localhost", false, true)
	model.Redis.Close()
	context.JSON(http.StatusOK,gin.H{"code":200,"message":Id+"发布签到成功"})
}
//查看签到

func CheckSignIn(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.Query("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂,不能查看你的签到情况"})
		return
	}
	model.Init()
	name,err:=redis.String(model.Redis.Do("GET","SignIn"+Id))
	if err!=nil{
		if err.Error()=="redigo: nil returned"{
			context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"该缓冲记录已被消除"})
			return
		}
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统Redis数据库存在问题"})
		return
	}
	SumString,err1:=redis.String(model.Redis.Do("GET","StudentSum"+Id))
	model.Redis.Close()
	if err1!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统Redis数据库存在问题"})
		return
	}
	SumInt,_:=strconv.Atoi(SumString)
	nameSlice:=strings.Split(name," ")
	nameSlice=nameSlice[1:]
	context.JSON(http.StatusOK,gin.H{"code":200,"message":Id+"查看签到成功","data":gin.H{"已签到人数":len(nameSlice),
		"总人数":SumInt,
		"签到名单":nameSlice}})
}
