package model

type Title struct {
	Id int   `gorm:"primarykey"`   //话题序号
	ClassCode string     //课堂码
	TitleName    string   //话题内容
	TeacherEmail string   //教师邮箱号
	CreateTime int64   //发布话题时间
}

func (Title)Init(){
	DB.AutoMigrate(&Title{})
}