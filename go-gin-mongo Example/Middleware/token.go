package middleware

import (
	model "Example/Model"
	redis "Example/Redis"
	utils "Example/Utils"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// // 4、实现校验的中间件
// // JWTAuthMiddleware 基于JWT的认证中间件

// One 半全局用户信息
var One model.User

// JWTAuthMiddleware 获取接收到的token并解析
func JWTAuthMiddleware() func(c *gin.Context)   {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization") // 获取请求头中的数据
		if authHeader == "" {
			utils.Return(c,utils.TOKEN_MISSING,"Token丢失")
			// 不进行下面的请求处理了！
			c.Abort()
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Beaver") {
			utils.Return(c,utils.Token_Faild,"")
			// 不进行下面的请求处理了！
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.Return(c,utils.TOKEN_INVALID,"")
			// 不进行下面的请求处理了！
			c.Abort()
			return
		}
		//存储用户信息
		One = mc.User
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("user", mc.User)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

// One存储的是请求中token携带的用户信息

// TokenCheck 检验用户携带的token是否是本用户应该携带的token
func TokenCheck() func(c *gin.Context)  {
	return func(c *gin.Context) {
		reqIP := c.ClientIP()
		if reqIP == "::1" {
			reqIP = "127.0.0.1"
		}
		redisToken,err := redis.GetKey(One.ID.Hex()+reqIP)
		var tmp model.User
		mc, err := utils.ParseToken(redisToken)
		if err != nil {
			utils.Return(c,utils.TOKEN_EXPIRE,err.Error())
			// 不进行下面的请求处理了！
			c.Abort()
			return
		}
		tmp = mc.User
		if reflect.DeepEqual(One,tmp){
			c.Next()
			return
		}
		utils.Return(c,utils.TOKEN_INVALID,"Token不匹配")
		c.Abort()
	}
}
