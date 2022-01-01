package Controller

import (
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//下载课件

func DownLoadPPT(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	Id:=context.PostForm("classCode")
	joinClass:=model.JoinClass{}
	result=model.DB.Where("student_email=? and class_code=?",emailString,Id).Find(&joinClass)
	if result.RowsAffected==0{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"你未加入该课堂"})
		return
	}
	filename:=context.PostForm("filename")
	context.File("./static/ppt/"+Id+"/"+filename)
}

