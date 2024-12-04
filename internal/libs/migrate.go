package libs

import (
	"yug_server/internal/data/user/model"
	"yug_server/pkg"

	"gorm.io/gorm"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.UserSetting{},
		&model.User{},
		&model.Friend{},
		// &chat.ChatMsg{},
	)
	if err != nil {
		return err
	}
	pkg.Info("数据库表迁移成功")
	return nil
}
