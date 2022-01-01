package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
	"strconv"
	"time"
)
//创建课堂,返回随机生成的6位课堂码(不可重复)

func CreateClass(context *gin.Context)  {
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Name:=context.PostForm("name")
	ClassName:=context.PostForm("classname")
	Year:=context.PostForm("year")
	Term:=context.PostForm("term")
	fmt.Println(ClassName,Name,Term,Year)
	var Id string
	var class model.Class
	for {
		Id=model.RandomString()
		class=model.Class{Id: Id}
		result=model.DB.Find(&class)
		if result.RowsAffected==0{
			break
		}
	}
	fmt.Println(ClassName,Name,Term,Year)
	class.ClassName=ClassName
	class.Id=Id
	class.Name=Name
	class.StudentNum=0
	class.Term=Term
	class.Year=Year
	class.TeacherEmail=emailString
	class.Status=1
	fmt.Println(class)
	model.DB.Create(&class)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"课堂创建成功","classCode":class.Id})
}
//删除课堂

func DeleteClass(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{
		Id:Id,
	}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂，不能删除"})
		return
	}
	model.DB.Delete(&class)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"删除课堂成功"})
}
//禁用课堂,修改课堂状态

func NoUseClass(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂，不能删除"})
		return
	}
	class.Status=2
	model.DB.Save(&class)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"停用该课堂成功"})
}
//使用该课堂

func UseClass(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	if emailString==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"未登录"})
		return
	}
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂，不能删除"})
		return
	}
	class.Status=1
	model.DB.Save(&class)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"启动该课堂成功"})
}
//更换课堂的课堂码

func UpdateClassCode(context *gin.Context)  {
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString {
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂，不能删除"})
		return
	}
	var class1 model.Class
	var joinClass []model.JoinClass
	for  {
		class1.Id=model.RandomString()
		result=model.DB.Find(&class1)
		if result.RowsAffected==0{
			break
		}
	}
	model.DB.Where("class_code=?",class.Id).Find(&joinClass)
	for i:=0;i<len(joinClass);i++{
		joinClass[i].ClassCode=class1.Id
	}
	model.DB.Save(&joinClass)
	class.Id=class1.Id
	model.DB.Save(&class)
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"修改课程码成功为:"+class.Id})
}
//查看课堂信息

func CheckClass(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.Query("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂"})
		return
	}
	var joinclass []model.JoinClass
	model.DB.Where("class_code=? and teacher_email=?",Id,emailString).Find(&joinclass)
	context.JSON(http.StatusOK,gin.H{"课堂码为":Id+"的首页"})
	for i:=1;i<=len(joinclass);i++{
		var user model.User
		model.DB.Where("email=?",joinclass[i-1].StudentEmail).Find(&user)
		context.JSON(http.StatusOK,gin.H{"第"+strconv.Itoa(i)+"个学生":user.Name})
	}
	var title []model.Title
	model.DB.Where("class_code=? and teacher_email=?",Id,emailString).Find(&title)
	if len(title)>0{
		for i:=1;i<=len(title);i++{
			context.JSON(http.StatusOK,gin.H{"话题":title[i].TitleName+"     "+time.Unix(title[i].CreateTime,0).Format("2006-01-02 15:04"),
				"TitleId":title[i].Id})
		}}else{
		context.JSON(http.StatusOK,gin.H{"话题":"还没有话题"})
	}
	pathname1:="./static/ppt"+"/"+Id
	filename1,err1:=model.GetAllFile(pathname1)
	pathname:="./static/homework"+"/"+Id
	filename,err:=model.GetAllFile(pathname)
	if err1!=nil && err==nil{
		context.JSON(http.StatusOK,gin.H{"作业":filename,"ppt":"还没有ppt"})
	}else if err!=nil && err1==nil{
		context.JSON(http.StatusOK,gin.H{"作业":"还没有作业","ppt":filename1})
	}else if err!=nil && err1!=nil{
		context.JSON(http.StatusOK,gin.H{"作业":"还没有作业","ppt":"还没有ppt"})
	}else {
		context.JSON(http.StatusOK,gin.H{"作业":filename,"ppt":filename1})
	}
	return
}
