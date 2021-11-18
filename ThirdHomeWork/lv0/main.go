package main

//@Title  ThirdHomework/lv0/main.go
//@Description   改写加锁
//@Author  Dbs(董百顺)
//@Update 2021-11-18

import (
	"fmt"
)

var(
	myRes    =make(map[int]int,20)  //存储n的阶乘数据
	Receiver =make(chan int)    //定义无缓冲通道，实现加锁
)

//@Title  factorial
//@Description   并发进行计算n的阶乘
//@Author  Dbs     2021-11-18 23:45
//@Param   n   所计算数的值

func factorial(n int){
	var res=1
	for i:=1;i<=n;i++{
		res*=i
	}
	myRes[n]=res
	<-Receiver
}
func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)    //开启20个goroutine
		Receiver <- i     //实现同步，接受值的同时，发送值也必须进行
	}
	for i, v := range myRes {
		fmt.Printf("myres[%d]=%d\n", i, v)
	}
}



