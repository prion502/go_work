package routers

import (
	"fifthHomework/controller"
	"fifthHomework/middleWare"
	"github.com/gin-gonic/gin"
)
func Routers(r *gin.Engine){
	loginRouter:=r.Group("/login")
	{
		loginRouter.POST("",middleWare.LoginMiddle(),controller.Login)//携带token登录
	}
	registerRouter:=r.Group("/register")
	{
		registerRouter.POST("/",controller.Register)//注册账户
		registerRouter.DELETE("/",controller.CancelRegister)//注销账户
	}
	articleRouter:=r.Group("/article")
	{
		articleRouter.GET("/",controller.SeeArticle)//查看文章
		articleRouter.DELETE("/",controller.DeleteArticle)//删除文章
		articleRouter.POST("/",controller.UpdateArticle)//修改文章
	}
	likeRouter:=r.Group("/thumbsUp")
	{
		likeRouter.POST("/",controller.AddLikeCount)//点赞
		likeRouter.GET("/",controller.Like)//查看点赞量
	}
}
