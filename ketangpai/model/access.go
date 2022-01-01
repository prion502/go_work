package model

type Access struct {
	Id int   `gorm:"primaryKey"`  //主键
	Identity int       //角色(学生或老师)
	AccessName string   //权限名称
	Path string    //域名路径
}
func (Access)TableName()string{
	return "access"
}