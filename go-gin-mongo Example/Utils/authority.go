package utils

import (
	model "Example/Model"
	redis "Example/Redis"
	"fmt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// SetCK 设置cookie
func SetCK(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 600, "/", ".localhost", false, true) // maxAge:0是永久,可设置为1进行删除cookie
		cookie, err = c.Cookie("gin_cookie")
	}
	fmt.Println("Cookie value:", cookie)
	c.String(200, "Login success!")
}

// SetSession 设置session
func SetSession(c *gin.Context)  {
	session := sessions.Default(c)
	fmt.Println("session:",session)
	//如果浏览器第一次访问返回状态码401，第二次访问则返回状态码200
	if session.Get("user") != "sessiontest" {
		session.Set("user", "sessiontest")
		_:session.Save()
		c.JSON(http.StatusUnauthorized, gin.H{"user": session.Get("user")})
	} else {
		c.String(http.StatusOK, "Successful second visit")
	}
}

// MyClaims 自定义签名
type MyClaims struct {
	//Name string `json:"name" form:"name"`
	User  model.User `json:"user"`
	jwt.StandardClaims
}

// 1、生成JWT

// TokenExpireDuration 设置Token过期时间
const TokenExpireDuration = time.Hour * 24

// JWTSecret 设置JWT自定义密钥
var JWTSecret = []byte("GarfieldIsAHero!")

// GenToken 生成Token
func GenToken(user model.User,reqIP string) (string, error) {

	c := MyClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "jwt_test",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获取完整的编码后的字符串token
	encodeToken,_:=token.SignedString(JWTSecret)
	//将token根据用户ID存于redis中
	err := redis.SetKey(user.ID.Hex()+reqIP, encodeToken)
	if err != nil {
		fmt.Println("Token failed to insert redis")
	}
	return token.SignedString(JWTSecret)
}

// 2、解析JWT

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 后面是一个匿名函数
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}



