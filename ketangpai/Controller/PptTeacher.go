package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
	"os"
	"path"
)
//上传课件

func UploadPpt(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂"})
		return
	}
	dir := "./static/ppt/" + class.Id
	if err := os.MkdirAll(dir, 0666); err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统问题"})
	}
	form,_:=context.MultipartForm()
	fmt.Println(form)
	files:=form.File["ppt[]"]
	fmt.Println(files)
	for _, file := range files {
		dst := path.Join(dir, file.Filename)
		err:=context.SaveUploadedFile(file, dst)
		fmt.Println(err)
	}
	context.JSON(http.StatusOK, gin.H{ "code":200,"message": "ppt上传成功"})
}
