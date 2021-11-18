package main

//@Title  ThirdHomework/lv0/main.go
//@Description   输出打印
//@Author  Dbs(董百顺)
//@Update 2021-11-18

import (
	"fmt"
	"sync"
)
var (
	Wg sync.WaitGroup
	ChA=make(chan bool,1)
	ChB=make(chan bool,1)
	ChC=make(chan bool,1)
)

//@Title  A
//@Description   打印10次A
//@Author  Dbs     2021-11-18 23:45

func A() {
	for i:=0;i<10;i++{
		if <-ChA{
			fmt.Printf("A ")
			ChB<-true
		}
	}
    Wg.Done()
}

//@Title  B
//@Description   打印10次B
//@Author  Dbs     2021-11-18 23:45

func B()  {
    for i:=0;i<10;i++{
    	if <-ChB{
    		fmt.Printf("B ")
    		ChC<-true
		}
	}
	Wg.Done()
}

//@Title  C
//@Description   打印10次C
//@Author  Dbs     2021-11-18 23:45

func C(){
	for i:=0;i<10;i++{
		if <-ChC{
			fmt.Printf("C ")
			ChA<-true
		}
	}
	Wg.Done()
}

func main() {
	//实现原理，同时开启3个goroutine打印ABC，按顺序打10次
	//先对A通道传入值，这时A函数检测到有值，则A通道发送值，打印A，并向B通道传入值，打印B,B向C传入值，打印C，重复10次
    Wg.Add(3)
    ChA<-true
    go A()
    go B()
    go C()
	Wg.Wait()
}