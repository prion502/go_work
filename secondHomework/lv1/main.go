package main

//@Title  secondHomework/lv1/main.go
//@Description   实现接口
//@Author  Dbs(董百顺)
//@Update 2021-11-12
import "fmt"

//定义Person结构体

type Person struct {
	name string  //姓名
	age int //年龄
	gender string //性别
}
// dove 鸽子
type dove interface {
	gugugu()   //鸽
}
//repeater 复读机
type repeater interface {
	repeat(string2 string)   //复读
}
//lemonMaster 柠檬精
type lemonMaster interface {
	lemon()    //酸
}
//trueFragrantMaster   真香怪
type trueFragrantMaster interface {
	trueFragrant()  //我王境泽就是死，从这跳下去，也不会吃你们一口饭---真香
}
//真香方法
func (p *Person)trueFragrant()  {
	fmt.Printf("%s就是死，从这跳下去，也不会吃你们一口饭---真香\n",p.name)
}
//复读方法
func (p *Person)repeat(string2 string){
	fmt.Printf("%s说:%s,%s,%s\n",p.name,string2,string2,string2)
}
//柠檬精方法
func (p *Person)lemon()  {
	fmt.Printf("%s总是自带酸臭味!\n",p.name)
}
//鸽子方法
func (p *Person)gugugu()  {
	fmt.Println(p.name,"又鸽了")
}
func main() {
	//Person类型实例化
	p:= &Person{
		name: "DBS",
		age: 20,
		gender: "男",
	}
	//使用复读机
	r:=repeater(p)
	r.repeat("复读机")
	//使用鸽子
	d:=dove(p)
	d.gugugu()
	//使用真香怪
	t:=trueFragrantMaster(p)
	t.trueFragrant()
	//使用柠檬精
	l:=lemonMaster(p)
	l.lemon()
}
