package middleWare

import (
	"fifthHomework/conf"
	"fifthHomework/tools"
	"github.com/gin-gonic/gin"
	"strings"
)
type User struct {
	Name string
	Password string
}
func LoginMiddle()gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString:=context.GetHeader("Authorization")
		if tokenString=="" || !strings.HasPrefix(tokenString,"Bearer "){
			context.Set("user",nil)
			return
		}
		tokenString=tokenString[7:]
		token,claims,err:=tools.ParseToken(tokenString)
		if err!=nil || !token.Valid{
			context.Set("user",nil)
			return
		}
		err=conf.DB.Ping()
		if err!=nil{
			context.Set("user",nil)
			return
		}
		var UserStruct User
		sql:="select * from user where username=?"
		conf.DB.QueryRow(sql,claims.UseName).Scan(&UserStruct.Name,&UserStruct.Password)
		if UserStruct.Name!=claims.UseName{
			context.Set("user",nil)
			return
		}
		context.Set("user",claims.UseName)
	}
}
