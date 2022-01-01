package model
//课堂数据库
type Class struct{
	Name string   //课程名称
	TeacherEmail string//教师邮箱
	ClassName string//教学班级
	StudentNum int//学生人数
	Id string    `gorm:"primaryKey"`   //课堂码
	Year string   //学年
	Term string   //学期
	Status int   //状态
}
func (Class)Init(){
	DB.AutoMigrate(&Class{})
}
