package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role 模型
type Role struct {
	ID 			primitive.ObjectID 			`bson:"_id,omitempty"      json:"_id,omitempty" `
	Name     	string 						`bson:"name"               json:"name" binding:"required" form:"name"`
	Apis     	[]primitive.ObjectID   		`bson:"apis"               json:"apis" form:"apis"`
	Description string   					`bson:"description"        json:"description" form:"description"`
}