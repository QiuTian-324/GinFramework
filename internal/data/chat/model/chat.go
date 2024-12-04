package model

import (
	"yug_server/internal/data"
)

// Chat 表的结构定义
type ChatMsg struct {
	data.BaseModel
	SenderID       uint64 `gorm:"type:bigint;not null" json:"sender_id"`         // 发送者 ID
	ReceiverID     uint64 `gorm:"type:bigint" json:"receiver_id"`                // 接收者 ID（用户私聊时使用）
	GroupID        uint64 `gorm:"type:bigint" json:"group_id"`                   // 群聊 ID（群聊时使用）
	Content        string `gorm:"type:text;not null" json:"content"`             // 消息内容
	ContentType    string `gorm:"type:tinyint;default:0" json:"content_type"`    // 消息类型
	IsRead         int    `gorm:"type:tinyint;default:0" json:"is_read"`         // 是否已读（0: 未读, 1: 已读）
	Url            string `gorm:"type:varchar(255)" json:"url"`                  //  URL
	FileName       string `gorm:"type:varchar(255)" json:"file_name"`            // 文件名
	OfflineMessage bool   `gorm:"type:tinyint;default:0" json:"offline_message"` // 离线消息
}

// 表名
func (ChatMsg) TableName() string {
	return "chat_msg"
}
