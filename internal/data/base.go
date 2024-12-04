package data

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 所有模型的基类，包含ID，创建时间，更新时间，删除时间（软删除字段）以及删除标志
type BaseModel struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // 软删除字段，使用 gorm 的 DeletedAt 类型
	IsDeleted bool           `json:"is_deleted"`                        // 自定义软删除标志，标记记录是否被软删除
}

// 软删除方法
func (b *BaseModel) SoftDelete(db *gorm.DB) error {
	// 设置 IsDeleted 为 true，并标记 DeletedAt 为当前时间
	b.IsDeleted = true
	b.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	return db.Save(b).Error
}

// 恢复已软删除的记录
func (b *BaseModel) Restore(db *gorm.DB) error {
	// 恢复 IsDeleted 标志并清除 DeletedAt
	b.IsDeleted = false
	b.DeletedAt = gorm.DeletedAt{}

	return db.Save(b).Error
}
