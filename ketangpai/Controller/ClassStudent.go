package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//加入课堂

func JoinClass(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	fmt.Println(Id)
	class:=model.Class{Id: Id}
	result=model.DB.Find(&class)
	if result.RowsAffected==0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"不存在该课堂"})
		return
	}
	joinClass:=model.JoinClass{}
	result=model.DB.Where("student_email=? and teacher_email=? and class_code=?",emailString,class.TeacherEmail,Id).Find(&joinClass)
	if result.RowsAffected!=0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"已加过该课堂"})
		return
	}
	joinClass.ClassCode=Id
	joinClass.TeacherEmail=class.TeacherEmail
	joinClass.StudentEmail=emailString
	err:=model.DB.Create(&joinClass)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统服务问题"})
	}
	class.StudentNum++
	model.DB.Save(&class)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"加入该课堂成功"})
}
//退出课堂

func ExitClass(context *gin.Context)  {
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	Id:=context.PostForm("classCode")
	joinClass:=model.JoinClass{StudentEmail: emailString,ClassCode: Id}
	model.DB.Find(&joinClass)
	if model.DB.RowsAffected==0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"之前未加入该课堂"})
		return
	}
	model.DB.Delete(&joinClass)
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	class.StudentNum--
	model.DB.Save(&class)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"退出课堂成功"})
}
