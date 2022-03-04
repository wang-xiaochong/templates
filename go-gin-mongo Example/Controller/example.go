package controller

import (
	model "Example/Model"
	service "Example/Service"
	utils "Example/Utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//------------------------------用户操作----------------------------//

// Login 用户登录
func Login(c *gin.Context)  {
	var user model.Login
	if err := c.Bind(&user); err != nil {
		utils.Return(c,utils.PARA_ERROR, err.Error())
		return
	}
	if err :=utils.CheckData(user);err!=nil {
		utils.Return(c, utils.PARA_ERROR, err)
		return
	} 
	service.Login(c,user)
	return
}

// FindAllUser 查找所有用户
func FindAllUser(c *gin.Context)  {
	jsonStr:=utils.Retrieve(c,"user","name","")
	//二进制格式转化成结构体
	var user  []model.User
	_:json.Unmarshal(jsonStr,&user)
	utils.Return(c,utils.SUCCESS,user)
	return
}

// FindAllUserByName 按名称查找用户
func FindAllUserByName(c *gin.Context)  {
	var  key model.KEY
	key.Key= c.Query("key")
	jsonStr:=utils.Retrieve(c,"user","name",key.Key)
	//格式转化成结构体
	var user  []model.User
	_:json.Unmarshal(jsonStr,&user)
	utils.Return(c,utils.SUCCESS,user)
}

// UserInsert 增加用户
func UserInsert(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		utils.Return(c, utils.PARA_ERROR , err.Error())
		return
	}

//自定义检验规则
	if err :=utils.CheckData(user);err!=nil {
		utils.Return(c, utils.PARA_ERROR, err.Error())
		return
	}
	service.UserInsert( c ,user)
	return
}
// UserUpdate 更新用户信息
func UserUpdate(c *gin.Context)  {
	var  user model.User
	if err:=c.Bind(&user);err !=nil {
		utils.Return(c,  utils.PARA_ERROR, err.Error())
		return
	}
	service.UserUpdate(c,user)
	return
}

// UserDelete 用户删除
func UserDelete(c *gin.Context)  {
	var user model.User
	if err:=c.Bind(&user);err !=nil {
		utils.Return(c,utils.PARA_ERROR, err.Error())
		return
	}
	utils.DeleteOne(c,"user","_id",user.ID)
	return
}

//-----------------------------------------------------------------//

//  -------------------简单增删改查模板-----------------------------//

// Add 增加示例
func Add(c *gin.Context) {
	var example model.Example
	if err := c.Bind(&example); err != nil {
		utils.Return(c, utils.PARA_ERROR , err.Error())
		return
	}
	utils.CreateOne(c,"example",example,"name","")
	return
}

// Search 查找示例
func Search(c *gin.Context)  {
	var  key model.KEY
	key.Key= c.Query("key")
	jsonStr:=utils.Retrieve(c,"example","name",key.Key)
	//格式转化成结构体
	var example  []model.Example
	_:json.Unmarshal(jsonStr,&example)
	utils.Return(c,utils.SUCCESS,example)
	return
}

// Update 更改示例
func Update(c *gin.Context)  {
	var  example model.Example
	if err:=c.Bind(&example);err !=nil {
		utils.Return(c,  utils.PARA_ERROR, err.Error())
		return
	}
	utils.UpdateOne(c,"example",example,"_id",example.ID)
	return
}

// Delete 删除示例
func Delete(c *gin.Context)  {
	var key model.KEY
	if err:=c.Bind(&key);err !=nil {
		utils.Return(c,utils.PARA_ERROR, err.Error())
		return
	}
	utils.DeleteOne(c,"example","_id",key.ID)
	return
}

//-----------------------------------------------------------------//

//---------------------------中间件拦截-----------------------------//

// Home 路由处cookie中间件进行拦截
func Home(c *gin.Context)  {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "访问成功",
		"data": "携带cookie",
	})
}
// Home2 路由出token中间件进行拦截
func Home2(c *gin.Context)  {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "访问成功",
		"data": "携带token",
	})
}