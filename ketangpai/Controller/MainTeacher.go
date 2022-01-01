package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//教师课堂主页

func TeacherMain(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	var class []model.Class
	model.DB.Where("teacher_email=?",emailString).Find(&class)
	fmt.Println(emailString)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":emailString+"教师首页"})
	for i:=0;i<len(class);i++{
		context.JSON(http.StatusOK,gin.H{
			"课程名称":class[i].Name,
			"教学班级":class[i].ClassName,
			"总人数":class[i].StudentNum,
			"课堂码":class[i].Id,
			"学期":class[i].Year+"第"+class[i].Term+"学期",
		})
	}
}
