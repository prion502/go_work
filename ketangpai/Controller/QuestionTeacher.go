package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)
//发布问题

var QuestionStatus=""   //记录回答状态
var Answer string    //记录回答答案
var ResponseName string  //记录回答人名字

func CreateQuestion(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	Id:=context.PostForm("classCode")
	class:=model.Class{Id: Id}
	model.DB.Find(&class)
	if class.TeacherEmail!=emailString{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"这不是你的课堂"})
		return
	}
	question:=context.PostForm("question")
	way:=context.PostForm("way")
	context.JSON(http.StatusOK,gin.H{"dn":"dwed"})
	if way=="抢答"{
		questionStruct:=model.Question{ClassCode: Id,Content: question,Way: 1}
		model.DB.Create(&questionStruct)
		context.JSON(http.StatusOK,gin.H{"code":200,"message":"发布问题"+strconv.Itoa(questionStruct.ID)+"题目为:"+question+"请同学们抢答"})
		now:=time.Now()
		fmt.Println(time.Now())
		for {
			if QuestionStatus==strconv.Itoa(questionStruct.ID){
				questionStruct.Answer=Answer
				questionStruct.ResponseStudent=ResponseName
				model.DB.Save(&questionStruct)
				context.JSON(http.StatusOK,gin.H{"code":200,"message":questionStruct.ResponseStudent+"抢答成功"+"答案为:"+Answer})
				QuestionStatus=""
				Answer=""
				ResponseName=""
				return
			}
			if time.Now().Sub(now)>=time.Second*10{
				context.JSON(http.StatusOK,gin.H{
					"code":"200",
					"message":"无人抢答",
				})
				return
			}
		}
	}
	if way=="抽答"{
		questionStruct:=model.Question{ClassCode: Id,Content: question,Way: 1}
		joinClass:=[]model.JoinClass{}
		model.DB.Where("id=? and teacher_email=?",Id,emailString).Find(&joinClass)
		rand.Seed(time.Now().Unix())
		n:=rand.Intn(len(joinClass))
		User:=model.User{Email: joinClass[n].StudentEmail}
		model.DB.Find(&User)
		context.JSON(http.StatusOK,gin.H{"code":200,"message":"发布问题"+strconv.Itoa(questionStruct.ID)+"题目为:"+question+"请"+User.Name+"回答"})
		now:=time.Now()
		for {
			if QuestionStatus==strconv.Itoa(questionStruct.ID){
				if ResponseName==User.Name{
					questionStruct.Answer=Answer
					model.DB.Save(&questionStruct)
					context.JSON(http.StatusOK,gin.H{"code":200,"message":ResponseName+"回答成功"+"答案为:"+Answer})
					QuestionStatus=""
					Answer=""
					ResponseName=""
					return
				}else{
					QuestionStatus=""
				}
			}
			if time.Now().Sub(now)>=time.Second*120{
				context.JSON(http.StatusOK,gin.H{
					"code":"200",
					"message":"无人抢答",
				})
				return
			}
		}
	}
	return
}
