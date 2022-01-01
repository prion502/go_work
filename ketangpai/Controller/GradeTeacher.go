package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
	"strconv"
)
//添加成绩

func ADDGrade(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂"})
		return
	}
	studentgrade:=context.PostForm("grade")
	studentgradeint,_:=strconv.Atoi(studentgrade)
	homeworkfile:=context.PostForm("homework")
	studentemail:=context.PostForm("studentemail")
	fmt.Println(studentemail,homeworkfile,studentgradeint)
	joinclass:=model.JoinClass{}
	result=model.DB.Where("class_code=? and student_email=?",Id,studentemail).Find(&joinclass)
	if result.RowsAffected==0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"该学生未加入你课堂，是否输错信息!"})
		return
	}
	User:=model.User{}
	model.DB.Where("email=?",studentemail).Find(&User)
	grade:=&model.Grade{ClassCode: Id,StudentGrade: studentgradeint,HomeworkFile: homeworkfile,StudentName: User.Name,StudentEmail: studentemail}
	model.DB.Create(&grade)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"为"+User.Name+"添加成绩成功"})
}
//查看成绩

func CheckGrade(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.Query("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂"})
		return
	}
	homeworkfile:=context.Query("homework")
	grade:=[]model.Grade{}
	result=model.DB.Where("class_code=? and homework_file=?",Id,homeworkfile).Find(&grade)
	if result.RowsAffected==0{
		context.JSON(http.StatusOK,gin.H{"code":200,"message":"没有该次的成绩"})
		return
	}
	for i:=0;i<len(grade);i++{
		context.JSON(http.StatusOK,gin.H{"message":gin.H{"姓名":grade[i].StudentName,"作业":homeworkfile,"成绩":grade[i].StudentGrade}})
	}
}
//删除成绩

func DeleteGrade(context *gin.Context)  {
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂"})
		return
	}
	homeworkfile:=context.PostForm("homework")
	studentemail:=context.PostForm("studentemail")
	grade:=model.Grade{ClassCode: Id,HomeworkFile: homeworkfile,StudentEmail: studentemail}
	result=model.DB.Delete(&grade)
	if result.RowsAffected==0{
		context.JSON(http.StatusOK,gin.H{"code":200,"message":"未找到该人成绩"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":grade.StudentName+"的成绩已删除"})
}