package Controller

import (
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//查看自己的成绩

func LookGrade(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.Query("classCode")
	joinclass:=model.JoinClass{StudentEmail:emailString,ClassCode: Id}
	result=model.DB.Find(&joinclass)
	if result.RowsAffected==0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂"})
		return
	}
	homeworkfile:=context.Query("homework")
	grade:=model.Grade{ClassCode: Id,HomeworkFile: homeworkfile,StudentEmail: emailString}
	result=model.DB.Find(&grade)
	if result.RowsAffected==0{
		context.JSON(http.StatusOK,gin.H{"code":200,"message":"未找到该人成绩"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":gin.H{"姓名":grade.StudentName,"作业":homeworkfile,"成绩":grade.StudentGrade}})
}

