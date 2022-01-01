package model

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	ClassCode string  //课堂码
	StudentEmail string  //学生邮箱
	StudentGrade int   //成绩
	StudentName  string  //学生姓名
	HomeworkFile string  //作业名称
}
func(Grade)Init()  {
	DB.AutoMigrate(&Grade{})
}
