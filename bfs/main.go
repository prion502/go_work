package main

//@Title  CSA/bfs/main.go
//@Description   词语接龙
//@Author  Dbs(董百顺)
//@Update 2021-11-05

import (
	"errors"
	"fmt"
)
//pair 定义坐标结构体
type pair struct{
	x int   //横坐标
	y int   //纵坐标
}
const MAX int=100  //定义二维数组最大存储量

var (
	N,M,T,x,y int   //N 行数，M 列数 T 障碍个数  x,y 坐标输入中间量
	begin pair     //开始坐标
	end  pair      //结束坐标
	Num [MAX][MAX]bool  //存储坐标点是否已走过或是障碍点   是 true 否 false
)
var Path =make([]pair,0,100)

//@Title  Bfs
//@Description   进行广度优先遍历
//@Author  Dbs     2021-11-05 10:23
//@Param   begin   pair   "每一次Bfs的起始点"
//@return value,err3  int,error   "value存储计算结果，error存储计算过程中的错误"

func Bfs(begin pair)error{
	Num[begin.x][begin.y]=true
	if begin.x==end.x && begin.y==end.y{
		return nil
	}else {
		if begin.y-1 >= 0 && begin.y-1 < M && Num[begin.x][begin.y-1] == false {    //向上
			Num[begin.x][begin.y-1] = true
			Temp := pair{begin.x, begin.y - 1,}
			Path = append(Path, Temp)
			Bfs(Temp)
		} else if begin.y+1 >= 0 && begin.y+1 < M && Num[begin.x][begin.y+1] == false { //向下
			Num[begin.x][begin.y+1] = true
			Temp := pair{begin.x, begin.y + 1,}
			Path = append(Path, Temp)
			Bfs(Temp)
		}else if begin.x-1 >= 0 && begin.x-1 < N && Num[begin.x-1][begin.y] == false { //向左
			Num[begin.x-1][begin.y] = true
			Temp := pair{begin.x - 1, begin.y,}
			Path = append(Path, Temp)
			Bfs(Temp)
		}else if begin.x+1 >= 0 && begin.x+1 < N && Num[begin.x+1][begin.y] == false {//向右
			Num[begin.x+1][begin.y] = true
			Temp := pair{begin.x + 1, begin.y,}
			Path = append(Path, Temp)
			Bfs(Temp)
		}else{//返回一个错误类型，说明迷宫走不通
			err:=errors.New("无法走通迷宫")
		    return err
		}
	}
	return nil
}
func main() {
	fmt.Scanf("%d%d%d",&N,&M,&T)
	//初始化Num数组
	for i:=0;i<M;i++{
		for j:=0;j<N;j++{
			Num[i][j]=false
		}
	}
	fmt.Scanf("%d%d%d%d\n",&begin.x,&begin.y,&end.x,&end.y)
	//修改障碍点的Num数组值
	for i:=0;i<T;i++{
		fmt.Scanf("%d%d\n",&x,&y)
		Num[x][y]=true
	}
	Path=append(Path,begin)
	err:=Bfs(begin)
	if err!=nil{
		fmt.Println(err)
	}else{
		for i:=range Path{
			fmt.Println(Path[i].x,Path[i].y)
		}
	}
}
