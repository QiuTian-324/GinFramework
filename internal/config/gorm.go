package config

import (
	"time"
	"yug_server/internal/libs"
	"yug_server/pkg"

	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// GormInternal 初始化 GORM 数据库
func GormInternal() (*gorm.DB, error) {
	// 设置日志模式
	logMode := logger.Warn
	if viper.GetBool("development.develop") {
		logMode = logger.Info
	}

	// 打开数据库连接
	db, err := gorm.Open(Mysqlinternal(), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加复数
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		pkg.Error("GORM 初始化失败", err)
		return nil, err
	}
	pkg.Info("GORM 初始化成功")

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		pkg.Error("获取 SQL DB 实例失败", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(viper.GetInt("mysql.conn_max_lifetime")))

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		pkg.Error("数据库连接测试失败", err)
		return nil, err
	}
	pkg.Info("数据库连接成功")

	// 自动迁移表结构
	err = libs.AutoMigrate(db)
	if err != nil {
		pkg.Error("数据库表迁移失败", err)
		return nil, err
	}

	return db, nil
}
