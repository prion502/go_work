package Controller

import (
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
	"os"
)
//发布作业

func AssignHomework(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂号，不能上传作业"})
		return
	}
	file,err:=context.FormFile("homework")
	if err==nil{
		dir := "./static/homework/" + class.Id
		if err2:= os.MkdirAll(dir, 0666); err2!= nil {
			context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统问题"})
		}
		err1 := context.SaveUploadedFile(file,dir+"/"+file.Filename)
		if err1==nil{
			model.DB.Save(&class)
			context.JSON(http.StatusOK,gin.H{"filename":file.Filename+"上传成功"})
			return
		}else {
			context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统问题，作业未上传成功"})
			return
		}
	}else {
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"上传作业失败"})
		return
	}
}
