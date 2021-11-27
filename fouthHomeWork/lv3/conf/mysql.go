package conf

import "database/sql"
var DB *sql.DB
var err error
func init()  {
	dsn := "root:20010712.@tcp(127.0.0.1:3307)/gin"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

