package model

import "go.mongodb.org/mongo-driver/bson/primitive"


// API 模型
type API struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"  json:"_id,omitempty"`
	Name        string             `bson:"name" json:"name" form:"name"  binding:"required"`
	URL         string             `bson:"url"  json:"url" form:"url" binding:"required"`
	Method      string             `bson:"method" json:"method" form:"method" binding:"required"`
	IsValid 	string			   `bson:"is_valid" json:"is_valid" form:"isValid" `
	Type		string			   `bson:"type" json:"type" form:"type" `
	Description string             `bson:"description" json:"description" form:"description"`

}

// KEY 模型
type KEY struct {
	ID			primitive.ObjectID `bson:"_id" json:"_id"`
	Key			string		       `bson:"key" json:"key" form:"key"`
}