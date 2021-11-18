package main

//@Title  ThirdHomework/lv3/main.go
//@Description   利用通道简写context包
//@Author  Dbs(董百顺)
//@Update 2021-11-18

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

//@Title  Printer
//@Description   所开协程
//@Author  Dbs     2021-11-18 23:45
//@Param   Chan  chan struct{}   空结构体通道，实现通知协程结束

func Printer(Chan chan struct{}) {
	defer wg.Done()
	for {
		fmt.Println("go作业")
		time.Sleep(time.Second)  //协程接近一秒运行1次
		select {
		case <-Chan:
			return
		default:
		}
	}
}

func main() {
	var cancel = make(chan struct{})//初始化channel
	wg.Add(1)
	go Printer(cancel)
	time.Sleep(time.Second * 10)
	cancel <- struct{}{}  //主程序休息10秒后，通知协程停止运行
	close(cancel)
	wg.Wait()
	fmt.Println("over")
}
//使用方法:使用通道对所开起协程通知是否关闭，协程大约1秒执行一次，主进程停止10秒，则协程大约要输出10次后，便通知协程停止退出！
