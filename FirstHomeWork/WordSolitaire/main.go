package main

//@Title  CSA/WordSolitaire/main.go
//@Description   词语接龙
//@Author  Dbs(董百顺)
//@Update 2021-11-05

import (
	"fmt"
)
func main() {
	var StringSet =make([]string,0,100)  //StringSet存储字符串
	var IndexSet=make([]bool,0,100)   //存储StringSet中字符串使用情况，未使用false，使用true
	var String string    //定义输入字符串变量
	var Char byte    //定义词语接龙的首字母变量
	fmt.Scanf("%s",&String)
	for String!="nil"{
		StringSet=append(StringSet,String)
		IndexSet=append(IndexSet,false)
		fmt.Scanf("%s",&String)
	}
	fmt.Scanf("%c",&Char)
	//for 循环进行词语接龙
	for {
		S:=Char
		for key, value := range StringSet {
			if byte(value[0]) == Char {
				if IndexSet[key] == false {
					IndexSet[key]=true
					fmt.Println(value)
					Char = byte(value[len(value)-1])
				}
			}
		}
		if S==Char{
			break;
		}
	}
	return

}
