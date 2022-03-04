package service

import (
	database "Example/Database"
	model "Example/Model"
	utils "Example/Utils"
	"context"
	"fmt"
	"log"
	"math/rand"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login 用户登录验证
func Login(c *gin.Context, user model.Login) {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	collection := database.DB.Collection("user")
	var one model.User
	//通过account从数据库中把该用户信息拿出来
	err := collection.FindOne(context.Background(), bson.M{"account": user.Account}).Decode(&one)
	if err != nil {
		utils.Return(c, utils.NOT_EXIST, err.Error())
		return
	} 
	//密码加密进行比较
	err = bcrypt.CompareHashAndPassword([]byte(one.Password), []byte(user.Password+one.Salt))
	if err != nil {
		fmt.Println("Login err:", err)
		utils.Return(c, utils.LOGIN_FAIL, err.Error())
		return
	}
	token, err := utils.GenToken(one,reqIP)
	if err != nil {
		utils.Return(c, utils.SYS_BUSY, err.Error())
		return
	}
	utils.Return(c, utils.SUCCESS, token)
}

// UserInsert 插入用户信息
func UserInsert(c *gin.Context, user model.User) {
	collection := database.DB.Collection("user")
	var one model.User
	//随机生成一个账号
	//for  {
	//	account:=fmt.Sprintf("%d",rand.Int31())
	//	err := collection.FindOne(context.Background(), bson.M{"account": account}).Decode(&one)
	//	if err!=nil {
	//		user.Account = account
	//		break
	//	}
	//}

	//接收前台账号注册
	//判断账号是否注册
	err := collection.FindOne(context.Background(), bson.M{"account": user.Account}).Decode(&one)
	if err == nil {
		utils.Return(c, utils.HAS_EXIST, err)
		return
	} 
	salt := fmt.Sprintf("%d", rand.Int31())
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost) //加密处理
	encodePWD := string(hash)
	user.Password = encodePWD
	user.Salt = salt
	insertOneResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err)
		utils.Return(c, utils.ERROR, err.Error())
		return
	}
	log.Println("collection.InsertOne: ", insertOneResult.InsertedID)
	//util.Return(c,util.SUCCESS,user.Account)
	utils.Return(c, utils.SUCCESS, insertOneResult)
	return
}
// UserUpdate 用户信息更新
func UserUpdate(c *gin.Context,user model.User)  {
	utils.UpdateOne(c,"user",user,"_id",user.ID)
	return
}