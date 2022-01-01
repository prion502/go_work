package model
type Question struct {
	ID int   //主键
	ClassCode string  //课堂码
	Content string  //问题内容
	Way int    //回答方式（抢答或抽答）
	ResponseStudent string   //回答学生
	Answer string   //回答答案
}
func (Question)TableName() string {
	return "question"
}
