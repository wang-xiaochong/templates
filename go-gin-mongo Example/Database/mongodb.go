package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
var client *mongo.Client

// DB 保留数据库连接
var DB *mongo.Database

// InitDB 初始化mongo数据库
func InitDB() (client *mongo.Client) {
	//设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")//本地无密码连接
	// clientOptions := options.Client().ApplyURI("mongodb://用户名:密码@远程主机IP:27017")	//远程带密码连接
	//连接到mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

// GetMongoDB 获取数据库连接
func GetMongoDB() *mongo.Database {
	if client == nil {
		client = InitDB()
	}
	DB = client.Database("example")
	return  DB
}

