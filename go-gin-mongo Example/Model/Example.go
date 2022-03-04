package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Example 示例模型
type Example struct {
	ID 			primitive.ObjectID 			`bson:"_id,omitempty"      json:"_id,omitempty" `
	Name     	string 						`bson:"name"               json:"name"  form:"name"  binding:"required" `
	Sex     	[]primitive.ObjectID   		`bson:"sex"                json:"sex"   form:"sex"`
	Phone       string   					`bson:"phone"              json:"phone" form:"phone"`
}