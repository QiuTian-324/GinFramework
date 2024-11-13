package group

import (
	"gin_template/internal/data"
	"time"
)

// Group 表的结构定义
type Group struct {
	data.BaseModel
	GroupName   string    `gorm:"type:varchar(100);not null" json:"group_name"` // 群组名称
	OwnerID     int64     `gorm:"type:bigint;not null" json:"owner_id"`         // 群主 ID
	AvatarUrl   string    `gorm:"type:varchar(255)" json:"avatar_url"`          // 群组头像
	Description string    `gorm:"type:text" json:"description"`                 // 群组简介
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// 表名
func (Group) TableName() string {
	return "groups"
}

// GroupMember 表的结构定义
type GroupMember struct {
	data.BaseModel
	GroupID  int64     `gorm:"type:bigint;not null" json:"group_id"` // 群组 ID
	UserID   int64     `gorm:"type:bigint;not null" json:"user_id"`  // 用户 ID
	Role     int       `gorm:"type:tinyint;default:0" json:"role"`   // 角色（0: 普通成员, 1: 管理员, 2: 群主）
	JoinedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"joined_at"`
}

// 表名
func (GroupMember) TableName() string {
	return "group_members"
}
