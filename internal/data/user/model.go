package user

import (
	"gin_template/internal/data/model"
)

// User 表的结构定义
type User struct {
	model.BaseModel
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
}

// 表名
func (User) TableName() string {
	return "users"
}
