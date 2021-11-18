package main

//@Title  CSA/calculate/main.go
//@Description   词语接龙
//@Author  Dbs(董百顺)
//@Update 2021-11-05

import (
	"errors"
	"fmt"
)
//定义全局变量
var (
	firstNum int   //第一个数
	secondNum int  //第二个数
	operator byte  //运算符
)

//@Title  switchOperator
//@Description   进行简单的计算
//@Author  Dbs     2021-11-05 10:23
//@Param nil
//@return value,err3  int,error   "value存储计算结果，error存储计算过程中的错误"

func Calculate()(int,error){
	var err error
	switch operator {
	case '+':
		return firstNum+secondNum,err
	case '-':
		return firstNum-secondNum,err
	case '*':
		return firstNum*secondNum,err
	case '/':
		if secondNum==0{
			err=errors.New("除数不能为0")
			return 0,err
		}
		return firstNum/secondNum,err
	default:
		break
	}
	err=errors.New("运算符出错")
	return 0,err
}
func main() {
	//for 循环持续输入计算的数进行计算，输入运算符为%时，退出计算
	//样例
	// 1 + 2
	// 3
	// 2 % 4
	for {
		_,err := fmt.Scanf("%d",&firstNum)
		if err != nil {
			fmt.Println("第一个数输入错误",err);
		}
		_, err1 := fmt.Scanf("%c",&operator)
		if err1 != nil {
			fmt.Println("运算符输入错误",err1)
		}
		if operator=='%'{
			break
		}
		_, err2 :=fmt.Scanf("%d",&secondNum)
		if err2 !=nil{
			fmt.Println("第二个数输入错误:",err2)
		}
		value, err3 :=Calculate()
		if err3 !=nil{
			fmt.Println("运算过程错误:", err3)
		}else {
			fmt.Println(value)
		}
	}
}