package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
var err error
func init(){
	dsn:="root:20010712.@tcp(127.0.0.1:3307)/CSA?charset=utf8&parseTime=True&loc=Local"
	DB,err=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		fmt.Println(err)
	}
}
