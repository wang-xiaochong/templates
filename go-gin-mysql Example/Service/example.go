package service

import (
	model "Example/Model"
	utils "Example/Utils"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login 用户登录验证
func Login(c *gin.Context, login model.Login) {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	// 通过account从数据库中把该用户信息拿出来
	strArray:=utils.Find(c,"user","account",login.Account)
	if len(strArray) == 0 {
		utils.Return(c,utils.NOT_EXIST,nil)
		return
	}
	var users []model.User
	var user model.User
	for k := range strArray {
		json.Unmarshal([]byte(string(strArray[k])), &user)
		users = append(users, user)
	}
	one := users[0]
	// 密码加密进行比较
	err := bcrypt.CompareHashAndPassword([]byte(one.Password), []byte(login.Password+one.Salt))
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

	//接收前台账号注册
	//判断账号是否注册
	strArray:=utils.Find(c,"user","account",user.Account)
	if len(strArray) != 0 {
		utils.Return(c,utils.HAS_EXIST,nil)
		return
	}
	salt := fmt.Sprintf("%d", rand.Int31())
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost) //加密处理
	encodePWD := string(hash)
	user.Password = encodePWD
	user.Salt = salt
	if err != nil {
		fmt.Println(err.Error())
		utils.Return(c, utils.ERROR, err.Error())
		return
	}
	utils.CreateOne(c,"user",user,"account",user.Account)
	return
}

// UserUpdate 用户信息更新
func UserUpdate(c *gin.Context,user model.User)  {
	utils.UpdateOne(c,"user",user,"id",user.ID)
	return
}