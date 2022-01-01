
package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//学生课堂主页

func StudentMain(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	var joinclass []model.JoinClass
	model.DB.Where("student_email=?",emailString).Find(&joinclass)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":emailString+"学生首页"})
	fmt.Println(joinclass)
	for i:=0;i<len(joinclass);i++{
		var class model.Class
		model.DB.Where("teacher_email=?",joinclass[i].TeacherEmail).Find(&class)
		title:=[]model.Title{}
		model.DB.Where("class_code=?",class.Id).Find(&title)
		titleSlice:=make(map[int]string,len(title))
		for k:=0;k<len(title);k++{
			titleSlice[k]=title[k].TitleName
		}
			user:=model.User{Email: class.TeacherEmail}
			model.DB.Find(&user)
			context.JSON(http.StatusOK,gin.H{
				"课程名称":class.Name,
				"教学班级":class.ClassName,
				"负责人(老师)":user.Name,
				"课堂码":class.Id,
				"学期":class.Year+"第"+class.Term+"学期",
			})
		}
}