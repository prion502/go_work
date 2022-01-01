package model
//学生加入课堂表

type JoinClass struct{
	Id int    `gorm:"primarykey"`   //主键，无实际意义
	StudentEmail string   //学生邮箱号
	TeacherEmail string   //老师邮箱号
	ClassCode string    //课堂码
}
func(JoinClass)Init(){
	DB.AutoMigrate(&JoinClass{})
}
