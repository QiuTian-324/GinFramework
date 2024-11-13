package libs

import (
	"gin_template/internal/data/user"
	"gin_template/pkg"

	"gorm.io/gorm"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&user.UserSetting{},
		&user.User{},
	)
	if err != nil {
		return err
	}
	pkg.Info("数据库表迁移成功")
	return nil
}
