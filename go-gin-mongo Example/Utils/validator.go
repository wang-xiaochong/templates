package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"strconv"
)

var validate = validator.New()

// InitValidator 初始化验证规则
func InitValidator() {
	_ = validate.RegisterValidation("NotAdmin", NotAdmin)
	_ = validate.RegisterValidation("Sex", Sex)
	_ = validate.RegisterValidation("StringNumber", StringNumber)
}

//自定义验证规则

// NotAdmin 不是'admin'字符串
func NotAdmin(f validator.FieldLevel) bool {
	return f.Field().String() != "admin"
}

// Sex 性别限制
func Sex(f validator.FieldLevel) bool {
	return f.Field().String()=="男"||f.Field().String() =="女"||f.Field().String()==""
}

// StringNumber 数字字符串限制
func StringNumber(f validator.FieldLevel) bool {
	role, err := strconv.ParseInt( f.Field().String(), 10, 64)
	if err != nil{
		log.Println(err)
		return false
	}
	return  role<10 &&  role>=0
}

// CheckData 验证函数
func CheckData(data interface{}) error {
	//fmt.Println("this is type", reflect.TypeOf(data))
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("参数错误:",err) //Key: 'Users.Passwd' Error:Field validation for 'Passwd' failed on the 'min' tag
		}
		return err
	}
	return nil
}
