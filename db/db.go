package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wonderivan/logger"
	"operations-platform/config"
)

var (
	isInit bool
	DB     *gorm.DB
	err    error
)

// 初始化方法
func Init() {
	//组装数据库连接配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)

	//初始化gorm实例
	DB, err = gorm.Open(config.DBType, dsn)
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	//设置debug日志开关
	DB.LogMode(config.LogMode)
	//设置连接池
	DB.DB().SetMaxIdleConns(config.MaxIdleConns)
	DB.DB().SetMaxOpenConns(config.MaxOpenConns)
	DB.DB().SetConnMaxLifetime(config.MaxLifeTime)

	//isInit = true
	logger.Info("数据库连接成功")
}
