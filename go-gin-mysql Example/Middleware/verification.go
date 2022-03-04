package middleware

import (
	// database "Example/Database"
	// model "Example/Model"
	// utils "Example/Utils"
	// "context"

	// "github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
)

// Verification 权限判断
// func Verification() func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		//获取要访问的URL
// 		url := c.Request.URL.Path
// 		method := c.Request.Method
// 		var  api model.API
// 		var permission = false
// 		//数据库查询要访问的URL
// 		err := database.GetMongoDB().Collection("api").FindOne(context.Background(), bson.M{"url": url,"method":method}).Decode(&api)
// 		if err!=nil {
// 			utils.Return(c,utils.NOT_EXIST,err.Error())
// 			c.Abort()
// 			return
// 		}
// 		//获取用户数据角色
// 		//verifications := One.Verification
// 		//获取用户功能角色
// 		authority := One.Authority
// 		//遍历用户数据角色ID
// 		//for _,verification := range verifications{
// 			var one model.Role
// 			err := database.GetMongoDB().Collection("role").FindOne(context.Background(), bson.M{"_id": authority}).Decode(&one)
// 			if err!=nil {
// 				utils.Return(c,utils.NOT_EXIST,err.Error())
// 				c.Abort()
// 				return
// 			}else {
// 				//遍历该角色下拥有的API数组
// 				for _,apiId := range one.Apis{
// 					//如果要访问的URL存在于该角色的API数组中,进行赋权
// 					if api.Id==apiId {
// 						permission=true
// 					}
// 				}
// 			}
// 		//}
// 		//判断是否有权限访问
// 		if permission==true {
// 			c.Next()
// 			return
// 		}else {
// 			util.Return(c,200,util.NO_PERMISSION,"")
// 			c.Abort()
// 			return
// 		}
		
// 	}
// }