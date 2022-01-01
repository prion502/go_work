package model
import "gorm.io/gorm"
type User struct {
	gorm.Model
	Email string `gorm:"primaryKey"`   //邮箱号
	Password string    //密码
	Name string   //姓名
	Identity int //身份（学生或教师）
	School string //学校
}
func (User)Init(){
	DB.AutoMigrate(&User{})
}
