package Controller

import (
	"github.com/gin-gonic/gin"
	"ketangpai/model"
	"net/http"
)
//回答问题

func AnswerQuestion(context *gin.Context){
	email,_:=context.Get("email")
	emailString,_:=email.(string)
	QuestionId:=context.PostForm("id")
	QuestionStatus=QuestionId
	Answer=context.PostForm("answer")
	User:=model.User{Email: emailString}
	model.DB.Find(&User)
	ResponseName=User.Name
	context.JSON(http.StatusOK,gin.H{"code":200,"message":User.Name+"回答成功"})
}
