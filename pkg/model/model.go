package model

import (
	"fmt"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {

	var err error

	// 初始化 MySQL 连接信息
	gormConfig := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
			config.GetString("database.mysql.username"),
			config.GetString("database.mysql.password"),
			config.GetString("database.mysql.host"),
			config.GetString("database.mysql.port"),
			config.GetString("database.mysql.database"),
			config.GetString("database.mysql.charset")),
	})

	var level gormlogger.LogLevel
	if config.GetBool("app.debug") {
		// 读取不到数据也会显示，需要打印所有sql记录则设置为info
		level = gormlogger.Warn
	} else {
		// 只有错误才会显示
		level = gormlogger.Error
	}

	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})

	logger.LogError(err)

	return DB
}
