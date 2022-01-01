package Controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ketangpai/model"
	"net/http"
)
var result  *gorm.DB  //记录查询结果
//注册用户

func Register(context *gin.Context)  {
	email:=context.PostForm("email")
	password:=context.PostForm("password")
	ReEnterPassword:=context.PostForm("reenter password")
	name:=context.PostForm("name")
	school:=context.PostForm("school")
	identity:=context.PostForm("identity")
	if email==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"邮箱为空"})
		return
	}
	if name==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"姓名为空"})
		return
	}
	if school==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"学校为空"})
		return
	}
	if identity==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"身份未选"})
		return
	}
	if password!=ReEnterPassword{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"密码输入确认与第一次不同"})
		return
	}
	if len(password)<8 || len(password)>20{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"密码长度不可小于8位或大于20位"})
		return
	}
	user:=model.User{Email: email}
	result=model.DB.Find(&user)
	if result.RowsAffected!=0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"该用户已存在"})
		return
	}
	user.Name=name
	user.Password=password
	if identity=="教师"{
		user.Identity=1
	}else if identity=="学生"{
		user.Identity=2
	}else{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"身份信息输入错误"})
		return
	}
	user.School=school
	result=model.DB.Create(&user)
	if result.RowsAffected!=1{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统异常"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"恭喜"+name+"注册成功"})
}
//删除用户

func DeleteUser(context *gin.Context){
	email:=context.PostForm("email")
	if email==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"邮箱为空,无法注销账户"})
		return
	}
	user:=model.User{Email: email}
	result=model.DB.Delete(&user)
	if result.RowsAffected==0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"该用户不存在"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"注销用户成功"})
}
