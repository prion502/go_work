package main

import (
	"fouthHomeWork/lv3/admin"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()
	admin.AdminRouter(r)
	r.Run()
}
