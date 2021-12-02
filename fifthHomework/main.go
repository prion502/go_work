package main

import (
	"fifthHomework/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()
	routers.Routers(r)
	r.Run()
}