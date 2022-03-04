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
	strArray:=utils.Retrieve(c,"user","name","")
	//二进制格式转化成结构体
	var users []model.User
	var user  model.User
	for k := range strArray {
		json.Unmarshal([]byte(string(strArray[k])), &user)
		users = append(users, user)
    }
	utils.Return(c,utils.SUCCESS,users)
}

// FindAllUserByName 按名称查找用户
func FindAllUserByName(c *gin.Context)  {
	var  key model.KEY
	key.Key= c.Query("key")
	strArray:=utils.Retrieve(c,"user","name",key.Key)
	//二进制格式转化成结构体
	var users []model.User
	var user  model.User
	for k := range strArray {
		json.Unmarshal([]byte(string(strArray[k])), &user)
		users = append(users, user)
    }
	utils.Return(c,utils.SUCCESS,users)
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
	utils.DeleteOne(c,"user","id",user.ID)
	return
}

//-----------------------------------------------------------------//

//  -------------------简单增删改查模板-----------------------------//

// Search 查找示例
func Search(c *gin.Context) {
	var  key model.KEY
	key.Key= c.Query("key")
	strArray:=utils.Retrieve(c,"example","name",key.Key)
	var examples []model.Example
	var example model.Example
	for k := range strArray {
		json.Unmarshal([]byte(string(strArray[k])), &example)
		examples = append(examples, example)
    }
	utils.Return(c,utils.SUCCESS,examples)
}

// Add 增加示例
func Add(c *gin.Context) {
	var example model.Example
	if err := c.Bind(&example); err != nil {
		utils.Return(c, utils.PARA_ERROR , err.Error())
		return
	}
	utils.CreateOne(c,"example",example,"name",example.Name)
}

// Update 更新数据
func Update(c *gin.Context) {
	var  example model.Example
	if err:=c.Bind(&example);err !=nil {
		utils.Return(c,  utils.PARA_ERROR, err.Error())
		return
	}
	utils.UpdateOne(c,"example",example,"id",example.ID)
	return
}

// Delete 删除数据
func Delete(c *gin.Context) {
	// key 接收要删除的id
	type ID struct {
		ID			int `bson:"id" json:"id"`
	}
	var id ID
	if err:=c.Bind(&id);err !=nil {
		utils.Return(c,  utils.PARA_ERROR, err.Error())
		return
	}
	utils.DeleteOne(c,"example","id",int64(id.ID))
}

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
