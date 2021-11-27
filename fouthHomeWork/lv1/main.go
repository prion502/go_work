package main

import (
	"fouthHomeWork/lv1/Context"
	"github.com/gin-gonic/gin"
)
func main() {
	r:=gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/main",Context.LoginStruct{}.Login)
	r.POST("/LoginSign",Context.LoginStruct{}.DoLogin)
	r.Run()
}
