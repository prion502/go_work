#conf包
配置mysql数据库文件
#controller包
article.go,文章的控制器

like.go,点赞的控制器

login.go,登录的控制器

register.go,注册的控制器
#middleWare包
loginmiddleware.go,登录的中间件，实现token检验
#routers包
routers.go,配置路由
#tmp包
热加载的运行包
#tools包
token.go,实现token发放及token检验函数
#业务逻辑实现截图
```go
//实现携带token的登录
loginRouter.POST("",middleWare.LoginMiddle(),controller.Login)
```
#####未携带token登录

![img.png](E:\Goproject\src\fifthHomework\png\img.png)
#####携带token登录

![img_1.png](E:\Goproject\src\fifthHomework\png\img_1.png)
```go
//注册账户
registerRouter.POST("/",controller.Register)
```
#####注册账户

![img_2.png](E:\Goproject\src\fifthHomework\png\img_2.png)
```go
//注销账户
registerRouter.DELETE("/",controller.CancelRegister)
```
#####注销用户

![img_3.png](E:\Goproject\src\fifthHomework\png\img_3.png)
```go
//查看文章
articleRouter.GET("/",controller.SeeArticle)
```
#####查看文章

![img_4.png](E:\Goproject\src\fifthHomework\png\img_4.png)
```go
//修改文章
articleRouter.POST("/",controller.UpdateArticle)
```
#####修改文章

![img_5.png](E:\Goproject\src\fifthHomework\png\img_5.png)
```go
//删除文章
articleRouter.DELETE("/",controller.DeleteArticle)
```
#####删除文章

![img_6.png](E:\Goproject\src\fifthHomework\png\img_6.png)
```go
//点赞
likeRouter.POST("/",controller.AddLikeCount)
```
#####点赞

![img_7.png](E:\Goproject\src\fifthHomework\png\img_7.png)
```go
//查看点赞量
likeRouter.GET("/",controller.Like)
```
#####查看点赞量

![img_8.png](E:\Goproject\src\fifthHomework\png\img_8.png)
