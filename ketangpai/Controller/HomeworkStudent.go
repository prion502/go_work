package Controller

import (
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
	"os"
)
//下载作业

func DownloadHomework(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	Id:=context.PostForm("classCode")
	joinClass:=model.JoinClass{ClassCode: Id}
	model.DB.Find(&joinClass)
	if joinClass.StudentEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"你未加入该课堂"})
		return
	}
	filename:=context.PostForm("filename")
	context.File("./static/homework/"+Id+filename)
}
//上交作业

func TakeInHomework(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	Id:=context.PostForm("classCode")
	joinClass:=model.JoinClass{ClassCode: Id}
	model.DB.Find(&joinClass)
	if joinClass.StudentEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"你未加入该课堂"})
		return
	}
	dir := "./static/ppt/" + joinClass.ClassCode
	if err := os.MkdirAll(dir, 0666); err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统问题"})
	}
	file,err:=context.FormFile("homework")
	if err==nil{
		err1 := context.SaveUploadedFile(file,dir+file.Filename)
		if err1==nil{
			context.JSON(http.StatusOK,gin.H{"code":200,"message":"作业上传成功"})
			return
		}else {
			context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统问题，作业未交成功"})
			return
		}
	}else {
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"作业未交成功"})
		return
	}
}
