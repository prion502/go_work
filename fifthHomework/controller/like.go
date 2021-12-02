package controller

import (
	"fifthHomework/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddLikeCount(context *gin.Context)  {
	username:=context.PostForm("username")
	title:=context.PostForm("title")
	sql:="update article set likecount=likecount+1 where username=? and title=?"
	conf.DB.Ping()
	ret,err:=conf.DB.Exec(sql,username,title)
	if err!=nil{
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"message":"系统错误",
		})
		return
	}
	if n,err1:=ret.RowsAffected();err1!=nil || n<1{
		context.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"不存在该文章",
		})
		return
	}
	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"message":"点赞文章成功",
	})
}
func Like(context *gin.Context)  {
	username:=context.Query("username")
	title:=context.Query("title")
	var likecount int
	sql:="select likecount from article where username=? and title=?"
	conf.DB.Ping()
	err:=conf.DB.QueryRow(sql,username,title).Scan(&likecount)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"message":"未找到该文章",
		})
		return
	}
	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{
			"username":username,
			"title":title,
			"likecount":likecount,
		},
		"message":"查看文章点赞量成功",
	})

}
