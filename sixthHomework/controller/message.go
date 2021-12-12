package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sixthHomework/conf"
	"time"
)

type Message struct{
	Receiver string
	Sender string
	Message string
	Time   int64
}
func SetMessage(context *gin.Context){
	var Username string
	master,err:=context.Cookie("username")
	if err!=nil || master==""{
		fmt.Println(err,master)
		context.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"没有账户登录,无法留言",
		})
		return
	}
	username:=context.PostForm("username")
	if username==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"数据传输错误"})
		return
	}
	sql:="select username from user where username=?"
	err3:=conf.DB.QueryRow(sql,username).Scan(&Username)
	if err3!=nil{
		fmt.Println(err3)
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"不存在该留言对象"})
		return
	}
	message:=context.PostForm("message")
	if message==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"数据传输错误"})
		return
	}
	Time:=time.Now().Unix()
	sql="insert into message(receiver,sender,message,time)values(?,?,?,?)"
	_,err1:=conf.DB.Exec(sql,username,master,message,Time)
	if err1!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"留言失败"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":master+"对"+username+"留言成功"})
}
func GetMessage(context *gin.Context){
	master,err:=context.Cookie("username")
	if err!=nil{
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"没有账户登录,无法查看自己的留言",
		})
	}
	sql:="select * from message where receiver=? ORDER BY time ASC;"
	row,err1:=conf.DB.Query(sql,master)
	if err1!=nil{
		context.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"该用户暂时没有留言",
		})
	}
	for row.Next(){
		var message Message
		err2:=row.Scan(&message.Receiver,&message.Sender,&message.Message,&message.Time)
		if err2!=nil{
			context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统错误"})
			return
		}
		if message.Receiver==""{
			context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"该用户没有留言"})
            return
		}
		context.JSON(http.StatusOK,gin.H{"留言":message.Sender+":"+message.Message+time.Unix(message.Time,0).Format("     2006/01/02 15:04")})
	}
	defer row.Close()
}
