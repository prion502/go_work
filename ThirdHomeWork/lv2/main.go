package main

//@Title  ThirdHomework/lv2/main.go
//@Description   并发求素数
//@Author  Dbs(董百顺)
//@Update 2021-11-18

import "fmt"

//@Title  Generate
//@Description   返回生成自然数序列的管道: 2, 3, 4, ...
//@Author  Dbs     2021-11-18 23:45
//@Param   chan int  返回自然数通道

const n=50000

func Generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2;;i++ {
			ch <- i
		}
	}()
	return ch
}

//@Title  IsPrime
//@Description   根据生成的自然数通道，剔除掉通道值的倍数
//@Author  Dbs     2021-11-18 23:45
//@Param   in <-chan int  产生的自然数通道   prime要剔除的素数
//@return chan int   剔除素数倍数后的通道

func IsPrime(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := Generate()
	for {
		prime:= <-ch // 新出现的素数
		if prime>n{   //判断素数的大小
			break
		}
		fmt.Printf("%v\n",prime)
		ch = IsPrime(ch, prime)
	}
}

