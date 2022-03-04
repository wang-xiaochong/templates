package database

import (
	"database/sql"
	"fmt"
)

// DB 连接池对象
var DB *sql.DB

// InitDB 数据库连接
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:xs000809@tcp(127.0.0.1:3306)/example"
	DB, err = sql.Open("mysql", dsn) //open不会校验用户名和密码是否正确
	if err != nil {
		fmt.Println("数据库连接失败",err.Error())
		return err
	}
	err = DB.Ping()
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	DB.SetMaxOpenConns(10) //设置数据库连接池的最大连接数 10
	DB.SetMaxIdleConns(5)  //设置最大空闲连接数
	return nil
}
