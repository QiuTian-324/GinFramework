package file

import (
	"yug_server/internal/data"
)

// File 表的结构定义
type File struct {
	data.BaseModel
	UserID   uint6464 `gorm:"type:bigint;not null" json:"user_id"`         // 上传者 ID
	FileName string   `gorm:"type:varchar(255);not null" json:"file_name"` // 文件名
	FilePath string   `gorm:"type:varchar(255);not null" json:"file_path"` // 文件路径
	FileType string   `gorm:"type:varchar(50)" json:"file_type"`           // 文件类型
	FileSize int64    `gorm:"type:bigint" json:"file_size"`                // 文件大小（字节）
}

// 表名
func (File) TableName() string {
	return "files"
}
