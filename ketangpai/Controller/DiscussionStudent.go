package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
	"strconv"
	"time"
)
//参与话题讨论

func AttendDiscussion(context *gin.Context)  {
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	TitleId:=context.PostForm("TitleId")
	TitleIdInt,_:=strconv.Atoi(TitleId)
	User:=model.User{Email: emailString}
	model.DB.Find(&User)
	Message:=context.PostForm("message")
	message:=model.Message{TitleID: TitleIdInt,Sender: User.Name,Content: Message,SendTime: time.Now().Unix()}
	model.DB.Create(&message)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"留言成功"})
}
//查看话题讨论

func LookDiscussion(context *gin.Context){
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
	model.DB.Where("title_id=?",title.Id).Find(&message)
	fmt.Println(title)
	context.JSON(http.StatusOK,gin.H{"话题":title.TitleName+"     "+time.Unix(title.CreateTime,0).Format("2006-01-02 15:04"),
		"TitleId":title.Id})
	fmt.Println(message)
	for i:=0;i<len(message);i++{
		context.JSON(http.StatusOK,gin.H{message[i].Sender:message[i].Content +"     "+time.Unix(message[i].SendTime,0).Format("2006-01-02 15:04")})
	}
	return
}
