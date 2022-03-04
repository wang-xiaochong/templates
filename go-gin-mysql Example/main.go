package main

import (
	database "Example/Database"
	redis "Example/Redis"
	utils "Example/Utils"
	router "Example/Router"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init(){
	_ ,err := redis.GetRedisConn()
	err = utils.InitValidator() //验证
	if err !=nil {
		fmt.Println("自定义验证规则初始化失败:",err.Error())
	}
	err = database.InitDB()
	if err !=nil {
		fmt.Println("数据库连接失败:",err.Error())
	}
}

// main 主函数
func main() {
	r := router.SetupRouter()
	//r.Use(middleware.MiddleWare())    //全局中间件
	router.InitRouter(r)
	if err := r.Run(":7001"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
