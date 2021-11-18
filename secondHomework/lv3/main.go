package main


//@Title  secondHomework/lv3/main.go
//@Description   B站视频
//@Author  Dbs(董百顺)
//@Update 2021-11-12

import (
	"fmt"
)

//Author BiliBili用户

type Author struct {
	Name string  //名字
	VIP bool    //是否是高贵的Vip
	Icon string  //头像
	Signatures string  //签名
	Focus int  //关注人数
}
//Video 视频详情
type Video struct {
	Name string //视频名字
	Author   //视频作者
	Duration  string  //视频时长
	PraiseCount  int     //点赞量
	Collection  int   //收藏量
	CoinAmount  int   //投币量
	TripleAmount  int  //三连量
}
//UploadVideo 发布视频
func(a Author)UploadVideo(name string,duration string) *Video {
	return &Video{
		Name:name,
		Author:a,
		Duration: duration,
		PraiseCount: 0,
		Collection: 0,
		CoinAmount: 0,
		TripleAmount: 0,
	}
}

//ClickPraise 点赞增加

func(v *Video)ClickPraise(){
	v.PraiseCount++
	fmt.Println("本视频已有点赞量:",v.PraiseCount)
}

//ClickCollection  收藏增加

func(v *Video) ClickCollection()  {
	v.Collection++
	fmt.Println("本视频已有收藏量:",v.Collection)

}

//ClickCoin  硬币增加

func(v *Video)ClickCoin()  {
	v.CoinAmount++
	fmt.Println("本视频已有投币量:",v.CoinAmount)
}

//ClickTripe 三连增加

func(v *Video)ClickTriple()  {
	v.TripleAmount++
	v.Collection++
	v.CoinAmount++
	v.PraiseCount++
	fmt.Println("本视频已有三连量:",v.TripleAmount)
}
//menu 菜单函数
func(v *Video) menu(){
	fmt.Println("****************B站视频界面****************")
	fmt.Println("****************视频名称:",v.Name,"****************")
	fmt.Println("****************点击0投币****************")
	fmt.Println("****************点击1收藏****************")
	fmt.Println("****************点击2点赞****************")
	fmt.Println("****************点击3三连****************")
	fmt.Println("****************点击4退出****************")

}
func main() {
	var (
		name string
		time string
		n=1
	)
	Aut:=&Author{
		Name: "弗朗西斯",
		VIP: true,
		Icon: "人物:Lebron James",
		Signatures: "why not!",
		Focus: 20,
	}
	fmt.Printf("输入发布视频名称:")
	fmt.Scanf("%s",&name)
	fmt.Printf("输入视频时长:")
	fmt.Scanf("%s",&time)
	video:=Aut.UploadVideo(name,time)
	video.menu()
	for n!=4{
		fmt.Scanf("%d",&n)
		switch n {
		case 0:
			video.ClickCoin()
			break
		case 1:
			video.ClickCollection()
			break
		case 2:
			video.ClickPraise()
			break
		case 3:
			video.ClickTriple()
			break
		default:
			if n==4{
				break
			}
			fmt.Println("无效操作!")
		}
	}
}
