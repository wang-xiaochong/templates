package utils

import (
	database "Example/Database"
	"context"
	"encoding/json"
	"log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Retrieve 单字段模糊搜索查找接口
func Retrieve(c *gin.Context,collection string,searchKey string,value string) []byte {
	Collection := database.DB.Collection(collection)
	cur, err := Collection.Find(context.Background(), bson.M{searchKey:  primitive.Regex{Pattern: value}})
	if err != nil {
		Return(c,EXCEPTION,err.Error())
		log.Println(err)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	var all []bson.M
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Println(err)
	}
	_:cur.Close(context.Background())
	//log.Println("collection.Find curl.All: ", all)
	// for _, one := range all {
	// 	log.Println(one)
	// }
	jsonStr, err := json.Marshal(all)
	return jsonStr
}

// UpdateOne 根据ID进行单个更新接口
func UpdateOne(c *gin.Context,collection string,data interface{},searchKey string,value primitive.ObjectID )  {
	Collection := database.DB.Collection(collection)
	update := bson.M{"$set": data}
	//fmt.Println("data:",data,"value:",value)
	updateResult, err := Collection.UpdateOne(context.Background(), bson.M{searchKey:value}, update)
	if err != nil {
		Return(c,EXCEPTION,err.Error())
		log.Println(err)
		return
	}
	log.Println("collection.UpdateOne:", updateResult)
	Return(c,SUCCESS,data)
}

// DeleteOne 根据ID进行单个删除接口
func DeleteOne(c *gin.Context,collection string,deleteKey string,value primitive.ObjectID )  {
	Collection := database.DB.Collection(collection)
	deleteResult, err := Collection.DeleteOne(context.Background(), bson.M{deleteKey:value})
	if err != nil {
		Return(c,EXCEPTION,err.Error())
		log.Println(err)
		return
	}
	Return(c,SUCCESS,deleteResult)
	return
}

// CreateOne 增加一条数据,相同字段值不可重复增加
func CreateOne(c *gin.Context,collection string,data interface{},searchKey string,value string)  {
	Collection := database.DB.Collection(collection)
	var one bson.M
	var One interface{}
	err := Collection.FindOne(context.Background(), bson.M{searchKey: value}).Decode(&one)
	if err != nil {
		One = data
		insertOneResult, err := Collection.InsertOne(context.Background(), One)
		if err != nil {
			Return(c,EXCEPTION,err.Error())
			log.Println(err)
		}
		log.Println("collection.InsertOne: ", insertOneResult.InsertedID)
		Return(c,SUCCESS,insertOneResult)
		return
	}
	Return(c,HAS_EXIST,one)
	return
}