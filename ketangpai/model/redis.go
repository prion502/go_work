package model

import (
	"github.com/garyburd/redigo/redis"
)
var Redis redis.Conn
func Init() {
	Redis, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic("REDIS IS WRONG!")
		return
	}

}
