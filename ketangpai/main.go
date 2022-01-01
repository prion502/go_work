package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"ketangpai/Routers"
	"ketangpai/model"
)
var store=cookie.NewStore([]byte("20010712."))
var R *gin.Engine
func main() {
	R=gin.Default()
	model.User{}.Init()
	model.Class{}.Init()
	model.JoinClass{}.Init()
	model.Message{}.Init()
	model.Grade{}.Init()
	model.Title{}.Init()
	R.Use(sessions.Sessions("mySessions",store))
	Routers.RouterInit(R)
	R.Run()
}
