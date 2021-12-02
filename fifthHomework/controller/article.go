package controller

import (
	"fifthHomework/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
type Article struct {
	Username string
	Title  string
	Content string
	LikeCount  string
}
func SeeArticle(context *gin.Context){
	username:= context.Query("username")
	title:=context.Query("title")
	var content string
	sql:="select content from article where username=? and title=?"
	err:=conf.DB.Ping()
	err=conf.DB.QueryRow(sql,username,title).Scan(&content)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"不存在该文章"})
		return
	}
	if content==""{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"文章内容为空"})
		return
	}
	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{
			"content":content,
			"title":title,
		},
		"message":"查看文章成功",
	})
}
func DeleteArticle(context *gin.Context)  {
	username:=context.PostForm("username")
	title:=context.PostForm("title")
	sql:="delete from article where username=? and title=?"
	conf.DB.Ping()
	rows,err:=conf.DB.Exec(sql,username,title)
	if err!=nil{
		context.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"message":"系统错误"})
		return
	}
	n,err1:=rows.RowsAffected()
	if err1!=nil{
		context.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"message":"系统错误"})
	}
	if n<1{
		context.JSON(http.StatusBadRequest,gin.H{"code":400,"message":"不存在该文章"})
	}
	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{
			"username":username,
			"title":title,
		},
		"message":"文章删除成功",
	})
}
func UpdateArticle(context *gin.Context){
	username:=context.PostForm("username")
	title:=context.PostForm("title")
	NewContent:=context.PostForm("content")
	var NewTitle string
	sql:="select title from article where username=?"
	conf.DB.Ping()
	rows,err:=conf.DB.Query(sql,username)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"该用户不存在",
		})
		return
	}
	for rows.Next(){
		rows.Scan(&NewTitle)
		if NewTitle==title{
			sqlStr := "update article set content=? where username = ? and title=?"
			_,err=conf.DB.Exec(sqlStr,NewContent,username,title)
			if err!=nil{
				context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"系统错误"})
				return
			}
			context.JSON(http.StatusOK,gin.H{"code":200,"message":"修改文章成功"})
			return
		}
	}
	sql= "insert into article(username, title,content,likecount) values (?,?,?,?)"
	_,err1:=conf.DB.Exec(sql,username,title,NewContent,0)
	if err1!=nil{
		fmt.Println(err1)
		context.JSON(http.StatusInternalServerError,gin.H{"code":500,"message":"添加文章出错"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"code":200,"message":"添加文章成功"})
}

//	{ //点赞
//		adminRouter.POST("/clickPraise", func(context *gin.Context) {
//			UseName, err := context.Cookie("username")
//			if err != nil {
//				context.JSON(http.StatusForbidden, gin.H{
//					"message": "请先登录！",
//				})
//				return
//			}
//			username := context.Query("username")
//			title := context.Query("title")
//			sqlStr := "update article set Prasie=Prasie+1 where username=? and title = ?"
//			ret, err1 := conf.DB.Exec(sqlStr, username, title)
//			if err1 != nil {
//				context.JSON(http.StatusInternalServerError, gin.H{
//					"status":  http.StatusInternalServerError,
//					"message": "点赞失败！",
//					"err":     err,
//				})
//				return
//			}
//			n, err2 := ret.RowsAffected()
//			if err2 != nil {
//				context.JSON(http.StatusOK, gin.H{
//					"status":  http.StatusOK,
//					"message": "点赞失败！",
//				})
//				return
//			}
//			if n == 0 {
//				context.JSON(http.StatusForbidden, gin.H{
//					"status":  http.StatusForbidden,
//					"message": "点赞失败！",
//					"reason":  "文章不存在!",
//				})
//				return
//			}
//			context.JSON(http.StatusOK, gin.H{
//				"message": "点赞成功！"+UseName,
//				"title":   title,
//			})
//		})
//	}
//}
