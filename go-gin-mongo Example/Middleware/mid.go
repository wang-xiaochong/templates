package middleware

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

// MiddleWare 中间件示例
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
// Sleeping 中间件执行过程
func Sleeping() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("中间件开始执行了")
		c.Next()
		stop := time.Since(start)
		fmt.Println("time:", stop)
	}
}

// AuthMiddleWare   cookie!=test 不能访问
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("gin_cookie"); cookie == "test"&&err==nil {
			c.Next()
			return
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "访问失败"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
