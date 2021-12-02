package tools

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)
var jwtSecret =[]byte("dbs20001")
type Claims struct {
	UseName string
	jwt.StandardClaims
}
func ReleaseToken(username string)(string,error){
	expirationTime:=time.Now().Add(7*24*time.Hour)//设置token期限
	claims:=&Claims{
		UseName: username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "dbs",
			Subject: "userToken",
		},
	}
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=tokenClaims.SignedString(jwtSecret)
    if err==nil{
    	return tokenString,err
	}
	return "",err
}
func ParseToken(tokenString string)(*jwt.Token,*Claims,error){
	claims:=&Claims{}
	token,err:=jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret,nil
	})
	return token,claims,err
}
