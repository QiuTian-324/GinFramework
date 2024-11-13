package user

import (
	"gin_template/internal/data"
	"time"
)

// User 表的结构定义
type User struct {
	data.BaseModel
	Username    string    `gorm:"type:varchar(50);unique;not null" json:"username"` // 用户名
	Password    string    `gorm:"type:varchar(255);not null" json:"password"`       // 密码（加密存储）
	Nickname    string    `gorm:"type:varchar(100);default:null" json:"nickname"`   // 昵称
	Email       string    `gorm:"type:varchar(100);unique" json:"email"`            // 邮箱
	Phone       string    `gorm:"type:varchar(20);unique" json:"phone"`             // 手机号
	AvatarUrl   string    `gorm:"type:varchar(255)" json:"avatar_url"`              // 头像 URL
	Bio         string    `gorm:"type:text" json:"bio"`                             // 个人简介
	Status      int       `gorm:"type:tinyint;default:1" json:"status"`             // 用户状态（1: 正常，0: 禁用）
	LastLoginAt time.Time `gorm:"type:datetime" json:"last_login_at"`               // 最后登录时间
}

// 表名
func (User) TableName() string {
	return "users"
}

// UserSetting 表的结构定义
type UserSetting struct {
	data.BaseModel
	UserID               int64     `gorm:"type:bigint;not null" json:"user_id"`                 // 用户 ID
	ReceiveNotifications int       `gorm:"type:tinyint;default:1" json:"receive_notifications"` // 是否接收通知（1: 接收, 0: 不接收）
	AllowStrangers       int       `gorm:"type:tinyint;default:0" json:"allow_strangers"`       // 是否允许陌生人添加好友（1: 允许, 0: 禁止）
	Theme                string    `gorm:"type:varchar(50);default:'light'" json:"theme"`       // 界面主题（light/dark）
	CreatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// 表名
func (UserSetting) TableName() string {
	return "user_settings"
}

// // 表迁移
// func MigrateUser() {
// 	libs.AutoMigrate(global.DB, &User{}, "user")
// 	libs.AutoMigrate(global.DB, &UserSetting{}, "user_setting")
// }
