package main


//@Title  secondHomework/lv2/main.go
//@Description   判断数据类型
//@Author  Dbs(董百顺)
//@Update 2021-11-12


import "fmt"

//@Title  Receiver
//@Description   判断数据类型
//@Author  Dbs     2021-11-12 10:23
//@Param   v interface{}     "空接口"

func Receiver(v interface{}){
	switch v.(type) {
	case int:
		fmt.Println("这是一个int",v)
	case string:
		fmt.Println("这是一个string",v)
	case bool:
		fmt.Println("这是一个bool",v)
	case byte:
		fmt.Println("这是一个byte",v)
	default:
		return
	}
}

func main() {
	var s string
	fmt.Scanf("%s",&s)//判断字符串
	Receiver(s)//判断字符串
	var b bool
	fmt.Scanln(&b)
	Receiver(b)//判断布尔值
	var by byte
	fmt.Scanf("%c\n",&by)
	Receiver(by)//判断字节类型
	var i int
	fmt.Scanf("%d",&i)
	Receiver(i)//判断int类型
}
