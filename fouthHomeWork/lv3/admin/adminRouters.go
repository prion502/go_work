package admin

import (
	"fmt"
	"fouthHomeWork/lv3/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRouter(r *gin.Engine)  {
	adminRouter:=r.Group("/")
	{
		//登录界面方法
		adminRouter.POST("/sign", func(context *gin.Context) {
			var PasswordSql string
			UseName := context.PostForm("username")
			PassWord := context.PostForm("password")
			err := conf.DB.Ping()
			if err != nil {
				panic(err)
			}
			sql := "select password from user where username=?"
			err1 := conf.DB.QueryRow(sql, UseName).Scan(&PasswordSql)
			if err1 != nil {
				context.String(http.StatusOK, "该用户不存在，无法登录!请注册用户")
			} else {
				if PassWord == PasswordSql {
					context.String(http.StatusOK, "%v登录成功，进入博客页面", UseName)
					context.SetCookie("文章",UseName,3600,"/","localhost",false,true)
				} else {
					context.String(http.StatusOK, "%v登录失败,密码输入错误", UseName)
				}
			}
		})
	}

	{
		//注册界面的Post方法
		adminRouter.POST("/login", func(context *gin.Context) {
			UseName:=context.PostForm("username")
			PassWord:=context.PostForm("password")
			err:=conf.DB.Ping()
			if err!=nil{
				panic(err)
			}
			sqlStr := "insert into user(username, password) values (?,?)"
			_, err1 := conf.DB.Exec(sqlStr, UseName, PassWord)
			if err != nil {
				fmt.Printf("insert failed, err:%v\n", err1)
				return
			}
			context.String(http.StatusOK,"注册成功!")
		})
	}
	{
		//查看文章
		adminRouter.GET("/article", func(context *gin.Context) {
			username, _ := context.Cookie("文章")
			var article string
			sql := "select article from articlecontext where username=?"
			rows, err := conf.DB.Query(sql, username)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"message": "未查询到该用户的文章",
				})
			}
			defer rows.Close()
			for rows.Next() {
				err1 := rows.Scan(&article)
				if err1 != nil {
					context.JSON(http.StatusOK, gin.H{
						"message": "文章读取过程中出错",
					})
				}
				context.JSON(http.StatusOK, gin.H{
					"文章": article,
					"用户":username,
					"message":"文章读取成功",
				})
			}

		})
		//发表文章
		adminRouter.POST("/article", func(context *gin.Context) {
			username, err := context.Cookie("username")
			if err != nil {
				context.JSON(http.StatusForbidden, gin.H{
					"message": "请先登录！",
				})
				return
			}
			title := context.Query("title")
			content := context.Query("content")
			if content == "" || title == "" {
				context.JSON(http.StatusForbidden, gin.H{
					"message": "发表失败",
					"reason":  "文章为空",
				})
				return
			}
			sqlStr := "insert into article(username, title, articlecontent) values (?,?,?)"
			_, err = conf.DB.Exec(sqlStr, username, title, content)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"message": "发表失败！",
					"err":     err,
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"message":  "发表成功！",
				"username": username,
				"title":    title,
				"content":  content,
			})
		})
	}

	{ //点赞
		adminRouter.POST("/clickPraise", func(context *gin.Context) {
			UseName, err := context.Cookie("username")
			if err != nil {
				context.JSON(http.StatusForbidden, gin.H{
					"message": "请先登录！",
				})
				return
			}
			username := context.Query("username")
			title := context.Query("title")
			sqlStr := "update article set Prasie=Prasie+1 where username=? and title = ?"
			ret, err1 := conf.DB.Exec(sqlStr, username, title)
			if err1 != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
					"status":  http.StatusInternalServerError,
					"message": "点赞失败！",
					"err":     err,
				})
				return
			}
			n, err2 := ret.RowsAffected()
			if err2 != nil {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "点赞失败！",
				})
				return
			}
			if n == 0 {
				context.JSON(http.StatusForbidden, gin.H{
					"status":  http.StatusForbidden,
					"message": "点赞失败！",
					"reason":  "文章不存在!",
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"message": "点赞成功！"+UseName,
				"title":   title,
			})
		})
	}
}