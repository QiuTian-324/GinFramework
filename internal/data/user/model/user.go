package model

import (
	"time"
	"yug_server/internal/data"
)

// User 表的结构定义
type User struct {
	data.BaseModel
	Username    string    `gorm:"type:varchar(50);unique;default:null" json:"username"` // 用户名
	Password    string    `gorm:"type:varchar(255)" json:"password"`                    // 密码（加密存储）
	Nickname    string    `gorm:"type:varchar(100);default:null" json:"nickname"`       // 昵称
	Email       string    `gorm:"type:varchar(100);unique;default:null" json:"email"`   // 邮箱
	Phone       string    `gorm:"type:varchar(20);unique;default:null" json:"phone"`    // 手机号
	AvatarUrl   string    `gorm:"type:varchar(255);default:null" json:"avatar_url"`     // 头像 URL
	Bio         string    `gorm:"type:text" json:"bio"`                                 // 个人简介
	LoginType   int       `gorm:"type:tinyint;default:0" json:"login_type"`             // 登录方式（0:账号密码，1:邮箱，2:手机号，3:第三方）
	Online      int       `gorm:"type:tinyint;default:0" json:"online"`                 // 是否在线（1: 在线，0: 离线）
	Status      int       `gorm:"type:tinyint;default:1" json:"status"`                 // 用户状态（1: 正常，0: 禁用）
	LastLoginAt time.Time `gorm:"type:datetime;default:null" json:"last_login_at"`      // 最后登录时间
	Friends     []Friend  `gorm:"many2many:friends;"`
}

// 表名
func (User) TableName() string {
	return "users"
}

// UserSetting 表的结构定义
type UserSetting struct {
	data.BaseModel
	UserID               int64  `gorm:"type:bigint;not null" json:"user_id"`                 // 用户 ID
	ReceiveNotifications int    `gorm:"type:tinyint;default:1" json:"receive_notifications"` // 是否接收通知（1: 接收, 0: 不接收）
	AllowStrangers       int    `gorm:"type:tinyint;default:0" json:"allow_strangers"`       // 是否允许陌生人添加好友（1: 允许, 0: 禁止）
	Theme                string `gorm:"type:varchar(50);default:'light'" json:"theme"`       // 界面主题（light/dark）
}

// 表名
func (UserSetting) TableName() string {
	return "user_settings"
}

// Friend 表的结构定义
type Friend struct {
	data.BaseModel
	UserID   uint64 `gorm:"type:bigint;not null" json:"user_id"`   // 当前用户 ID
	FriendID uint64 `gorm:"type:bigint;not null" json:"friend_id"` // 好友用户 ID
	Status   int    `gorm:"type:tinyint;default:0" json:"status"`  // 好友状态（0: 待接受, 1: 已接受, 2: 拒绝）
	User     User   `gorm:"foreignKey:UserID;references:ID"`       // 当前用户
	Friend   User   `gorm:"foreignKey:FriendID;references:ID"`     // 好友用户
}

// 表名
func (Friend) TableName() string {
	return "friends"
}
