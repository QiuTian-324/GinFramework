package initialize

import (
	"gin_template/global"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// GormInitialize  初始化gorm
func GormInitialize() (*gorm.DB, error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	logModel := logger.Info
	if viper.GetBool("development.develop") {
		logModel = logger.Warn
	}
	db, err := gorm.Open(MysqlInitialize(), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logModel),
	})
	if err != nil {
		global.Logger.Error("GORM初始化失败", err)
		return nil, err
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(viper.GetInt("mysql.MaxIdleCons"))
	sqlDb.SetMaxOpenConns(viper.GetInt("mysql.MaxOpenCons"))
	sqlDb.SetConnMaxLifetime(time.Minute * 60)
	return db, nil
}
