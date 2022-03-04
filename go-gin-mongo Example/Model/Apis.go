package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// APIGroup 模型
type APIGroup struct {
	ID   primitive.ObjectID `json:"_id,omitempty"`
	Name string             `json:"name" binding:"required" form:"name"`
	Apis []API              `json:"apis" form:"apis"`
	Description string      `json:"description"  form:"description"`
}


