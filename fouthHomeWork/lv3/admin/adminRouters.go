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
			username,_:=context.Cookie("文章")
			var article string
			sql := "select article from articleContext where username=?"
			rows,err:=conf.DB.Query(sql,username)
			if err!=nil {
				context.JSON(http.StatusOK,gin.H{
					"message":"未查询到该用户的文章",
				})
			}
			defer rows.Close()
			for rows.Next(){
				err1:=rows.Scan(&article)
				if err1!=nil{
					context.JSON(http.StatusOK,gin.H{
						"message":"文章读取过程中出错",
					})
				}
				context.JSON(http.StatusOK,gin.H{
					"文章":article,
				})
			}

		})
		//修改文章
		adminRouter.POST("/article", func(context *gin.Context) {

		})
		//删除文章
		adminRouter.DELETE("/article", func(context *gin.Context) {

		})
	}
	{   //点赞
		adminRouter.GET("/praise", func(context *gin.Context) {
		})
		//取消点赞
		adminRouter.POST("/praise", func(context *gin.Context) {
		})
	}
}
