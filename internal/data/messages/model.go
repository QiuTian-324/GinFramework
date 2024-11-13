package message

import (
	"gin_template/internal/data"
	"time"
)

// Message 表的结构定义
type Message struct {
	data.BaseModel
	SenderID    int64     `gorm:"type:bigint;not null" json:"sender_id"`      // 发送者 ID
	ReceiverID  *int64    `gorm:"type:bigint" json:"receiver_id"`             // 接收者 ID（用户私聊时使用）
	GroupID     *int64    `gorm:"type:bigint" json:"group_id"`                // 群聊 ID（群聊时使用）
	Content     string    `gorm:"type:text;not null" json:"content"`          // 消息内容
	ContentType int       `gorm:"type:tinyint;default:0" json:"content_type"` // 消息类型（0: 文本, 1: 图片, 2: 文件, 3: 视频）
	IsRead      int       `gorm:"type:tinyint;default:0" json:"is_read"`      // 是否已读（0: 未读, 1: 已读）
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// 表名
func (Message) TableName() string {
	return "messages"
}
