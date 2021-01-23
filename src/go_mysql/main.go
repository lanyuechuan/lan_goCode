package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Gift struct{
	id string
	name string
	statement string
	status int
	user_id string
}

func main() {
	password := os.Getenv("MYSQLTEST_PASSWORD")
	database, err := sqlx.Open("mysql", "devel_test:" + password + "@tcp(10.252.252.18:3306)/qdtest")
    //database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")  
	if err != nil {
        fmt.Println("open mysql failed,", err)
        return
	}
	var gift Gift
	err1 := database.Select(&gift, "select id,name,statement,status,user_id from order_gift",1)
	if err != nil {
		fmt.Println("exec failed, ", err1)
	}
	fmt.Println(gift)
}