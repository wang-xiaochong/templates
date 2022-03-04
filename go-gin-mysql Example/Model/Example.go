package model


// Example 示例模型
type Example struct {
	ID 			int64						`bson:"id"       json:"id" `
	Name     	string 						`bson:"name"               json:"name"  form:"name"   `
	Sex     	string   		            `bson:"sex"                json:"sex"   form:"sex"`
	Phone       int64   					`bson:"phone"              json:"phone" form:"phone"`
}