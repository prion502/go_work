package model

import (
	"io/ioutil"
	"math/rand"
	"time"
)
//生成6位随机字符串（课堂码）

func  RandomString() string{
	s:=[]byte("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	str:=make([]byte,6)
	rand.Seed(time.Now().Unix())
	for i:=0;i<6;i++{
		str[i]=s[rand.Intn(len(s))]
	}
	return string(str)
}
//读取文件夹下的文件

func GetAllFile(pathname string) ([]string,error){
	var filename []string
	rd, err1 := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		filename=append(filename,fi.Name())
	}
	return filename,err1
}
