package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
	"strconv"
	"time"
)
//创建话题讨论

func CreateTitle(context *gin.Context)  {
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂,不能创建话题"})
		return
	}
	TitleName:=context.PostForm("title")
	Title:=model.Title{ClassCode:Id,TitleName:TitleName,TeacherEmail: emailString,CreateTime: time.Now().Unix()}
	result=model.DB.Create(&Title)
	if result.RowsAffected!=0{
		context.JSON(http.StatusOK,gin.H{"code":200,"message":"创建讨论成功","TitleID为":Title.Id})
		return
	}
	context.JSON(http.StatusBadRequest,gin.H{"code":200,"message":"创建讨论失败"})
}
//查看话题讨论

func CheckDiscussion(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	TitleId:=context.Query("TitleId")
	TitleIdInt,_:=strconv.Atoi(TitleId)
	fmt.Println(TitleIdInt)
	message:=[]model.Message{}
	title:=model.Title{}
	result=model.DB.Where("id=?",TitleIdInt).Find(&title)
	if result.RowsAffected==0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":TitleId+"话题不存在"})
		return
	}
	model.DB.Where("id=?",uint(title.Id)).Find(&message)
	context.JSON(http.StatusOK,gin.H{"话题":title.TitleName+"     "+time.Unix(title.CreateTime,0).Format("2006-01-02 15:04"),
		"TitleId":title.Id})
	for i:=0;i<len(message);i++{
		context.JSON(http.StatusOK,gin.H{message[i].Sender:message[i].Content +"     "+time.Unix(message[i].SendTime,0).Format("2006-01-02 15:04")})
	}
	return
}

